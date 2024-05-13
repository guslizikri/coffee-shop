package repository

import (
	"coffee-shop/config"
	"coffee-shop/internal/models"

	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock.Mock
}

func (r *RepoMock) CreateUser(data *models.User) (*config.Result, error) {
	args := r.Mock.Called(data)
	return args.Get(0).(*config.Result), args.Error(1)
}

func (r *RepoMock) ReadUser(params models.Query) (*config.Result, error) {
	args := r.Mock.Called()
	return args.Get(0).(*config.Result), args.Error(1)
}
func (r *RepoMock) UpdateUser(data *models.User) (*config.Result, error) {
	args := r.Mock.Called(data)
	return args.Get(0).(*config.Result), args.Error(1)
}
func (r *RepoMock) DeleteUser(data *models.User) (*config.Result, error) {
	args := r.Mock.Called(data)
	return args.Get(0).(*config.Result), args.Error(1)
}

func (r *RepoMock) GetAuthData(id string) (*models.User, error) {
	args := r.Mock.Called(id)
	return args.Get(0).(*models.User), args.Error(1)
}
func (r *RepoMock) GetUserById(user string) (*config.Result, error) {
	args := r.Mock.Called(user)
	return args.Get(0).(*config.Result), args.Error(1)
}
