package google

import (
	"strconv"

	"git.resultys.com.br/lib/lower/convert/encode"
	"git.resultys.com.br/lib/lower/str"
	"git.resultys.com.br/sdk/crawlers-golang/lib/request"
)

// Client struct
type Client struct {
	IP      string
	proxy   string
	timeout int
}

// New cria um client
func New(IP string, proxy string, timeout int) *Client {
	return &Client{IP: IP, proxy: proxy, timeout: timeout}
}

// Counter ...
func (client *Client) Counter(nome string, telefone string) ([]Count, bool) {
	nome = encode.URL(nome)
	telefone = encode.URL(telefone)

	url := client.createURL("search", str.Format("/counter?keyword={0}&sentence={1}", telefone, nome))

	return getCounter(url)
}

// IsContador ...
func (client *Client) IsContador(telefone string) (bool, bool) {
	telefone = encode.URL(telefone)
	url := client.createURL("contador", str.Format("/verify?telefone={0}", telefone))
	response, isBlock := request.Get(url)
	if isBlock {
		return false, true
	}

	if response == nil {
		return false, false
	}

	return response.(bool), isBlock
}

// SearchTelefones pesquisa telefones na pagina principal do google
// Retorna array de telefones e se ocorreu bloqueio
func (client *Client) SearchTelefones(nome string, cidade string, cep string, language string) (arr []string, isBlock bool) {
	nome = encode.URL(nome)
	cidade = encode.URL(cidade)
	cep = encode.URL(cep)
	language = encode.URL(language)

	url := client.createURL("search", str.Format("/phones?nome={0}&cidade={1}&cep={2}&language={3}", nome, cidade, cep, language))

	return request.GetArrayString(url)
}

// SearchFacebook pesquisa todos os links do facebook na pagina principal do google
// Return array string e se ocorreu bloqueio
func (client *Client) SearchFacebook(nome string, cidade string) (arr []string, isBlock bool) {
	nome = encode.URL(nome)
	cidade = encode.URL(cidade)

	url := client.createURL("search", str.Format("/facebook?nome={0}&cidade={1}", nome, cidade))

	return request.GetArrayString(url)
}

// SearchTwitter pesquisa todos os links do twitter na pagina principal do google
// Return array string e se ocorreu bloqueio
func (client *Client) SearchTwitter(nome string, cidade string) (arr []string, isBlock bool) {
	nome = encode.URL(nome)
	cidade = encode.URL(cidade)

	url := client.createURL("search", str.Format("/twitter?nome={0}&cidade={1}", nome, cidade))

	return request.GetArrayString(url)
}

// SearchSite pesquisa todos os links do site na pagina principal do google
// Return array string e se ocorreu bloqueio
func (client *Client) SearchSite(nome string, cidade string) (arr []string, isBlock bool) {
	nome = encode.URL(nome)
	cidade = encode.URL(cidade)

	url := client.createURL("search", str.Format("/site?nome={0}&cidade={1}", nome, cidade))

	return request.GetArrayString(url)
}

// SearchLinkedin pesquisa todos os links do twitter na pagina principal do google
// Return array string e se ocorreu bloqueio
func (client *Client) SearchLinkedin(nome string, cidade string) (arr []string, isBlock bool) {
	nome = encode.URL(nome)
	cidade = encode.URL(cidade)

	url := client.createURL("search", str.Format("/linkedin?nome={0}&cidade={1}", nome, cidade))

	return request.GetArrayString(url)
}

func (client *Client) createURL(service string, params string) string {
	return str.Format("http://{0}/google/{1}{2}&proxy={3}&timeout={4}", client.IP, service, params, client.proxy, strconv.Itoa(client.timeout))
}
