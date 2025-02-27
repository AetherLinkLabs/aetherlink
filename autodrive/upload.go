package autodrive

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
)

// Upload uploads a file to the autodrive server
func (c *Client) Upload(filename string, r io.Reader) (string, error) {
	res, err := c.createUploadSession(filename)
	if err != nil {
		return "", err
	}

	// TODO: Split large files into multiple chunks and upload concurrently
	if err := c.uploadChunk(res.ID, "0", filename, r); err != nil {
		return "", err
	}

	uploadComplete, err := c.uploadComplete(res.ID)
	if err != nil {
		return "", err
	}

	return uploadComplete.CID, nil
}

type uploadSessionRequest struct {
	Filename      string `json:"filename"`
	UploadOptions any    `json:"uploadOptions"`
}

type uploadSessionResponse struct {
	ID string `json:"id"`
}

func (c *Client) createUploadSession(filename string) (*uploadSessionResponse, error) {
	url := fmt.Sprintf("%s/uploads/file/", c.endpoint)

	res := &uploadSessionResponse{}
	payloadBytes, err := json.Marshal(&uploadSessionRequest{Filename: filename})
	if err != nil {
		return nil, err
	}

	return res, c.post(url, bytes.NewReader(payloadBytes), res, "")
}

type uploadChunkRequest struct {
	Message string `json:"message"`
}

func (c *Client) uploadChunk(uploadID, index, filename string, chunk io.Reader) error {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(part, chunk)
	if err != nil {
		return err
	}

	if err = writer.WriteField("index", index); err != nil {
		return err
	}

	if err = writer.Close(); err != nil {
		return err
	}

	url := fmt.Sprintf("%s/uploads/file/%s/chunk", c.endpoint, uploadID)
	res := &uploadChunkRequest{}
	if err = c.post(url, body, res, writer.FormDataContentType()); err != nil {
		return err
	}

	if res.Message != "Chunk uploaded" {
		return fmt.Errorf("chunk upload failed")
	}

	return nil
}

type uploadCompleteResponse struct {
	CID string `json:"cid"`
}

func (c *Client) uploadComplete(uploadID string) (*uploadCompleteResponse, error) {
	url := fmt.Sprintf("%s/uploads/%s/complete", c.endpoint, uploadID)
	res := &uploadCompleteResponse{}
	return res, c.post(url, nil, res, "")
}
