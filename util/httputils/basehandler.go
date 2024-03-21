package httputils

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

type ClientHandler struct {
	Client *resty.Client
}

func NewClientHandle() *ClientHandler {
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetTimeout(3 * time.Minute)
	return &ClientHandler{
		Client: client,
	}
}

// CommonPost 通用post请求
func (c *ClientHandler) CommonPost(url string, body interface{}, headers map[string]string) (*resty.Response, error) {
	resp, err := c.Client.R().SetHeaders(headers).SetBody(body).Post(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("error post, status: %d", resp.StatusCode())
	}
	return resp, nil
}
