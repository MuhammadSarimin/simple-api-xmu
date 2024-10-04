package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muhammadsarimin/simple-api-xmu/helper"
	"github.com/muhammadsarimin/simple-api-xmu/types"
	log "github.com/sirupsen/logrus"
)

type handler struct {
	store Store
}

func NewHandler(r *gin.RouterGroup, s Store) {
	h := &handler{store: s}

	r.GET("/movies", h.FindAll)
	r.POST("/movies", h.Create)
	r.GET("/movies/:id", h.FindByID)
	r.PATCH("/movies/:id", h.Update)
	r.DELETE("/movies/:id", h.Delete)

}

func (h *handler) FindAll(c *gin.Context) {

	movies, err := h.store.FindAll()
	if err != nil {
		log.Error("Happened error when Find All", err)
		helper.Error(c, err)
		return
	}

	c.JSON(200, types.Response{
		ResponseCode:    "200",
		ResponseMessage: "Success",
		ResponseData:    movies,
	})

}

func (h *handler) Create(c *gin.Context) {

	var movie types.Movie

	if err := c.ShouldBindJSON(&movie); err != nil {
		log.Error("Happened error when bind json", err)
		helper.Error(c, err)
		return
	}

	if err := helper.Validate(movie); err != nil {
		log.Error("Happened error when Validate", err)
		helper.Error(c, err)
		return
	}

	if err := h.store.Create(&movie); err != nil {
		log.Error("Happened error when Create", err)
		helper.Error(c, err)
		return
	}

	c.JSON(200, types.Response{
		ResponseCode:    "200",
		ResponseMessage: "Success",
		ResponseData:    movie,
	})

}

func (h *handler) FindByID(c *gin.Context) {

	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Error("Happened error when convert string to int", err)
		helper.Error(c, err)
		return
	}

	movie, err := h.store.FindByID(id)
	if err != nil {
		log.Error("Happened error when Find By ID", err)
		helper.Error(c, err)
		return
	}

	c.JSON(200, types.Response{
		ResponseCode:    "200",
		ResponseMessage: "Success",
		ResponseData:    movie,
	})

}

func (h *handler) Update(c *gin.Context) {

	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Error("Happened error when convert string to int", err)
		helper.Error(c, err)
		return
	}

	var movie types.Movie

	if err := c.ShouldBindJSON(&movie); err != nil {
		log.Error("Happened error when bind json", err)
		helper.Error(c, err)
		return
	}

	if err := helper.Validate(movie); err != nil {
		log.Error("Happened error when Validate", err)
		helper.Error(c, err)
		return
	}

	movie.ID = uint(id)

	if err := h.store.Update(&movie); err != nil {
		log.Error("Happened error when Update", err)
		helper.Error(c, err)
		return
	}

	c.JSON(200, types.Response{
		ResponseCode:    "200",
		ResponseMessage: "Success",
		ResponseData:    movie,
	})

}

func (h *handler) Delete(c *gin.Context) {

	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Error("Happened error when convert string to int", err)
		helper.Error(c, err)
		return
	}

	if err := h.store.Delete(id); err != nil {
		log.Error("Happened error when Delete", err)
		helper.Error(c, err)
		return
	}

	c.JSON(200, types.Response{
		ResponseCode:    "200",
		ResponseMessage: "Success",
	})
}
