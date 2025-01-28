package test

import (
	"gen_learn/domain"
	"gen_learn/internal/book/service"
	"log"
	"time"
)

var bs = service.BookService{}

func CreateBook() {
	err := bs.Create()
	if err != nil {
		panic("创建书籍失败")
	}
	log.Println("创建三本顶级功法成功，已关联id为1的用户")
}

func FindBookByUserId() {
	books, err := bs.Find()
	if err != nil {
		panic("查找书籍失败")
	}
	log.Println("查找书籍成功：", books)
}

func UpdateBook() {
	oldBook := domain.Book{
		Id:          "3",
		Name:        "如何被富婆包养之秘诀三",
		Price:       999.99,
		PublishDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		Uid:         "1",
	}
	newBook := domain.Book{
		Id:          "3",
		Name:        "好男儿志在四方",
		Price:       0.99,
		PublishDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		Uid:         "1",
	}
	if bs.Update(oldBook, newBook) != nil {
		panic("更新书籍失败")
	}
	log.Println("更新书籍成功")
}

// ClearBookAssociation 清除关联
func ClearBookAssociation() {
	if bs.Clear() != nil {
		panic("清除书籍与用户关联失败")
	}
	log.Println("清除书籍与用户关联成功")
}
