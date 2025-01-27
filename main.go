package main

import (
	"gen_learn/dao"
	"gen_learn/dao/query"
	"gen_learn/global"
	"gen_learn/test"
)

func init() {
	dao.InitDB()
	query.SetDefault(global.DB)
}

func main() {
	test.CreateUser("1", "pzj", "pzj") // 增
	test.DeleteUser("1")               // 删

	test.CreateUser("2", "pzj2", "pzj2") // 增
	test.CreateUser("3", "pzj3", "pzj3") // 增
	test.CreateUser("4", "pzj4", "pzj4") // 增

	test.UpdateUser("2", map[string]any{
		"username": "edit2",
	})
}
