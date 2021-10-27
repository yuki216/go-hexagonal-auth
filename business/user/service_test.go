package user_test

import (
	"go-hexagonal-auth/business"
	"go-hexagonal-auth/business/user"
	userMock "go-hexagonal-auth/business/user/mocks"

	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	id       = 1
	name     = "name"
	username = "username"
	password = "password"
	creator  = "creator"

	modifier = "modifier"
	version  = 1
)

var (
	userService    user.Service
	userRepository userMock.Repository

	data user.User
	userData       user.User
	insertUserData user.InsertUserSpec
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindUserByID(t *testing.T) {
	t.Run("Expect found the user", func(t *testing.T) {
		userRepository.On("FindUserByID",  mock.AnythingOfType("int")).Return(&data, nil).Once()

		user, err := userService.FindUserByID(id)
		assert.Nil(t, err)

		assert.NotNil(t, user)

		assert.Equal(t, id, user.ID)
		assert.Equal(t, name, user.Name)
		assert.Equal(t, username, user.Username)
		assert.Equal(t, password, user.Password)

	})

	t.Run("Expect user not found", func(t *testing.T) {
		userRepository.On("FindUserByID",  mock.AnythingOfType("int")).Return(nil, business.ErrNotFound).Once()

		user, err := userService.FindUserByID(id)

		assert.NotNil(t, err)

		assert.Nil(t, user)

		assert.Equal(t, err, business.ErrNotFound)
	})
}

func TestInsertUserByID(t *testing.T) {
	t.Run("Expect insert user success", func(t *testing.T) {
		userRepository.On("InsertUser", mock.AnythingOfType("user.User"), mock.AnythingOfType("string")).Return(nil).Once()

		err := userService.InsertUser(insertUserData, creator)

		assert.Nil(t, err)

	})

	t.Run("Expect insert user not found", func(t *testing.T) {
		userRepository.On("InsertUser", mock.AnythingOfType("user.User"), mock.AnythingOfType("string")).Return(business.ErrInternalServerError).Once()

		err := userService.InsertUser(insertUserData, creator)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInternalServerError)
	})
}

func TestUpdateUserByID(t *testing.T) {
	t.Run("Expect update user success", func(t *testing.T) {
		userRepository.On("FindUserByID", mock.AnythingOfType("int")).Return(&userData, nil).Once()
		userRepository.On("UpdateUser", mock.AnythingOfType("user.User"), mock.AnythingOfType("int")).Return(nil).Once()

		err := userService.UpdateUser(id, name, modifier, version)

		assert.Nil(t, err)

	})

	t.Run("Expect update user failed", func(t *testing.T) {
		userRepository.On("FindUserByID", mock.AnythingOfType("int")).Return(&userData, nil).Once()
		userRepository.On("UpdateUser", mock.AnythingOfType("user.User"), mock.AnythingOfType("int")).Return(business.ErrInternalServerError).Once()

		err := userService.UpdateUser(id, name, modifier, version)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInternalServerError)
	})
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
		creator,
		time.Now(),
	)

	insertUserData = user.InsertUserSpec{
		Name:     name,
		Username: username,
		Password: password,
	}

	userService = user.NewService(&userRepository)
}
