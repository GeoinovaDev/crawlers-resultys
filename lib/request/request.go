package request

import (
	"git.resultys.com.br/lib/lower/exception"
	"git.resultys.com.br/lib/lower/exec/try"
	"git.resultys.com.br/lib/lower/net"
	"git.resultys.com.br/lib/lower/net/request"
	"git.resultys.com.br/sdk/crawlers-golang/lib/convert"
)

// GetArrayString executa um Get e retorna um array de string
// Return array string e error
func GetArrayString(url string, timeout int) (arr []string, code int, message string) {
	protocol := Get(url, timeout)

	code = protocol.Code
	message = protocol.Message

	if protocol.Code == 200 {
		arr = convert.ArrayInterfaceToArrayString(protocol.Data.([]interface{}))
	}

	return
}

// Get executa uma requisição get
// Retorna um json ou nil
func Get(url string, timeout int) (proto net.Protocol) {
	try.New().SetTentativas(3).Run(func() {
		err := request.New(url).GetJSON(&proto)
		if err != nil {
			panic(err)
		}
	}).Catch(func(message string) {
		proto.Code = 500
		proto.Message = message

		exception.Raise(message, exception.WARNING)
	})

	return
}
