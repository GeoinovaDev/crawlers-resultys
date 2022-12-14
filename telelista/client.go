package telelista

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

// SearchTelefones pesquisa telefones no telelista
// Retorna array de telefones e se ocorreu bloqueio
func (client *Client) SearchTelefones(nome string, estado string) ([]string, int, string) {
	nome = encode.URL(nome)
	estado = encode.URL(estado)

	url := client.createURL(str.Format("/search?nome={0}&estado={1}", nome, estado))

	return request.GetArrayString(url, 5)
}

func (client *Client) createURL(params string) string {
	return str.Format("http://{0}/telelista{1}", client.IP, params)
}
