package gplaces

import (
	"git.resultys.com.br/lib/lower/convert/encode"
	"git.resultys.com.br/lib/lower/exec"
	"git.resultys.com.br/lib/lower/net/request"
	"git.resultys.com.br/lib/lower/str"
	"git.resultys.com.br/motor/models/gplaces"
)

// Client struct
type Client struct {
	IP string
}

type protocol struct {
	Status  string           `json:"status"`
	Company *gplaces.Company `json:"data"`
	Message string           `json:"message"`
}

// New ...
func New(IP string) *Client {
	return &Client{IP: IP}
}

// Search ...
func (client *Client) Search(nome string, cidade string, cep string, language string) (company *gplaces.Company) {
	exec.Trying(3, func() {
		url := str.Format("/search?nome={0}&cidade={1}&cep={2}&language={3}", encode.URL(nome), encode.URL(cidade), encode.URL(cep), encode.URL(language))
		url = client.createURL(url)
		protocol := protocol{}

		err := request.New(url).GetJSON(&protocol)
		if err != nil {
			panic(err)
		}

		if protocol.Status != "ok" {
			panic(protocol.Message)
		}

		company = protocol.Company
	}, func() {

	}, func() {

	}, func() {

	})

	return
}

func (client *Client) createURL(params string) string {
	return str.Format("http://{0}/google/places{1}", client.IP, params)
}
