package controllers

import (
	"net/http"

	"github.com/Jac-Bo/gofigure/utils/db"
	figures "github.com/Jac-Bo/gofigure/utils/models"
	"github.com/gin-gonic/gin"
)

func CreateFigure(c *gin.Context) {
	var figure *figures.AFigure // figure is then referenced when creating the response
	err := c.ShouldBind(&figure)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	response := db.DB.Create(figure) //response is created, as a var with the value of the created figure.
	if response.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating a figure",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"figure": figure,
	})

}
