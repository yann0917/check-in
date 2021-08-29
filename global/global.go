package global

import (
	"net/http"

	"github.com/imroc/req"
	"github.com/yann0917/check-in/config"
)

var (
	Config    config.Server
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36"
)

type Client struct {
	*req.Req
	Headers http.Header
}

func NewClient(cookie, referer string) *Client {
	header := make(http.Header)
	header.Set("User-Agent", UserAgent)
	header.Set("Referer", referer)
	header.Set("Cookie", cookie)
	return &Client{
		Req:     req.New(),
		Headers: header,
	}
}
