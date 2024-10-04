package helper

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/muhammadsarimin/simple-api-xmu/types"
	"gorm.io/gorm"
)

func Error(c *gin.Context, err error) {

	var customErr *types.CustomError
	var jsonErr *json.UnmarshalTypeError

	if errors.As(err, &customErr) {
		c.JSON(400, types.Response{
			ResponseCode:    customErr.Code,
			ResponseMessage: customErr.Message,
		})
		return
	}

	if errors.As(err, &jsonErr) {
		c.JSON(400, types.Response{
			ResponseCode:    "400",
			ResponseMessage: jsonErr.Error(),
		})
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(404, types.Response{
			ResponseCode:    "404",
			ResponseMessage: "movie not found",
		})
		return
	}

	c.JSON(500, types.Response{
		ResponseCode:    "500",
		ResponseMessage: "General Error",
	})
}
