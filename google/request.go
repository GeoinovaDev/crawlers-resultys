package google

import (
	"git.resultys.com.br/lib/lower/exception"
	"git.resultys.com.br/lib/lower/exec/try"
	"git.resultys.com.br/lib/lower/net/request"
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

func getCounter(url string, timeout int) (counters []Count, isBlock bool) {
	try.New().SetTentativas(3).Run(func() {
		protocol := protocolCounter{}
		err := request.New(url).SetTimeout(timeout).GetJSON(&protocol)
		if err != nil {
			isBlock = false
			counters = nil
			panic(err)
		}

		if protocol.Code == 101 {
			isBlock = true
			counters = nil
			return
		}

		isBlock = false
		counters = protocol.Data
	}).Catch(func(msg string) {
		exception.Raise(msg, exception.WARNING)
	})

	return
}
