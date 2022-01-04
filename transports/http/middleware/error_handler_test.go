package middleware_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/libsv/p4-server/log"
	"github.com/libsv/p4-server/transports/http/middleware"
	"github.com/stretchr/testify/assert"
	validator "github.com/theflyingcodr/govalidator"
	"github.com/theflyingcodr/lathos/errs"
)

func TestErrorHandler(t *testing.T) {
	tests := map[string]struct {
		err           error
		expResp       interface{}
		expStatusCode int
	}{
		"client error 400": {
			err: validator.ErrValidation{
				"paymentID": []string{"no style", "no class"},
			},
			expResp: map[string]interface{}{
				"errors": map[string]interface{}{
					"paymentID": []interface{}{"no style", "no class"},
				},
			},
			expStatusCode: http.StatusBadRequest,
		},
		"internal error 500": {
			err: errors.New("ahnah"),
			expResp: map[string]interface{}{
				"code":    "500",
				"title":   "Internal Server Error",
				"message": "ahnah",
			},
			expStatusCode: http.StatusInternalServerError,
		},
		"not found 404": {
			err: errs.NewErrNotFound("my 404", "not found"),
			expResp: map[string]interface{}{
				"code":    "my 404",
				"title":   "Not found",
				"message": "not found",
			},
			expStatusCode: http.StatusNotFound,
		},
		"conflict 409": {
			err: errs.NewErrDuplicate("my 409", "collision"),
			expResp: map[string]interface{}{
				"code":    "my 409",
				"title":   "Item already exists",
				"message": "collision",
			},
			expStatusCode: http.StatusConflict,
		},
		"not auth'd 401": {
			err: errs.NewErrNotAuthenticated("my 401", "will ya login"),
			expResp: map[string]interface{}{
				"code":    "my 401",
				"title":   "Not authenticated",
				"message": "will ya login",
			},
			expStatusCode: http.StatusUnauthorized,
		},
		"forbidden 403": {
			err: errs.NewErrNotAuthorised("my 403", "lol nice try buddy"),
			expResp: map[string]interface{}{
				"code":    "my 403",
				"title":   "Permission denied",
				"message": "lol nice try buddy",
			},
			expStatusCode: http.StatusForbidden,
		},
		"cannot process 422": {
			err: errs.NewErrUnprocessable("my 422", "what did you even send?"),
			expResp: map[string]interface{}{
				"code":    "my 422",
				"title":   "Unprocessable",
				"message": "what did you even send?",
			},
			expStatusCode: http.StatusUnprocessableEntity,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			rec := httptest.NewRecorder()

			e := echo.New()
			ctx := e.NewContext(req, rec)
			middleware.ErrorHandler(log.Noop{})(test.err, ctx)

			response := rec.Result()
			defer response.Body.Close()

			var mm interface{}
			assert.NoError(t, json.NewDecoder(response.Body).Decode(&mm))
			if m, ok := mm.(map[string]interface{}); ok {
				delete(m, "id")
			}

			assert.Equal(t, test.expResp, mm)
			assert.Equal(t, test.expStatusCode, response.StatusCode)
		})
	}
}
