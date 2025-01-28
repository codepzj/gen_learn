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
func CreateUser[T domain.User | []domain.User](user T) {
	switch v := any(user).(type) {
	case domain.User:
		if err := us.Create(v); err != nil {
			panic(fmt.Sprintf("创建用户失败：%s", err.Error()))
		}
	case []domain.User:
		for _, u := range v {
			if err := us.Create(u); err != nil {
				panic(fmt.Sprintf("创建用户失败：%s", err.Error()))
			}
		}
	default:
		panic("用户类型不匹配")
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

// DeleteAllUser 删除全部用户
func DeleteAllUser() {
	if err := us.DeleteAll(); err != nil {
		panic(err.Error())
	}
	log.Println("删除全部用户成功")
}

// UpdateUser 更新用户
func UpdateUser(id string, userMap map[string]any) {
	if err := us.UpdateById(id, userMap); err != nil {
		panic(fmt.Sprintf("用户 %s 更新失败", id))
	}
	log.Println("用户更新成功")
}

// FindAllUsers 查找所有用户
func FindAllUsers() {
	users, err := us.FindAll()
	if err != nil {
		panic("查询所有用户失败")
	}
	fmt.Println("查询所有用户成功:", users)
}
