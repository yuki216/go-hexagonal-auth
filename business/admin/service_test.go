package admin_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-hexagonal-auth/business"
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

func TestFindUserByID(t *testing.T) {
	t.Run("Expect found the user", func(t *testing.T) {
		adminRepository.On("FindAdminByID", mock.AnythingOfType("int")).Return(&userData, nil).Once()

		user, err := adminService.FindAdminByID(id)

		assert.Nil(t, err)

		assert.NotNil(t, user)

		assert.Equal(t, id, user.ID)
		assert.Equal(t, name, user.Name)
		assert.Equal(t, username, user.Username)
		assert.Equal(t, password, user.Password)

	})

	t.Run("Expect user not found", func(t *testing.T) {
		adminRepository.On("FindAdminByID", mock.AnythingOfType("int")).Return(nil, business.ErrNotFound).Once()

		user, err := adminService.FindAdminByID(id)

		assert.NotNil(t, err)

		assert.Nil(t, user)

		assert.Equal(t, err, business.ErrNotFound)
	})
}
func TestFindUserByName(t *testing.T) {
	t.Run("Expect found the user", func(t *testing.T) {
		adminRepository.On("FindAdminByAdminnameAndPassword", mock.AnythingOfType("string"),mock.AnythingOfType("string")).Return(&userData, nil).Once()

		user, err := adminService.FindAdminByAdminnameAndPassword(username, password)

		assert.Nil(t, err)

		assert.NotNil(t, user)

		assert.Equal(t, id, user.ID)
		assert.Equal(t, name, user.Name)
		assert.Equal(t, username, user.Username)
		assert.Equal(t, password, user.Password)

	})

	t.Run("Expect user not found", func(t *testing.T) {
		adminRepository.On("FindAdminByAdminnameAndPassword", mock.AnythingOfType("string"),mock.AnythingOfType("string")).Return(nil, business.ErrNotFound).Once()

		user, err := adminService.FindAdminByAdminnameAndPassword(username, password)

		assert.NotNil(t, err)

		assert.Nil(t, user)

		assert.Equal(t, err, business.ErrNotFound)
	})
}

func TestInsertAdminByID(t *testing.T) {
	t.Run("Expect insert admin success", func(t *testing.T) {
		adminRepository.On("InsertAdmin", mock.AnythingOfType("admin.Admin"), mock.AnythingOfType("string")).Return(nil).Once()

		err := adminService.InsertAdmin(insertUserData, creator)

		assert.Nil(t, err)

	})

	t.Run("Expect insert user not found", func(t *testing.T) {
		adminRepository.On("InsertAdmin", mock.AnythingOfType("admin.Admin"), mock.AnythingOfType("string")).Return(business.ErrInternalServerError).Once()

		err := adminService.InsertAdmin(insertUserData, creator)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInternalServerError)
	})
}

func TestUpdateAdminByID(t *testing.T) {
	t.Run("Expect update user success", func(t *testing.T) {
		adminRepository.On("FindAdminByID", mock.AnythingOfType("int")).Return(&userData, nil).Once()
		adminRepository.On("UpdateAdmin", mock.AnythingOfType("admin.Admin"), mock.AnythingOfType("string")).Return(nil).Once()

		err := adminService.UpdateAdmin(id, name, modifier)

		assert.Nil(t, err)

	})

	t.Run("Expect update user failed", func(t *testing.T) {
		adminRepository.On("FindAdminByID", mock.AnythingOfType("int")).Return(&userData, nil).Once()
		adminRepository.On("UpdateAdmin", mock.AnythingOfType("admin.Admin"), mock.AnythingOfType("string")).Return(business.ErrInternalServerError).Once()

		err := adminService.UpdateAdmin(id, name, modifier)

		assert.NotNil(t, err)

		assert.Equal(t, err, business.ErrInternalServerError)
	})
}

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
