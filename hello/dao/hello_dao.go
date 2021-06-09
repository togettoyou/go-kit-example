package dao

import (
	"context"
	"log"
)

// HelloEntity 数据库实体模型
type HelloEntity struct {
	Name string
	Age  uint
}

type HelloDAO interface {
	Query(ctx context.Context) *HelloEntity
}

// ------------------接口实现------------------

type helloDAOImpl struct {
	//db *DB
}

func NewHelloDAOImpl() HelloDAO {
	return &helloDAOImpl{}
}

func (helloDAO *helloDAOImpl) Query(ctx context.Context) *HelloEntity {
	log.Println("dao层-Query")
	// 模拟查询数据库
	return &HelloEntity{Name: "togettoyou", Age: 22}
}
