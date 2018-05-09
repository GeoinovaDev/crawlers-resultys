package gplaces

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

// SearchTelefones pesquisa telefones no google places
// Retorna array de telefones e se ocorreu bloqueio
func (client *Client) SearchTelefones(nome string, cep, string, cidade string, language string) (arr []string, isBlock bool) {
	url := client.createURL(str.Format("/phone?nome={0}&cidade={1}&cep={2}&language={3}", nome, cidade, cep, language))
	return request.GetArrayString(url)
}

func (client *Client) createURL(params string) string {
	return str.Format("http://{0}/gplaces/search{1}", client.IP, params)
}
