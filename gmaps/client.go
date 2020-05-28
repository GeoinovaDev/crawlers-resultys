package gmaps

import (
	"strconv"

	"git.resultys.com.br/lib/lower/convert/decode"
	"git.resultys.com.br/lib/lower/convert/encode"
	"git.resultys.com.br/lib/lower/net/request"
	"git.resultys.com.br/lib/lower/str"
	"git.resultys.com.br/motor/models/gmaps"
)

// Client ...
type Client struct {
	IP      string
	proxy   string
	timeout int
}

// New cria um client
func New(IP string, proxy string, timeout int) *Client {
	return &Client{IP: IP, proxy: proxy, timeout: timeout}
}

// Search ...
func (c *Client) Search(nome string, cidade string) (gmaps.Company, int, string) {
	nome = encode.URL(nome)
	cidade = encode.URL(cidade)

	proto := Protocol{}
	url := str.Format("http://{0}/google/gmaps/search?nome={1}&cidade={2}&proxy={3}&timeout={4}", c.IP, nome, cidade, c.proxy, strconv.Itoa(c.timeout))
	rq := request.New(url)
	response, err := rq.Get()

	if err != nil {
		return proto.Data, 500, err.Error()
	}

	if proto.Code == 200 {
		decode.JSON(response, &proto)
	}

	return proto.Data, proto.Code, ""
}
