package google

import (
	"strconv"

	"github.com/GeoinovaDev/lower-resultys/convert/encode"
	"github.com/GeoinovaDev/lower-resultys/str"
	"github.com/GeoinovaDev/crawlers-resultys/lib/request"
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
func (client *Client) Counter(nome string, telefone string) ([]Count, int, string) {
	nome = encode.URL(nome)
	telefone = encode.URL(telefone)

	url := client.createURL("search", str.Format("/counter?keyword={0}&sentence={1}", telefone, nome))

	return getCounter(url, client.timeout+1)
}

// IsContador ...
func (client *Client) IsContador(telefone string) (bool, int, string) {
	telefone = encode.URL(telefone)
	url := client.createURL("contador", str.Format("/verify?telefone={0}", telefone))
	proto := request.Get(url, client.timeout+1)

	if proto.Code == 200 {
		return proto.Data.(bool), proto.Code, ""
	}

	return false, proto.Code, proto.Message
}

// SearchTelefones pesquisa telefones na pagina principal do google
// Retorna array de telefones e se ocorreu bloqueio
func (client *Client) SearchTelefones(nome string, cidade string) ([]string, int, string) {
	nome = encode.URL(nome)
	cidade = encode.URL(cidade)

	url := client.createURL("search", str.Format("/phones?nome={0}&cidade={1}", nome, cidade))

	return request.GetArrayString(url, client.timeout+1)
}

// SearchFacebook pesquisa todos os links do facebook na pagina principal do google
// Return array string e se ocorreu bloqueio
func (client *Client) SearchFacebook(nome string, cidade string) ([]string, int, string) {
	nome = encode.URL(nome)
	cidade = encode.URL(cidade)

	url := client.createURL("search", str.Format("/facebook?nome={0}&cidade={1}", nome, cidade))

	return request.GetArrayString(url, client.timeout+1)
}

// SearchTwitter pesquisa todos os links do twitter na pagina principal do google
// Return array string e se ocorreu bloqueio
func (client *Client) SearchTwitter(nome string, cidade string) ([]string, int, string) {
	nome = encode.URL(nome)
	cidade = encode.URL(cidade)

	url := client.createURL("search", str.Format("/twitter?nome={0}&cidade={1}", nome, cidade))

	return request.GetArrayString(url, client.timeout+1)
}

// SearchSite pesquisa todos os links do site na pagina principal do google
// Return array string e se ocorreu bloqueio
func (client *Client) SearchSite(nome string, cidade string) ([]string, int, string) {
	nome = encode.URL(nome)
	cidade = encode.URL(cidade)

	url := client.createURL("search", str.Format("/site?nome={0}&cidade={1}", nome, cidade))

	return request.GetArrayString(url, client.timeout+1)
}

// SearchLinkedin pesquisa todos os links do twitter na pagina principal do google
// Return array string e se ocorreu bloqueio
func (client *Client) SearchLinkedin(nome string, cidade string) ([]string, int, string) {
	nome = encode.URL(nome)
	cidade = encode.URL(cidade)

	url := client.createURL("search", str.Format("/linkedin?nome={0}&cidade={1}", nome, cidade))

	return request.GetArrayString(url, client.timeout+1)
}

func (client *Client) createURL(service string, params string) string {
	return str.Format("http://{0}/google/{1}{2}&proxy={3}&timeout={4}", client.IP, service, params, client.proxy, strconv.Itoa(client.timeout))
}
