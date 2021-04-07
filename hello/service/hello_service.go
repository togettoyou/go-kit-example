package service

import (
	"context"
	"fmt"
	"github.com/togettoyou/go-kit-example/hello/dao"
)

type HelloService interface {
	SayHello(ctx context.Context) string
}

type HelloServiceImpl struct {
	helloDAO dao.HelloDAO
}

func MakeHelloServiceImpl(helloDAO dao.HelloDAO) HelloService {
	return &HelloServiceImpl{
		helloDAO,
	}
}

func (s HelloServiceImpl) SayHello(ctx context.Context) string {
	fmt.Println("service层")
	return s.helloDAO.QueryMsg().Msg
}
