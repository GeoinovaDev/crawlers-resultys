package google

import (
	"github.com/GeoinovaDev/lower-resultys/exception"
	"github.com/GeoinovaDev/lower-resultys/exec/try"
	"github.com/GeoinovaDev/lower-resultys/net/request"
)

// Count ...
type Count struct {
	Word  string `json:"word"`
	Total int    `json:"total"`
}

type protocolCounter struct {
	Code    int     `json:"code"`
	Status  string  `json:"status"`
	Data    []Count `json:"data"`
	Message string  `json:"message"`
}

func getCounter(url string, timeout int) (counters []Count, code int, message string) {
	try.New().SetTentativas(3).Run(func() {
		protocol := protocolCounter{}

		err := request.New(url).SetTimeout(timeout).GetJSON(&protocol)
		if err != nil {
			panic(err)
		}

		code = protocol.Code
		counters = protocol.Data
		message = ""
	}).Catch(func(msg string) {
		counters = nil
		code = 500
		message = msg

		exception.Raise(msg, exception.WARNING)
	})

	return
}
