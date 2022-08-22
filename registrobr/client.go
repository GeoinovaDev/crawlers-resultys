package registrobr

import (
	"github.com/GeoinovaDev/lower-resultys/exception"
	"github.com/GeoinovaDev/lower-resultys/exec/try"
	"github.com/GeoinovaDev/lower-resultys/net/request"
	"github.com/GeoinovaDev/lower-resultys/str"
	"github.com/GeoinovaDev/crawlers-resultys/registrobr/document"
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
func (client *Client) SearchDocument(domain string) (*document.Document, int) {
	url := client.createURL(str.Format("/rdap?domain={0}", domain))

	return sendRequest(url, 3)
}

func (client *Client) createURL(params string) string {
	return str.Format("http://{0}/registro{1}", client.IP, params)
}

func sendRequest(url string, timeout int) (doc *document.Document, status int) {
	try.New().SetTentativas(3).Run(func() {
		protocol := Protocol{}
		err := request.New(url).GetJSON(&protocol)

		if err != nil {
			panic(err)
		}

		if protocol.Code == 200 {
			doc = protocol.Data
		}

		status = protocol.Code
	}).Catch(func(err string) {
		exception.Raise(err, exception.WARNING)
	})

	return
}
