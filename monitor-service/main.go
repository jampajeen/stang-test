package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/jampajeen/stang-test/monitor-service/api"
	"github.com/jampajeen/stang-test/monitor-service/core"
	"github.com/jampajeen/stang-test/monitor-service/ethereum"
)

var wg sync.WaitGroup

func main() {

	db := bootstrap()
	listener := ethereum.NewTransactionListener(db)

	wg.Add(1)
	go listener.Listen()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, echo.HeaderXRequestedWith},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))

	apiController := api.NewApiController(db)
	e.GET("/api/addresses/:id", apiController.GetTransactionByAddressHandler)

	port := fmt.Sprintf(":%d", core.Config.APP.BindPort)
	e.Logger.Fatal(e.Start(port))

	wg.Wait()
}
