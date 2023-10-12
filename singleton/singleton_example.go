package singleton

import (
	"sync"
)

type single struct {
	Text string
}

var lock = &sync.Mutex{}

var singleInstance *single

func GetInstance() *single {
	if singleInstance == nil {
		//Garante imutabilidade da intância da variável
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			//Criando a intância caso não exista
			singleInstance = &single{
				Text: "Olá, sou uma intância global e única",
			}
		}
	}

	return singleInstance
}
