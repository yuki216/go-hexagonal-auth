package admin

import (
	"go-hexagonal-auth/business"
	"go-hexagonal-auth/util/validator"
	"time"
)

//InsertAdminSpec create Admin spec
type InsertAdminSpec struct {
	Name     string `validate:"required"`
	Adminname string `validate:"required"`
	Password string `validate:"required"`
}

//=============== The implementation of those interface put below =======================
type service struct {
	repository Repository
}

//NewService Construct Admin service object
func NewService(repository Repository) Service {
	return &service{
		repository,
	}
}

//FindAdminByID Get Admin by given ID, return nil if not exist
func (s *service) FindAdminByID(id int) (*Admin, error) {
	return s.repository.FindAdminByID(id)
}


func (s *service) FindAdminByAdminnameAndPassword(Adminname string, password string) (*Admin, error) {
	return s.repository.FindAdminByAdminnameAndPassword(Adminname, password)
}

//FindAllAdmin Get all Admins , will be return empty array if no data or error occured
func (s *service) FindAllAdmin(skip int, rowPerPage int) ([]Admin, error) {

	admin, err := s.repository.FindAllAdmin(skip, rowPerPage)
	if err != nil {
		return []Admin{}, err
	}

	return admin, err
}

//InsertAdmin Create new Admin and store into database
func (s *service) InsertAdmin(insertAdminSpec InsertAdminSpec, createdBy string) error {
	err := validator.GetValidator().Struct(insertAdminSpec)
	if err != nil {
		return business.ErrInvalidSpec
	}

	Admin := NewAdmin(
		1,
		insertAdminSpec.Name,
		insertAdminSpec.Adminname,
		insertAdminSpec.Password,
		createdBy,
		time.Now(),
	)

	err = s.repository.InsertAdmin(Admin)
	if err != nil {
		return err
	}

	return nil
}

//UpdateAdmin will update found Admin, if not exists will be return error
func (s *service) UpdateAdmin(id int, name string, modifiedBy string) error {

	Admin, err := s.repository.FindAdminByID(id)
	if err != nil {
		return err
	} else if Admin == nil {
		return business.ErrNotFound
	}

	modifiedAdmin := Admin.ModifyAdmin(name, time.Now(), modifiedBy)

	return s.repository.UpdateAdmin(modifiedAdmin)
}
