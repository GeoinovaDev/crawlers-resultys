package facebook

import (
	"github.com/GeoinovaDev/crawlers-resultys/lib/request"
	"github.com/GeoinovaDev/lower-resultys/str"
	"github.com/GeoinovaDev/models-resultys/facebook"
)

// Client struct
type Client struct {
	IP string
}

// New cria o client
func New(ip string) *Client {
	return &Client{IP: ip}
}

// GetDados busca informações na pagina da empresa no facebook
func (client *Client) GetDados(url string, proxy string) (*facebook.Page, int) {
	page := &facebook.Page{}

	proto := request.Get(client.createURL(url, proxy), 5)
	if proto.Code == 200 {
		page.PopuleFromMap(proto.Data.(map[string]interface{}))
	}

	return page, proto.Code
}

func (client *Client) createURL(url string, proxy string) string {
	return str.Format("http://{0}/facebook/search?url={1}&proxy={2}", client.IP, url, proxy)
}
