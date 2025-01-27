package service

import (
	"errors"
	"gen_learn/dao/query"
	"gen_learn/domain"
)

var _ IUserService = (*UserService)(nil)

type UserService struct {
}

type IUserService interface {
	Create(u domain.User) error
	DeleteById(id string) error
	UpdateById(id string, userMap map[string]any) error
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) Create(u domain.User) error {
	return query.User.Create(&u)
}

func (us *UserService) DeleteById(id string) error {
	result, err := query.User.Where(query.User.Id.Eq(id)).Delete()
	if result.RowsAffected == 0 {
		return errors.New("删除失败")
	}
	return err
}

func (us *UserService) UpdateById(id string, userMap map[string]any) error {
	result, err := query.User.Where(query.User.Id.Eq(id)).Updates(userMap)
	if result.RowsAffected == 0 {
		return errors.New("更新失败")
	}
	return err
}
