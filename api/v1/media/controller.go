package media

import (
	"github.com/labstack/echo/v4"
	"go-hexagonal-auth/api/common"
	"go-hexagonal-auth/api/v1/media/response"
	"go-hexagonal-auth/business/media"
	"go-hexagonal-auth/config"
)

//Controller Get item API controller
type Controller struct {
	service media.Service
	cfg config.Config
}

//NewController Construct item API controller
func NewController(service media.Service, cfg config.Config) *Controller {
	return &Controller{
		service,
		cfg,
	}
}

//Login by given username and password will return JWT token
func (controller *Controller) MediaUpload(c echo.Context) error {



	form, err := c.FormFile("file")
	if err != nil {
		return c.JSON(common.NewErrorResponse(common.ControllerResponse{
			Code:    common.ErrBadRequest,
			Message: err.Error(),
			Data:    nil,
		}))
	}

	filename, err := controller.service.UploadMedia(form)

	if err != nil {
		return c.JSON(common.NewErrorResponse(common.ControllerResponse{
			Code:    common.ErrBadRequest,
			Message: err.Error(),
			Data:    nil,
		}))
	}
	response :=  response.NewMediaResponse(filename)

	return c.JSON(common.NewSuccessResponse(response))

}