package main

import (
	"fmt"
	"github.com/enginewang/microservice-mall/user/domain/repository"
	"github.com/enginewang/microservice-mall/user/handler"
	"github.com/enginewang/microservice-mall/user/proto/user"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"

	service2 "github.com/enginewang/microservice-mall/user/domain/service"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("lastest"),
	)
	srv.Init()
	db, err := gorm.Open("mysql", "root:qwert789@/mall?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.SingularTable(true)
	// 创建表，只需要调用一次即可
	//rp := repository.NewUserRepository(db)
	//err = rp.InitTable()
	//if err != nil {
	//	fmt.Println(err)
	//}
	userDataService := service2.NewUserDataService(repository.NewUserRepository(db))

	err = user.RegisterUserHandler(srv.Server(), &handler.User{UserDataService: userDataService})
	if err != nil {
		fmt.Println(err)
	}
	// Run service
	if err = srv.Run(); err != nil {
		fmt.Println(err)
	}
}
