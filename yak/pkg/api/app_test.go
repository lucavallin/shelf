package api

import (
	"github.com/lucavallin/yak-webshop/pkg/herd"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// make sure the app is alive
func TestApp_Run(t *testing.T) {
	herdRepo := herd.NewXMLFileRepository("../../data/herd.xml")
	app := NewApp(herdRepo)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	app.router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}