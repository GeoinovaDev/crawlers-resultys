package telelista

import (
	"git.resultys.com.br/lib/lower/str"
	"git.resultys.com.br/sdk/crawlers-golang/lib/request"
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
func (client *Client) SearchTelefones(nome string, estado string) (arr []string, isBlock bool) {
	url := client.createURL(str.Format("/search?nome={0}&estado={1}", nome, estado))

	return request.GetArrayString(url)
}

func (client *Client) createURL(params string) string {
	return str.Format("http://{0}/telelista{1}", client.IP, params)
}
