package apontador

import (
	"github.com/GeoinovaDev/crawlers-resultys/lib/request"
	"github.com/GeoinovaDev/lower-resultys/convert/encode"
	"github.com/GeoinovaDev/lower-resultys/str"
)

// Client struct
type Client struct {
	IP string
}

// New cria um client
func New(IP string) *Client {
	return &Client{IP: IP}
}

// SearchTelefones pesquisa telefones no apontador
// Return array string e se ocorreu bloqueio
func (client *Client) SearchTelefones(nome string, cidade string, estado string) ([]string, int, string) {
	nome = encode.URL(nome)
	cidade = encode.URL(cidade)
	estado = encode.URL(estado)

	url := client.createURL(str.Format("/search?nome={0}&cidade={1}&estado={2}", nome, cidade, estado))
	return request.GetArrayString(url, 5)
}

func (client *Client) createURL(params string) string {
	return str.Format("http://{0}/apontador{1}", client.IP, params)
}
