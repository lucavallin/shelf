package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lucavallin/yak-webshop/pkg/herd"
	"github.com/prometheus/common/log"
	"net/http"
)

// App contains all the dependencies needed for the API
type App struct {
	herdRepo herd.Repository
	router   *gin.Engine
}

// NewApp creates a new app with the needed dependencies
func NewApp(herdRepo herd.Repository) *App {
	router := gin.Default()
	router.Use(cors.Default())

	app := &App{herdRepo, router}
	app.router.GET("/ping", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	yakWebShop := app.router.Group("/yak-webshop")
	{
		yakWebShop.POST("/load", app.PostLoadHandler())
		yakWebShop.GET("/stock/:days", app.GetStockHandler())
		yakWebShop.GET("/herd/:days", app.GetHerdHandler())
		yakWebShop.POST("/order/:days", app.PostOrderHandler())
	}

	return app
}

// Run starts the APIs
func (app *App) Run(port string) {
	log.Fatal(app.router.Run(":" + port))
}
