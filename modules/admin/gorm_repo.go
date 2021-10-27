package admin

import (
	"go-hexagonal-auth/business/admin"
	"time"

	"gorm.io/gorm"
)

//GormRepository The implementation of Admin.Repository object
type GormRepository struct {
	DB *gorm.DB
}

type Admins struct {
	ID         int       `gorm:"id;primaryKey;autoIncrement"`
	Name       string    `gorm:"name"`
	Adminname   string    `gorm:"email;index:idx_email,unique"`
	Password   string    `gorm:"password"`
	CreatedAt  time.Time `gorm:"created_at"`
	CreatedBy  string    `gorm:"created_by"`
	ModifiedAt time.Time `gorm:"modified_at"`
	ModifiedBy string    `gorm:"modified_by"`
}

func newAdmin(Admin admin.Admin) *Admins {

	return &Admins{
		Admin.ID,
		Admin.Name,
		Admin.Username,
		Admin.Password,
		Admin.CreatedAt,
		Admin.CreatedBy,
		Admin.ModifiedAt,
		Admin.ModifiedBy,
	}

}

func (col *Admins) ToAdmin() admin.Admin {
	var Admin admin.Admin

	Admin.ID = col.ID
	Admin.Name = col.Name
	Admin.Username = col.Adminname
	Admin.Password = col.Password
	Admin.CreatedAt = col.CreatedAt
	Admin.CreatedBy = col.CreatedBy
	Admin.ModifiedAt = col.ModifiedAt
	Admin.ModifiedBy = col.ModifiedBy

	return Admin
}

//NewGormDBRepository Generate Gorm DB Admin repository
func NewGormDBRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

//FindAdminByID If data not found will return nil without error
func (repo *GormRepository) FindAdminByID(id int) (*admin.Admin, error) {

	var AdminData Admins

	err := repo.DB.First(&AdminData, id).Error
	if err != nil {
		return nil, err
	}

	Admin := AdminData.ToAdmin()

	return &Admin, nil
}

//FindAdminByID If data not found will return nil without error
func (repo *GormRepository) FindAdminByAdminnameAndPassword(Adminname string, password string) (*admin.Admin, error) {

	var AdminData Admins

	err := repo.DB.Where("Adminname = ?", Adminname).First(&AdminData).Error
	if err != nil {
		return nil, err
	}

	Admin := AdminData.ToAdmin()

	return &Admin, nil
}

//FindAllAdmin find all Admin with given specific page and row per page, will return empty slice instead of nil
func (repo *GormRepository) FindAllAdmin(skip int, rowPerPage int) ([]admin.Admin, error) {

	var Admins []Admins

	err := repo.DB.Offset(skip).Limit(rowPerPage).Find(&Admins).Error
	if err != nil {
		return nil, err
	}

	var result []admin.Admin
	for _, value := range Admins {
		result = append(result, value.ToAdmin())
	}

	return result, nil
}

//InsertAdmin Insert new Admin into storage
func (repo *GormRepository) InsertAdmin(admin admin.Admin) error {

	AdminData := newAdmin(admin)
	AdminData.ID = 0

	err := repo.DB.Create(AdminData).Error
	if err != nil {
		return err
	}

	return nil
}

//UpdateItem Update existing item in database
func (repo *GormRepository) UpdateAdmin(admin admin.Admin, currentVersion int) error {

	AdminData := newAdmin(admin)

	err := repo.DB.Model(&AdminData).Updates(Admins{Name: AdminData.Name}).Error
	if err != nil {
		return err
	}

	return nil
}
