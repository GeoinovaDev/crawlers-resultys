package google

import (
	"git.resultys.com.br/lib/lower/exec"
	"git.resultys.com.br/lib/lower/net/request"
)

// Count ...
type Count struct {
	Word  string    `json:"word"`
	Total int `json:"total"`
}

type protocolCounter struct {
	Code    int     `json:"code"`
	Status  string  `json:"status"`
	Data    []Count `json:"data"`
	Message string  `json:"message"`
}

func getCounter(url string) (counters []Count, isBlock bool) {
	exec.Trying(3, func() {
		protocol := protocolCounter{}
		err := request.New(url).GetJSON(&protocol)
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
	}, func() {

	}, func() {

	}, func() {

	})

	return
}
