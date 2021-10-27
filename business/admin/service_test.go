package admin_test

import (
	"go-hexagonal-auth/business/admin"
	userMock "go-hexagonal-auth/business/admin/mocks"

	"os"
	"testing"
	"time"
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
	adminService    admin.Service
	adminRepository userMock.Repository

	userData       admin.Admin
	insertUserData admin.InsertAdminSpec
)

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

//func TestFindUserByID(t *testing.T) {
//	t.Run("Expect found the user", func(t *testing.T) {
//		adminRepository.On("FindAdminByID", mock.AnythingOfType("string")).Return(&userData, nil).Once()
//
//		user, err := adminService.FindAdminByID(id)
//
//		assert.Nil(t, err)
//
//		assert.NotNil(t, user)
//
//		assert.Equal(t, id, user.ID)
//		assert.Equal(t, name, user.Name)
//		assert.Equal(t, username, user.Username)
//		assert.Equal(t, password, user.Password)
//
//	})
//
//	t.Run("Expect user not found", func(t *testing.T) {
//		adminRepository.On("FindAdminByID", mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()
//
//		user, err := adminService.FindAdminByID(id)
//
//		assert.NotNil(t, err)
//
//		assert.Nil(t, user)
//
//		assert.Equal(t, err, business.ErrNotFound)
//	})
//}
//
//func TestInsertUserByID(t *testing.T) {
//	t.Run("Expect insert user success", func(t *testing.T) {
//		adminRepository.On("InsertUser", mock.AnythingOfType("user.User"), mock.AnythingOfType("string")).Return(nil).Once()
//
//		err := adminService.InsertAdmin(insertUserData, creator)
//
//		assert.Nil(t, err)
//
//	})
//
//	t.Run("Expect insert user not found", func(t *testing.T) {
//		adminRepository.On("InsertUser", mock.AnythingOfType("user.User"), mock.AnythingOfType("string")).Return(business.ErrInternalServerError).Once()
//
//		err := adminService.InsertAdmin(insertUserData, creator)
//
//		assert.NotNil(t, err)
//
//		assert.Equal(t, err, business.ErrInternalServerError)
//	})
//}
//
//func TestUpdateUserByID(t *testing.T) {
//	t.Run("Expect update user success", func(t *testing.T) {
//		adminRepository.On("FindAdminByID", mock.AnythingOfType("string")).Return(&userData, nil).Once()
//		adminRepository.On("UpdateAdmin", mock.AnythingOfType("user.User"), mock.AnythingOfType("int")).Return(nil).Once()
//
//		err := adminService.UpdateAdmin(id, name, modifier, version)
//
//		assert.Nil(t, err)
//
//	})
//
//	t.Run("Expect update user failed", func(t *testing.T) {
//		adminRepository.On("FindAdminByID", mock.AnythingOfType("string")).Return(&userData, nil).Once()
//		adminRepository.On("UpdateAdmin", mock.AnythingOfType("user.User"), mock.AnythingOfType("int")).Return(business.ErrInternalServerError).Once()
//
//		err := adminService.UpdateAdmin(id, name, modifier, version)
//
//		assert.NotNil(t, err)
//
//		assert.Equal(t, err, business.ErrInternalServerError)
//	})
//}

func setup() {

	userData = admin.NewAdmin(
		id,
		name,
		username,
		password,
		creator,
		time.Now(),
	)

	insertUserData = admin.InsertAdminSpec{
		Name:     name,
		Adminname: username,
		Password: password,
	}

	adminService = admin.NewService(&adminRepository)
}
