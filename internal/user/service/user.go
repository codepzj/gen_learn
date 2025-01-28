package service

import (
	"errors"
	"gen_learn/dao/query"
	"gen_learn/domain"
)

var (
	_ IUserService = (*UserService)(nil)
)

type UserService struct {
}

type IUserService interface {
	Create(user domain.User) error
	DeleteById(id string) error
	DeleteAll() error
	UpdateById(id string, userMap map[string]any) error
	FindAll() ([]domain.User, error)
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) Create(user domain.User) error {
	u := query.User
	return u.Create(&user)
}

func (us *UserService) DeleteById(id string) error {
	u := query.User
	result, err := u.Where(u.Id.Eq(id)).Delete()
	if result.RowsAffected == 0 {
		return errors.New("删除用户失败")
	}
	return err
}

func (us *UserService) DeleteAll() error {
	u := query.User
	result, err := u.Where(u.Id.Like("%")).Delete()
	if result.RowsAffected == 0 {
		return errors.New("删除全部用户失败")
	}
	return err
}

func (us *UserService) UpdateById(id string, userMap map[string]any) error {
	u := query.User
	result, err := u.Where(u.Id.Eq(id)).Updates(userMap)
	if result.RowsAffected == 0 {
		return errors.New("更新失败")
	}
	return err
}

func (us *UserService) FindAll() ([]domain.User, error) {
	u := query.User
	usersPtr, err := u.Preload(u.Books).Find()
	if err != nil {
		return nil, err
	}
	users := make([]domain.User, len(usersPtr))
	for i, user := range usersPtr {
		users[i] = *user
	}
	return users, nil
}
