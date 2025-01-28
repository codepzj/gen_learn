package service

import (
	"gen_learn/dao/query"
	"gen_learn/domain"
	"time"
)

var (
	_ IBookService = (*BookService)(nil)
)

type BookService struct {
}

type IBookService interface {
	Create() error
	Find() ([]domain.Book, error)
	Update(oldBook domain.Book, newBook domain.Book) error
	Clear() error
}

func NewBookService() *BookService {
	return &BookService{}
}

func (bs *BookService) Create() error {
	u := query.User
	// 因为user和book之间具有一对多的关系
	books := []*domain.Book{
		{
			Id:          "1",
			Name:        "如何被富婆包养之秘诀一",
			Price:       9.99,
			PublishDate: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			Uid:         "1",
		},
		{
			Id:          "2",
			Name:        "如何被富婆包养之秘诀二",
			Price:       99.99,
			PublishDate: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			Uid:         "1",
		},
		{
			Id:          "3",
			Name:        "如何被富婆包养之秘诀三",
			Price:       999.99,
			PublishDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			Uid:         "1",
		},
	}
	// 将books和id为1的user创建关联
	return u.Books.Model(&domain.User{Id: "1"}).Append(books...)
}

func (bs *BookService) Find() ([]domain.Book, error) {
	u := query.User
	booksPtr, err := u.Books.Model(&domain.User{Id: "1"}).Find()
	if err != nil {
		return nil, err
	}
	books := make([]domain.Book, len(booksPtr))
	for i, book := range booksPtr {
		books[i] = *book
	}
	return books, err
}

func (bs *BookService) Update(oldBook domain.Book, newBook domain.Book) error {
	u := query.User
	// 什么鬼啊
	return u.Books.Model(&domain.User{Id: "1"}).Replace(&oldBook, &newBook)
}

func (bs *BookService) Clear() error {
	u := query.User
	return u.Books.Model(&domain.User{Id: "1"}).Clear()
}
