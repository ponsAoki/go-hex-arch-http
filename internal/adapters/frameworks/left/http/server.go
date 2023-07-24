package http

import (
	interactor_ports "hex/internal/ports/interactors"
	"log"
	"net/http"
)

type Adapter struct {
	api interactor_ports.ArithmeticApplicationServicePort
}

func NewAdapter(api interactor_ports.ArithmeticApplicationServicePort) *Adapter {
	return &Adapter{api: api}
}

func (httpa *Adapter) Run() {
	err := http.ListenAndServe(":9000", httpa.InitRouter())
	if err != nil {
		log.Fatal(err)
	}
}
