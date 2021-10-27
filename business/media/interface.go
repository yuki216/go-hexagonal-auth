package media

import (
	"mime/multipart"
)

//Service outgoing port for user
type Service interface {
	//Login If data not found will return nil without error
	UploadMedia(file *multipart.FileHeader) (string, error)
}

