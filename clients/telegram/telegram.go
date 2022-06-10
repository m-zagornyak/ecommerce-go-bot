package telegram

import (
	"net/http"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

func New(host string, token string) Client {
	return Client{
		host:     host,
		basePath: "bot" + token,
		client:   http.Client{},
	}
}

func newBasePAth(token string) string {
	return "bot" + token
}

func (c *Client) Updates() {

}

func (c *Client) SendMessage() {
	
}