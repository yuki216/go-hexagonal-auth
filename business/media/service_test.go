package media_test

import (
	"github.com/stretchr/testify/mock"
	"go-hexagonal-auth/api/v1/auth/request"
	"go-hexagonal-auth/business"
	"go-hexagonal-auth/business/admin"
	adminMock "go-hexagonal-auth/business/admin/mocks"
	"go-hexagonal-auth/business/auth"
	authMock "go-hexagonal-auth/business/auth/mocks"
	"go-hexagonal-auth/business/user"
	userMock "go-hexagonal-auth/business/user/mocks"
	"go-hexagonal-auth/config"

	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	id       = 1
	name     = "name"
	username = "username"
	password = "password"
	repassword = "password"
	creator  = "creator"

	modifier = "modifier"
	address  = "test"
	is_admin  = true

	token = "string token"
)

var (
	authService    auth.Service
	authServiceMock    authMock.Service
	authRepository authMock.Repository
	userService user.Service
	adminService admin.Service
	adminRepo   adminMock.Repository
	userRepo   userMock.Repository
	cfg        config.Config

	data user.User
	reqAdmin request.RegisterAdminRequest
	reqUser request.RegisterUserRequest
	userData       user.User
	insertUserData user.InsertUserSpec
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestLogin(t *testing.T) {
	t.Run("Expect login the user", func(t *testing.T) {
		authServiceMock.On("Login",   mock.AnythingOfType("string"), mock.AnythingOfType("bool")).Return(&data, nil).Once()


		user, err := authService.Login(username, is_admin)

		assert.Nil(t, err)
		assert.NotNil(t, user)

		assert.Equal(t, name, user.Name)
		assert.Equal(t, username, user.Username)
		assert.Equal(t, password, user.Password)

	})

	t.Run("Expect user not found", func(t *testing.T) {
		authServiceMock.On("Login",   mock.Anything).Return(nil, business.ErrNotFound).Once()


		user, err := authServiceMock.Login(username, is_admin)

		assert.NotNil(t, err)

		assert.Nil(t, user)

		assert.Equal(t, err, business.ErrNotFound)
	})
}

func TestRegisterAdmin(t *testing.T) {
	t.Run("Expect register the user", func(t *testing.T) {
		authServiceMock.On("RegisterAdmin",   mock.Anything).Return(&reqAdmin, nil).Once()
		authRepository.On("InsertAdmin",   mock.AnythingOfType("admin.Admin")).Return(&reqAdmin, nil).Once()

		user, err := authService.RegisterAdmin(reqAdmin)

		assert.Nil(t, err)
		assert.NotNil(t, user)

		assert.Equal(t, name, user.Name)
		assert.Equal(t, username, user.Username)
		assert.Equal(t, password, user.Password)
		assert.Equal(t, repassword, user.RePassword)

	})

	t.Run("Expect user not found", func(t *testing.T) {
		authServiceMock.On("RegisterAdmin",   mock.Anything).Return(nil, business.ErrNotFound).Once()


		user, err := authServiceMock.RegisterAdmin(reqAdmin)

		assert.NotNil(t, err)

		assert.Nil(t, user)

		assert.Equal(t, err, business.ErrNotFound)
	})

	t.Run("Expect user different password and re password", func(t *testing.T) {
		authServiceMock.On("RegisterAdmin",   mock.Anything).Return(&reqAdmin, nil).Once()

		user, err := authServiceMock.RegisterAdmin(reqAdmin)

		assert.Nil(t, err)
		assert.NotNil(t, user)

		assert.Equal(t, name, user.Name)
		assert.Equal(t, username, user.Username)
		assert.Equal(t, password, user.Password)
		assert.NotEqual(t, username, user.RePassword)
	})
}

func TestRegisterUser(t *testing.T) {
	t.Run("Expect register the user", func(t *testing.T) {
		authServiceMock.On("RegisterUser",   mock.Anything).Return(&reqUser, nil).Once()

		user, err := authServiceMock.RegisterUser(reqUser)

		assert.Nil(t, err)
		assert.NotNil(t, user)

		assert.Equal(t, name, user.Name)
		assert.Equal(t, username, user.Username)
		assert.Equal(t, password, user.Password)
		assert.Equal(t, repassword, user.RePassword)
		assert.Equal(t, address, user.Address)



	})

	t.Run("Expect user not found", func(t *testing.T) {
		authServiceMock.On("RegisterUser",   mock.Anything).Return(nil, business.ErrNotFound).Once()


		user, err := authServiceMock.RegisterUser(reqUser)

		assert.NotNil(t, err)

		assert.Nil(t, user)

		assert.Equal(t, err, business.ErrNotFound)
	})

	t.Run("Expect user different password and re password", func(t *testing.T) {
		authServiceMock.On("RegisterUser",   mock.Anything).Return(&reqUser, nil).Once()

		user, err := authServiceMock.RegisterUser(reqUser)

		assert.Nil(t, err)
		assert.NotNil(t, user)

		assert.Equal(t, name, user.Name)
		assert.Equal(t, username, user.Username)
		assert.Equal(t, password, user.Password)
		assert.NotEqual(t, username, user.RePassword)
		assert.Equal(t, address, user.Address)
	})
}

func setup() {
	reqAdmin = request.RegisterAdminRequest{
		Name:       name,
		Username:   username,
		Password:   password,
		RePassword: repassword,
	}

	reqUser = request.RegisterUserRequest{
		Name:       name,
		Username:   username,
		Password:   password,
		RePassword: password,
		Address: address,
	}

	data = user.User{
		ID:         id,
		Name:       name,
		Username:   username,
		Password:   password,
		Address:    "",
		CreatedAt:  time.Now(),
		CreatedBy:  creator,
		Version:    1,
	}

	userData = user.NewUser(
		name,
		username,
		password,
		"admin",
		time.Now(),
	)

	insertUserData = user.InsertUserSpec{
		Name:     name,
		Username: username,
		Password: password,
	}

	authService = auth.NewService(userService, adminService, &adminRepo, &userRepo, cfg )
}
