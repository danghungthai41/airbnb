package placehttp

import (
	placemodel "airbnb-golang/internal/places/model"
	"airbnb-golang/pkg/common"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IPlaceUsecase interface {
	CreatePlace(context.Context, *placemodel.Place) (*placemodel.Place, error)
	GetPlaces(context.Context, *common.Paging, *placemodel.Filter) ([]placemodel.Place, error)
	GetPlaceByID(context.Context, int) (*placemodel.Place, error)
	DeletePlace(context.Context, int) error
	UpdatePlace(context.Context, int, *placemodel.Place) error
}

type placeHandler struct {
	placeUC IPlaceUsecase
}

func NewPlaceHandler(placeuc IPlaceUsecase) *placeHandler {
	return &placeHandler{placeuc}

}
func (hdl *placeHandler) CreatePlace() gin.HandlerFunc {
	return func(c *gin.Context) {
		var place placemodel.Place
		if err := c.ShouldBind(&place); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		data, err := hdl.placeUC.CreatePlace(c.Request.Context(), &place)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"data": data})
	}
}

func (hdl *placeHandler) GetPlaces() gin.HandlerFunc {
	return func(c *gin.Context) {
		var paging common.Paging
		var filter placemodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		paging.FullFill()
		data, err := hdl.placeUC.GetPlaces(c.Request.Context(), &paging, &filter)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"data": data, "paging": paging})
	}
}
func (hdl *placeHandler) GetPlaceByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		fmt.Println("id", id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		data, err := hdl.placeUC.GetPlaceByID(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"data": data})

	}
}

func (hdl *placeHandler) DeletePlace() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return

		}
		if err := hdl.placeUC.DeletePlace(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"data": true})

	}
}

func (hdl *placeHandler) UpdatePlace() gin.HandlerFunc {
	return func(c *gin.Context) {
		var place placemodel.Place
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := c.ShouldBind(&place); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := hdl.placeUC.UpdatePlace(c.Request.Context(), id, &place); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"data": place})

	}

}
