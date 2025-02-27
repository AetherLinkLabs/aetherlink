package autodrive

import (
	"fmt"
	"io"
	"net/url"
)

// Download downloads a file from the autodrive server by specifying the CID
func (c *Client) Download(cid string) (io.Reader, error) {
	url := fmt.Sprintf("%s/objects/%s/download", c.endpoint, url.PathEscape(cid))

	return c.get(url)
}
