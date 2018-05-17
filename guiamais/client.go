package guiamais

import (
	"git.resultys.com.br/lib/lower/convert/encode"
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

// SearchTelefones pesquisa telefones no guiamas
// Retorna array de telefones e se ocorreu bloqueio
func (client *Client) SearchTelefones(nome string, cidade string, estado string) (arr []string, isBlock bool) {
	nome = encode.URL(nome)
	cidade = encode.URL(cidade)
	estado = encode.URL(estado)

	url := client.createURL(str.Format("/search?nome={0}&cidade={1}&estado={2}", nome, cidade, estado))

	return request.GetArrayString(url)
}

func (client *Client) createURL(params string) string {
	return str.Format("http://{0}/guiamais{1}", client.IP, params)
}
