package gmaps

import (
	"strconv"

	"github.com/GeoinovaDev/lower-resultys/convert/decode"
	"github.com/GeoinovaDev/lower-resultys/convert/encode"
	"github.com/GeoinovaDev/lower-resultys/net/request"
	"github.com/GeoinovaDev/lower-resultys/str"
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
func (c *Client) Search(nome string, cidade string) ([]gmaps.Company, int, string) {
	nome = encode.URL(nome)
	cidade = encode.URL(cidade)

	proto := Protocol{}
	url := str.Format("http://{0}/google/maps/search?nome={1}&cidade={2}&proxy={3}&timeout={4}", c.IP, nome, cidade, c.proxy, strconv.Itoa(c.timeout))
	rq := request.New(url)
	response, err := rq.Get()

	if err != nil {
		return proto.Data, 500, err.Error()
	}

	decode.JSON(response, &proto)

	return proto.Data, proto.Code, ""
}
