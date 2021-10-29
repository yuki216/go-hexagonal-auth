package auth

import (
	"go-hexagonal-auth/api/v1/auth/request"
	"go-hexagonal-auth/business"
	"go-hexagonal-auth/business/admin"
	"go-hexagonal-auth/business/user"
	"go-hexagonal-auth/config"
	utils "go-hexagonal-auth/util"
	"go-hexagonal-auth/util/validator"
)

//=============== The implementation of those interface put below =======================
type service struct {
	userService user.Service
	adminService admin.Service
	adminRepo   admin.Repository
	userRepo   user.Repository
	cfg        config.Config
}

//NewService Construct user service object
func NewService(userService user.Service, adminService admin.Service, adminRepo admin.Repository, userRepo user.Repository, cfg config.Config) Service {
	user := &service{
		userService,
		adminService,
		adminRepo,
		userRepo,
		cfg,
	}
	return user
}

//Login by given user Username and Password, return error if not exist
func (s *service) Login(username string, isAdmin bool) (*user.User, error) {
	var result user.User
	if isAdmin {
		adminData, err := s.adminService.FindAdminByAdminnameAndPassword(username, "")
		if err != nil {
			return nil, err
		}

		result = user.User{
			Name:       adminData.Name,
			Username:   adminData.Username,
			Password:   adminData.Password,
		}


	} else{
		userData, err := s.userService.FindUserByUsernameAndPassword(username, "")
		if err != nil {
			return nil, err
		}

		result = user.User{
			Name:       userData.Name,
			Username:   userData.Username,
			Password:   userData.Password,
		}
	}




	return &result, nil
}

func (s *service) RegisterAdmin(request request.RegisterAdminRequest) (*request.RegisterAdminRequest, error)  {

	err := validator.GetValidator().Struct(request)
	if err != nil {
		return nil,business.ErrInvalidSpec
	}
    pass, err :=utils.HashPassword(request.Password)
	if err != nil {
		return nil,business.ErrInvalidSpec
	}
	AdminReq := admin.Admin{
		Name:       request.Name,
		Username:   request.Username,
		Password:   pass,
	}
	err = s.adminRepo.InsertAdmin(AdminReq)
	if err != nil {
		return nil,err
	}

	return &request,nil
}

func (s *service) RegisterUser(request request.RegisterUserRequest) (*request.RegisterUserRequest, error)  {
	err := validator.GetValidator().Struct(request)
	if err != nil {
		return nil,business.ErrInvalidSpec
	}

	pass, err :=utils.HashPassword(request.Password)
	if err != nil {
		return nil,business.ErrInvalidSpec
	}

	UserReq := user.User{
		Name:       request.Name,
		Username:   request.Username,
		Password:   pass,
		Address: request.Address,
	}
	err = s.userRepo.InsertUser(UserReq)
	if err != nil {
		return nil,err
	}

	return &request,nil
}