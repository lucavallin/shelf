package api

import (
	"github.com/lucavallin/yak-webshop/pkg/herd"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	herdRepo = herd.NewXMLFileRepository("../../data/herd_test.xml")
	app = NewApp(herdRepo)
)

func TestApp_PostLoadHandler(t *testing.T) {
	requestBody := `
<herd>
<labyak name="Betty-1" age="4" sex="f"/>
<labyak name="Betty-2" age="8" sex="f"/>
<labyak name="Betty-3" age="9.5" sex="f"/>
</herd>
`

	req, _ := http.NewRequest("POST", "/yak-webshop/load", strings.NewReader(requestBody))
	w1 := httptest.NewRecorder()
	app.router.ServeHTTP(w1, req)
	// Returns 200 on correct request
	assert.Equal(t, http.StatusResetContent, w1.Code)

	req, _ = http.NewRequest("POST", "/yak-webshop/load", strings.NewReader(""))
	w2 := httptest.NewRecorder()
	app.router.ServeHTTP(w2, req)
	// Returns 400 on broken xml
	assert.Equal(t, http.StatusBadRequest, w2.Code)
}

func TestApp_GetStockHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/yak-webshop/stock/13", nil)
	w := httptest.NewRecorder()
	app.router.ServeHTTP(w, req)
	// Returns 200 on correct request
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"milk\":1104.48,\"wool\":3}", w.Body.String())
}

func TestApp_GetHerdHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/yak-webshop/herd/13", nil)
	w := httptest.NewRecorder()
	app.router.ServeHTTP(w, req)
	// Returns 200 on correct request
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t,
		"{\"herd\":[{\"name\":\"Betty-1\",\"age\":4.13,\"age-last-shaved\":4},{\"name\":\"Betty-2\",\"age\":8.13,\"age-last-shaved\":8},{\"name\":\"Betty-3\",\"age\":9.63,\"age-last-shaved\":9.5}]}",
		w.Body.String())
}