package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lucavallin/yak-webshop/pkg/order"
	"io/ioutil"
	"net/http"
	"strconv"
)

// PostOrderHandler handlers POST to /yak-webshop/order/:days
func (app *App) PostOrderHandler() func(c *gin.Context) {
	book := order.NewBook(app.herdRepo.Get())

	return func(c *gin.Context) {
		// Maybe validate elapsedDays for > 0
		elapsedDays, _ := strconv.Atoi(c.Param("days"))
		body, _ := ioutil.ReadAll(c.Request.Body)
		newOrder := order.Order{}
		if err := json.Unmarshal(body, &newOrder); err != nil {
			c.Status(http.StatusBadRequest)
		}

		items := book.AddOrder(newOrder, elapsedDays)
		if items.Milk > 0 && items.Skins > 0 {
			c.JSON(http.StatusCreated, items)
			return
		} else if items.Milk == 0 && items.Skins == 0 {
			c.Status(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusPartialContent, items)
	}
}