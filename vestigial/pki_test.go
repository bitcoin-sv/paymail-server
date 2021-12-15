package web

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/nch-bowstave/paymail"
	"github.com/nch-bowstave/paymail/mocks"
	"github.com/nch-bowstave/paymail/service"
	"github.com/stretchr/testify/assert"
)

func TestPKI(t *testing.T) {
	e := echo.New()
	tests := map[string]struct {
		mockAccFunc func(ctx context.Context, handle paymail.Handle) (*paymail.PublicAccount, error)
		handle      string
		exp         *paymail.PKI
		code        int
	}{
		"should return 404 for missing handle": {
			mockAccFunc: func(ctx context.Context, handle paymail.Handle) (*paymail.PublicAccount, error) {
				return nil, nil
			},
			code:   http.StatusNotFound,
			handle: "bob@somedomain.com",
		},
		"should return 200 and pki for valid handle": {
			mockAccFunc: func(ctx context.Context, handle paymail.Handle) (*paymail.PublicAccount, error) {
				return &paymail.PublicAccount{
					Alias:     "bob",
					Handle:    "bob@somedomain.com",
					PublicKey: "abc123",
				}, nil
			},
			code:   http.StatusOK,
			handle: "bob@somedomain.com",
			exp: &paymail.PKI{
				BsvAlias:  "1.0",
				Handle:    "bob@somedomain.com",
				PublicKey: "abc123",
			},
		},
	}

	for name, test := range tests {
		svc := service.NewPaymailService(&mocks.AccountReaderWriterMock{
			AccountFunc: test.mockAccFunc,
			CreateFunc:  func(ctx context.Context, account paymail.Account) error { return nil },
		}, "somedomain.com")
		a := NewPKI(svc)
		a.RegisterRoutes(e.Group(""))

		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, routePki, nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)

			ctx.SetParamNames("handle")
			ctx.SetParamValues(test.handle)

			if assert.NoError(t, a.PKI(ctx)) {
				assert.Equal(t, test.code, rec.Code)

				var pki *paymail.PKI
				json.Unmarshal(rec.Body.Bytes(), &pki)

				assert.Equal(t, test.code, rec.Code)
				assert.Equal(t, test.exp, pki)
			}
		})
	}
}
