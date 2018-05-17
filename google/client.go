package google

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

// IsContador ...
func (client *Client) IsContador(telefone string) (bool, bool) {
	url := str.Format("http://{0}/contador/verify?telefone={1}", client.IP, encode.URL(telefone))
	response, isBlock := request.Get(url)
	return response.(bool), isBlock
}

// SearchTelefones pesquisa telefones na pagina principal do google
// Retorna array de telefones e se ocorreu bloqueio
func (client *Client) SearchTelefones(nome string, cidade string, cep string, language string) (arr []string, isBlock bool) {
	nome = encode.URL(nome)
	cidade = encode.URL(cidade)
	cep = encode.URL(cep)
	language = encode.URL(language)

	url := client.createURL(str.Format("/phones?nome={0}&cidade={1}&cep={2}&language={3}", nome, cidade, cep, language))

	return request.GetArrayString(url)
}

// SearchFacebook pesquisa todos os links do facebook na pagina principal do google
// Return array string e se ocorreu bloqueio
func (client *Client) SearchFacebook(nome string, cidade string) (arr []string, isBlock bool) {
	nome = encode.URL(nome)
	cidade = encode.URL(cidade)

	url := client.createURL(str.Format("/facebook?nome={0}&cidade={1}", nome, cidade))

	return request.GetArrayString(url)
}

// SearchTwitter pesquisa todos os links do twitter na pagina principal do google
// Return array string e se ocorreu bloqueio
func (client *Client) SearchTwitter(nome string, cidade string) (arr []string, isBlock bool) {
	nome = encode.URL(nome)
	cidade = encode.URL(cidade)

	url := client.createURL(str.Format("/twitter?nome={0}&cidade={1}", nome, cidade))

	return request.GetArrayString(url)
}

// SearchSite pesquisa todos os links do site na pagina principal do google
// Return array string e se ocorreu bloqueio
func (client *Client) SearchSite(nome string, cidade string) (arr []string, isBlock bool) {
	nome = encode.URL(nome)
	cidade = encode.URL(cidade)

	url := client.createURL(str.Format("/site?nome={0}&cidade={1}", nome, cidade))

	return request.GetArrayString(url)
}

// SearchLinkedin pesquisa todos os links do twitter na pagina principal do google
// Return array string e se ocorreu bloqueio
func (client *Client) SearchLinkedin(nome string, cidade string) (arr []string, isBlock bool) {
	nome = encode.URL(nome)
	cidade = encode.URL(cidade)

	url := client.createURL(str.Format("/linkedin?nome={0}&cidade={1}", nome, cidade))

	return request.GetArrayString(url)
}

func (client *Client) createURL(params string) string {
	return str.Format("http://{0}/google/search{1}", client.IP, params)
}
