package dao

import (
	"fmt"
)

// 数据库实体模型
type HelloEntity struct {
	Msg string
}

func (HelloEntity) TableName() string {
	return "hello"
}

type HelloDAO interface {
	QueryMsg() *HelloEntity
}

// HelloDAO 接口实现
type HelloDAOImpl struct {
}

func (helloDAO *HelloDAOImpl) QueryMsg() *HelloEntity {
	fmt.Println("dao层")
	// 模拟查询数据库
	return &HelloEntity{Msg: "Hello World"}
}
