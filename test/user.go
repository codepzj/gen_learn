package test

import (
	"fmt"
	"gen_learn/domain"
	"gen_learn/internal/user/service"
	"log"
)

var (
	us = service.NewUserService()
)

// CreateUser 创建用户
func CreateUser(id string, username string, password string) {
	u := domain.User{
		Id:       id,
		Username: username,
		Password: password,
		Books:    nil,
	}
	if err := us.Create(u); err != nil {
		panic(fmt.Sprintf("创建用户失败：%s", err.Error()))
	}
	log.Println("创建用户成功")
}

// DeleteUser 删除用户
func DeleteUser(id string) {
	err := us.DeleteById(id)
	if err != nil {
		panic(err.Error())
	}
	log.Println("删除用户成功")
}

func UpdateUser(id string, userMap map[string]any) {
	if err := us.UpdateById(id, userMap); err != nil {
		panic(fmt.Sprintf("用户 %s 更新失败", id))
	}
	log.Println("用户更新成功")
}
