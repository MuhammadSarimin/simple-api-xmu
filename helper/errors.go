package helper

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhammadsarimin/simple-api-xmu/types"
	"gorm.io/gorm"
)

func Error(c *gin.Context, err error) {

	var customErr *types.CustomError
	var jsonErr *json.UnmarshalTypeError
	var strErr *strconv.NumError

	if errors.As(err, &customErr) {
		c.JSON(400, types.Response{
			ResponseCode:    customErr.Code,
			ResponseMessage: customErr.Message,
		})
		return
	}

	if errors.As(err, &jsonErr) {
		c.JSON(400, types.Response{
			ResponseCode:    "002",
			ResponseMessage: jsonErr.Field + " must be " + jsonErr.Type.String(),
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

	if errors.As(err, &strErr) {
		c.JSON(400, types.Response{
			ResponseCode:    "002",
			ResponseMessage: "id must be an integer",
		})
		return
	}

	c.JSON(500, types.Response{
		ResponseCode:    "500",
		ResponseMessage: "General Error",
	})
}
