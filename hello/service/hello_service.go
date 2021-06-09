package service

import (
	"context"
	"github.com/togettoyou/go-kit-example/hello/dao"
	"log"
)

type HelloService interface {
	GetName(ctx context.Context) string
	GetAge(ctx context.Context) uint
}

// ------------------接口实现------------------

type helloServiceImpl struct {
	helloDAO dao.HelloDAO
}

func NewHelloServiceImpl(helloDAO dao.HelloDAO) HelloService {
	return &helloServiceImpl{
		helloDAO,
	}
}

func (s helloServiceImpl) GetName(ctx context.Context) string {
	log.Println("service层-GetName")
	// 模拟业务处理
	return s.helloDAO.Query(ctx).Name
}

func (s helloServiceImpl) GetAge(ctx context.Context) uint {
	log.Println("service层-GetAge")
	// 模拟业务处理
	return s.helloDAO.Query(ctx).Age
}
