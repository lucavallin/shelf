package api

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"github.com/lucavallin/yak-webshop/pkg/herd"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
)

// PostLoadHandler handlers POST to /yak-webshop/load
func (app *App) PostLoadHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		xmlContent, _ := ioutil.ReadAll(c.Request.Body)

		tmpHerd := &herd.Herd{}
		err := xml.Unmarshal(xmlContent, &tmpHerd)
		if err != nil || tmpHerd.Yaks == nil {
			c.Status(http.StatusBadRequest)
			return
		}

		app.herdRepo.Save(tmpHerd)
		c.Status(http.StatusResetContent)
	}
}

// GetStockHandler handlers GET to /yak-webshop/stock/:days
func (app *App) GetStockHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		// Maybe validate elapsedDays for > 0
		elapsedDays, _ := strconv.Atoi(c.Param("days"))
		newHerd := app.herdRepo.Get()
		stock := newHerd.GetStock(elapsedDays)
		stock.Milk = math.Round(stock.Milk / 0.01) * 0.01

		c.JSON(http.StatusOK, stock)
	}
}

// GetHerdHandler handlers GET to /yak-webshop/herd/:days
func (app *App) GetHerdHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		elapsedDays, _ := strconv.Atoi(c.Param("days"))
		newHerd := app.herdRepo.Get()
		newHerd.Age(elapsedDays)

		c.JSON(http.StatusOK, newHerd)
	}
}