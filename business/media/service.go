package media

import (
	"fmt"
	"go-hexagonal-auth/config"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

//=============== The implementation of those interface put below =======================
type service struct {
	cfg        config.Config
}

//NewService Construct user service object
func NewService( cfg config.Config) Service {
	return &service{
		cfg,
	}
}

//Login by given user Username and Password, return error if not exist
func (s *service) UploadMedia(file *multipart.FileHeader) (string, error) {
	var filename =""

	currentTime := time.Now()

	alias := currentTime.Format("20060102150405000000")+"_image"

	uploadedFile, err := file.Open()
	if err != nil {
		return "", err
	}
	defer uploadedFile.Close()

	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	filename = file.Filename
	if alias != "" {
		filename = fmt.Sprintf("%s%s", alias, filepath.Ext(file.Filename))
	}
	folder := "public/products"
	if _, err := os.Stat(dir+"/public/products/"+currentTime.Format("2006-01-02")); os.IsNotExist(err) {
		err := os.Mkdir(dir+"//public//products//"+currentTime.Format("2006-01-02"), os.ModePerm)
		if err != nil {
			return "", err
		}
		folder = folder+"/"+currentTime.Format("2006-01-02")
	}

	fileLocation := filepath.Join(dir, folder, filename)
	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	defer targetFile.Close()

	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
		return "", err
	}

	return "/products/"+currentTime.Format("2006-01-02")+"/"+filename, nil
}
