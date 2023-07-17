package api

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/jampajeen/stang-test/monitor-service/db"
)

type ApiController struct {
	db *db.MongoDb
}

func NewApiController(db *db.MongoDb) *ApiController {
	return &ApiController{db: db}
}

func (a *ApiController) GetTransactionByAddressHandler(c echo.Context) error {

	id := c.Param("id")

	res, err := a.db.QueryTransactionRecord(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}
