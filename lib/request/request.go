package request

import (
	"git.resultys.com.br/lib/lower/exec"
	"git.resultys.com.br/lib/lower/net"
	"git.resultys.com.br/lib/lower/net/request"
	"git.resultys.com.br/sdk/crawlers-golang/lib/convert"
)

// GetArrayString executa um Get e retorna um array de string
// Return array string e error
func GetArrayString(url string, timeout int) (arr []string, isBlock bool) {
	response, isBlock := Get(url, timeout)

	if response != nil {
		arr = convert.ArrayInterfaceToArrayString(response.([]interface{}))
	}

	return
}

// Get executa uma requisição get
// Retorna um json ou nil
func Get(url string, timeout int) (response interface{}, isBlock bool) {
	exec.Trying(3, func() {
		protocol := net.Protocol{}
		err := request.New(url).GetJSON(&protocol)
		if err != nil {
			isBlock = false
			response = nil
			panic(err)
		}

		if protocol.Code == 101 {
			isBlock = true
			response = nil
			return
		}

		isBlock = false
		response = protocol.Data
	}, func() {

	}, func() {

	}, func() {

	})

	return
}
