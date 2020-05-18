package registrobr

import (
	"git.resultys.com.br/lib/lower/exception"
	"git.resultys.com.br/lib/lower/exec/try"
	"git.resultys.com.br/lib/lower/net/request"
	"git.resultys.com.br/lib/lower/str"
	"git.resultys.com.br/sdk/crawlers-golang/registrobr/document"
)

// Client struct
type Client struct {
	IP string
}

// New cria um client
func New(IP string) *Client {
	return &Client{IP: IP}
}

// SearchDocument ...
func (client *Client) SearchDocument(domain string) *document.Document {
	url := client.createURL(str.Format("/rdap?domain={0}", domain))

	return sendRequest(url, 3)
}

func (client *Client) createURL(params string) string {
	return str.Format("http://{0}/registro{1}", client.IP, params)
}

func sendRequest(url string, timeout int) (doc *document.Document) {
	try.New().SetTentativas(3).Run(func() {
		protocol := Protocol{}
		err := request.New(url).GetJSON(&protocol)
		if err != nil {
			doc = nil
			panic(err)
		}

		if protocol.Code == 101 {
			doc = nil
			return
		}

		doc = protocol.Data
	}).Catch(func(err string) {
		exception.Raise(err, exception.WARNING)
	})

	return
}
