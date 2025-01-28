package main

import (
	"gen_learn/dao"
	"gen_learn/dao/query"
	"gen_learn/domain"
	"gen_learn/global"
	"gen_learn/test"
	"log"
)

func init() {
	dao.InitDB()
	query.SetDefault(global.DB)
}

func UserAction() {
	log.Println("常规的增删改查")
	log.Println("------------------------------------------------------")
	test.CreateUser(domain.User{Id: "1", Username: "pzj", Password: "pzj"}) // 增
	test.DeleteUser("1")                                                    // 删

	users := []domain.User{
		{Id: "2", Username: "pzj2", Password: "pzj2"},
		{Id: "3", Username: "pzj3", Password: "pzj3"},
		{Id: "4", Username: "pzj4", Password: "pzj4"},
	}
	test.CreateUser(users) // 增

	test.UpdateUser("2", map[string]any{
		"username": "edit2",
	}) // 改

	test.FindAllUsers()  // 查
	test.DeleteAllUser() // 删
	log.Println("------------------------------------------------------")
}

func BookAction() {
	log.Println("带有关联关系的增删改查")
	log.Println("------------------------------------------------------")
	test.CreateUser(domain.User{Id: "1", Username: "pzj", Password: "pzj"}) // 增
	// 创建关联
	test.CreateBook()
	test.FindAllUsers()
	// 查找关联
	test.FindBookByUserId()
	// 更新关联
	test.UpdateBook()
	// 查找关联（查看更新关联后的书籍信息）
	test.FindBookByUserId()

	// 清除关联（清除book表中和用户id绑定逻辑外键映射关系）
	test.ClearBookAssociation()
	test.FindBookByUserId()

	test.DeleteAllUser()
	log.Println("------------------------------------------------------")
}

func main() {
	UserAction()
	BookAction()
}
