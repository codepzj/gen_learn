package service

import (
	"context"
	"gen_learn/dao/model"
	"gen_learn/dao/query"
)

var (
	ctx = context.Background()
)

type BookService struct {
}

type IBookService interface {
	Create(book model.Book) error
}

func NewBookService() *BookService {
	return &BookService{}
}

func (b *BookService) Create(book model.Book) error {
	return query.Book.WithContext(ctx).Create(&book)
}
