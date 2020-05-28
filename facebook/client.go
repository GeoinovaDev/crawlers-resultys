package facebook

import (
	"git.resultys.com.br/lib/lower/str"
	"git.resultys.com.br/motor/models/facebook"
	"git.resultys.com.br/sdk/crawlers-golang/lib/request"
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
func (client *Client) GetDados(url string) *facebook.Page {
	page := &facebook.Page{}

	proto := request.Get(client.createURL(url), 5)
	if proto.Code == 200 {
		page.PopuleFromMap(proto.Data.(map[string]interface{}))
	}

	return page
}

func (client *Client) createURL(url string) string {
	return str.Format("http://{0}/facebook/search?url={1}", client.IP, url)
}
