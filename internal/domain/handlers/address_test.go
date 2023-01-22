package handlers_test

import (
	"github.com/stretchr/testify/assert"
	httphdl "github.com/victoorraphael/coordinator/cmd/http"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddressHandler_Find(t *testing.T) {
	router := httphdl.Routes(srvs, true)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/address", nil)
	router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, 200, w.Code)
}
