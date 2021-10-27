package admin

//Service outgoing port for Admin
type Service interface {
	//FindAdminByID If data not found will return nil without error
	FindAdminByID(id int) (*Admin, error)

	//FindAdminByAdminnameAndPassword If data not found will return nil
	FindAdminByAdminnameAndPassword(username string, password string) (*Admin, error)

	//FindAllAdmin find all Admin with given specific page and row per page, will return empty slice instead of nil
	FindAllAdmin(skip int, rowPerPage int) ([]Admin, error)

	//InsertAdmin Insert new Admin into storage
	InsertAdmin(insertAdminSpec InsertAdminSpec, createdBy string) error

	//UpdateAdmin if data not found will return error
	UpdateAdmin(id int, name string, modifiedBy string) error
}

//Repository ingoing port for Admin
type Repository interface {
	//FindAdminByID If data not found will return nil without error
	FindAdminByID(id int) (*Admin, error)

	//FindAdminByAdminnameAndPassword If data not found will return nil
	FindAdminByAdminnameAndPassword(username string, password string) (*Admin, error)

	//FindAllAdmin find all Admin with given specific page and row per page, will return empty slice instead of nil
	FindAllAdmin(skip int, rowPerPage int) ([]Admin, error)

	//InsertAdmin Insert new Admin into storage
	InsertAdmin(admin Admin) error

	//UpdateAdmin if data not found will return error
	UpdateAdmin(admin Admin) error
}
