package auth_test

import (
	"fmt"
	"github.com/stretchr/testify/mock"
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
	creator  = "creator"

	modifier = "modifier"
	version  = 1
	is_admin  = false

	token = "string token"
)

var (
	authService    auth.Service
	authRepository authMock.Repository
	userService user.Service
	adminService admin.Service
	adminRepo   adminMock.Repository
	userRepo   userMock.Repository
	cfg        config.Config

	data user.User
	userData       user.User
	insertUserData user.InsertUserSpec
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestLogin(t *testing.T) {
	t.Run("Expect login the user", func(t *testing.T) {
		authRepository.On("FindAdminByAdminnameAndPassword",  mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&data, nil).Once()

		user, err := authService.Login(username, is_admin)
		fmt.Println(err)
		assert.Nil(t, err)
		//user := user.User{
		//	Name:       "name",
		//	Username:   "username",
		//	Password:   "password",
		//}

		assert.NotNil(t, user)

		assert.Equal(t, name, user.Name)
		assert.Equal(t, username, user.Username)
		assert.Equal(t, password, user.Password)

	})

	//t.Run("Expect user not found", func(t *testing.T) {
	//	userRepository.On("FindUserByID",  mock.AnythingOfType("int")).Return(nil, business.ErrNotFound).Once()
	//
	//	user, err := userService.FindUserByID(id)
	//
	//	assert.NotNil(t, err)
	//
	//	assert.Nil(t, user)
	//
	//	assert.Equal(t, err, business.ErrNotFound)
	//})
}

func setup() {

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
