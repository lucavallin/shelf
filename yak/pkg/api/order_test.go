package api

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// When availability is ok
func TestApp_PostOrderHandler(t *testing.T) {
	// Overwrite app to load the orders book again
	app = NewApp(herdRepo)
	body := strings.NewReader(`{"customer":"Medvedev","order":{"milk":1100,"skins":3}}`)
	req, _ := http.NewRequest("POST", "/yak-webshop/order/13", body)

	w := httptest.NewRecorder()
	app.router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, `{"milk":1100,"skins":3}`, w.Body.String())
}

// When availability is not ok
func TestApp_PostOrderHandlerLowAvailability(t *testing.T) {
	// Overwrite app to load the orders book again
	app = NewApp(herdRepo)
	body := strings.NewReader(`{"customer":"Medvedev","order":{"milk":1400,"skins":6}}`)
	req, _ := http.NewRequest("POST", "/yak-webshop/order/13", body)

	w := httptest.NewRecorder()
	app.router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

// When availability is partial
func TestApp_PostOrderHandlerPartialAvailability(t *testing.T) {
	// Overwrite app to load the orders book again
	app = NewApp(herdRepo)
	body := strings.NewReader(`{"customer":"Medvedev","order":{"milk":1200,"skins":3}}`)
	req, _ := http.NewRequest("POST", "/yak-webshop/order/13", body)

	w := httptest.NewRecorder()
	app.router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusPartialContent, w.Code)
	assert.Equal(t, `{"skins":3}`, w.Body.String())
}

// When availability is partial
func TestApp_PostOrderHandlerAvailabilityOnSuccessiveDays(t *testing.T) {
	// Overwrite app to load the orders book again
	app = NewApp(herdRepo)

	body := strings.NewReader(`{"customer":"Medvedev","order":{"milk":1100,"skins":3}}`)
	req, _ := http.NewRequest("POST", "/yak-webshop/order/13", body)
	w := httptest.NewRecorder()
	app.router.ServeHTTP(w, req)

	body2 := strings.NewReader(`{"customer":"Medvedev","order":{"milk":1100,"skins":3}}`)
	req2, _ := http.NewRequest("POST", "/yak-webshop/order/14", body2)
	w2 := httptest.NewRecorder()
	app.router.ServeHTTP(w2, req2)

	assert.Equal(t, http.StatusNotFound, w2.Code)
}