package request

import (
	"git.resultys.com.br/lib/lower/exec"
	"git.resultys.com.br/lib/lower/net"
	"git.resultys.com.br/lib/lower/net/request"
)

// Get executa uma requisição get
// Retorna um json ou nil
func Get(url string) (response interface{}) {
	exec.Trying(3, func() {
		protocol := net.Protocol{}
		err := request.New(url).GetJSON(&protocol)
		if err != nil {
			panic(err)
		}

		response = protocol.Data
	}, func() {

	}, func() {
		response = nil
	}, func() {

	})

	return
}
