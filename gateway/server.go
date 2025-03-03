package gateway

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/url"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"

	"github.com/aetherlink/aetherlink/autodrive"
	"github.com/aetherlink/aetherlink/contract"
)

type Server struct {
	ctx        context.Context
	port       uint64
	engine     *gin.Engine
	contract   *contract.AETHER
	autoClient *autodrive.Client
	db         *sql.DB
}

func NewServer(
	ctx context.Context,
	port uint64,
	contractAddr common.Address,
	autoEvmUrl string,
	autoDriveURL string,
	token string,
	dbPath string,
) (*Server, error) {
	client := autodrive.NewClient(ctx, autoDriveURL, token)

	ethClient, err := ethclient.Dial(autoEvmUrl)
	if err != nil {
		return nil, err
	}

	aetherContract, err := contract.NewAETHER(contractAddr, ethClient)
	if err != nil {
		return nil, err
	}

	db, err := initDB(dbPath)
	if err != nil {
		return nil, err
	}

	go NewEventListener(db, aetherContract).Run()

	return &Server{
		ctx:        ctx,
		port:       port,
		engine:     gin.Default(),
		contract:   aetherContract,
		autoClient: client,
		db:         db,
	}, nil
}

func (s *Server) Start() error {
	s.init()
	if err := s.engine.Run(fmt.Sprintf(":%d", s.port)); err != nil {
		return err
	}

	return nil
}

func (s *Server) init() {
	// CORS middleware
	s.engine.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	s.engine.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	s.engine.GET("/:filename", s.rootHandler)
	s.engine.GET("/files", func(ctx *gin.Context) {
		files, err := s.files()
		if err != nil {
			ctx.String(500, "Error getting files")
			return
		}

		ctx.JSON(200, files)
	})
}

// https://<subdomain>.autonomys.site/<filename>
func (s *Server) rootHandler(c *gin.Context) {
	host := c.Request.Host

	// Get the subdomain
	subdomain, err := getSubdomain(host)
	if err != nil {
		slog.Error("Error getting subdomain", "error", err, "host", host)
		c.String(500, "Error getting subdomain")
		return
	}

	if subdomain == "" {
		c.String(404, "subdomain not found")
		return
	}

	userAddress, err := s.contract.GetUserAddress(nil, subdomain)
	if err != nil {
		slog.Error("Error getting user address", "error", err, "username", subdomain)
		c.String(500, "Error getting user address")
		return
	}

	filename := c.Param("filename")
	cid, err := s.contract.GetFileCID(nil, filename, userAddress)
	if err != nil {
		slog.Error("Error getting CID", "error", err, "filename", filename, "userAddress", userAddress)
		c.String(500, "Error getting CID")
		return
	}

	reader, err := s.autoClient.Download(cid)
	if err != nil {
		slog.Error("Error downloading file", "error", err, "cid", cid)
		c.String(500, "Error downloading file")
		return
	}

	c.DataFromReader(200, -1, "text/html", reader, map[string]string{
		"Content-Disposition": `inline; filename="file.html"`,
	})
}

// getSubdomain extracts the subdomain from a given URL or domain.
func getSubdomain(domain string) (string, error) {
	// Check if the domain looks like a URL (http:// or https://)
	if !strings.HasPrefix(domain, "http://") && !strings.HasPrefix(domain, "https://") {
		// Treat the input as a simple domain name
		return extractSubdomain(domain)
	}

	// Parse URL to extract host
	parsedURL, err := url.Parse(domain)
	if err != nil {
		return "", err
	}

	// Extract host from the URL (e.g., aether.localhost:8081)
	host := parsedURL.Hostname()
	return extractSubdomain(host)
}

// extractSubdomain splits the domain into parts and returns the subdomain if it exists
func extractSubdomain(host string) (string, error) {
	// Remove any ports (e.g., example.com:8080)
	host = strings.Split(host, ":")[0]

	// Split the host into parts by '.'
	parts := strings.Split(host, ".")
	if len(parts) > 0 && parts[len(parts)-1] == "localhost" {
		return strings.Join(parts[:len(parts)-1], "."), nil
	}

	if len(parts) < 3 {
		// No subdomain present (only a domain and TLD)
		return "", nil
	}

	// Join all parts before the domain (second to last part) and TLD
	subdomain := strings.Join(parts[:len(parts)-2], ".")
	return subdomain, nil
}

type FileMetaData struct {
	FileName string `json:"filename"`
	UserName string `json:"username"`
}

func (s *Server) files() ([]*FileMetaData, error) {
	rows, err := s.db.Query("SELECT useraddress, username, cid, filename FROM users LIMIT 10")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []*FileMetaData
	for rows.Next() {
		var userAddress, userName, cid, fileName string
		if err := rows.Scan(&userAddress, &userName, &cid, &fileName); err != nil {
			return nil, err
		}

		files = append(files, &FileMetaData{
			FileName: fileName,
			UserName: userName,
		})
	}

	return files, nil
}
