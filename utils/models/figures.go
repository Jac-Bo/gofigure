package figures

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AFigure struct {
	gorm.Model
	Title       string  `json:"title"`
	Value       float64 `json:"valuea"`
	SecondValue float64 `json:"valueb"`
	FinalVal    float64 `json:"finalvalue"`
}

func (fig AFigure) Multiply(a float64, b float64, c *gin.Context) {

	funcFinalFigure := fig.Value / fig.SecondValue

	c.JSON(http.StatusAccepted, funcFinalFigure)
}

func (fig AFigure) SumOf(a float64, b float64, c *gin.Context) {
	funcFinalFigure := fig.Value + fig.SecondValue
	c.JSON(http.StatusAccepted, funcFinalFigure)
}

func (fig AFigure) Divide(a float64, b float64, c *gin.Context) {
	if fig.Value == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Divided By Zero",
		})
	}
	funcFinalFigure := fig.Value / fig.SecondValue
	fig.FinalVal = funcFinalFigure
	c.JSON(http.StatusAccepted, fig.FinalVal)
}
