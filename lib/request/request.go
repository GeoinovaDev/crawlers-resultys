package request

import (
	"github.com/GeoinovaDev/lower-resultys/exception"
	"github.com/GeoinovaDev/lower-resultys/exec/try"
	"github.com/GeoinovaDev/lower-resultys/net"
	"github.com/GeoinovaDev/lower-resultys/net/request"
	"github.com/GeoinovaDev/crawlers-resultys/lib/convert"
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
