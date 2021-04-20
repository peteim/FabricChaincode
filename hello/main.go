package main

import (
	"github.com/snlansky/coral/pkg/contract"
	"github.com/snlansky/coral/pkg/contract/impl"
)

type MyService struct {
}

func (s *MyService) SayHello(stub contract.IContractStub, name string) string {
	return "hello " + name
}

func main() {
	cc := impl.NewFabricChaincode()
	cc.Register(&MyService{})
	cc.Start()
}
