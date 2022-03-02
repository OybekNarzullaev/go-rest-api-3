package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServerHandler(t *testing.T) {
	// serverni init qilamiz:
	s := New(NewConfig())

	// http response ni yozib olish uchun o'zgaruvchi
	rec := httptest.NewRecorder()

	// reuestni shakllantiramiz:
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	// api ishlatib ko'ramiz:
	s.HandleHello().ServeHTTP(rec, req)

	assert.Equal(t, rec.Body.String(), "hello")
}