package web

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/nch-bowstave/paymail"
	"github.com/nch-bowstave/paymail/mocks"
	"github.com/nch-bowstave/paymail/service"
	"github.com/stretchr/testify/assert"
)

func TestRegisterAccount(t *testing.T) {
	e := echo.New()
	svc := service.NewPaymailService(&mocks.AccountReaderWriterMock{
		AccountFunc: func(context.Context, paymail.Handle) (*paymail.PublicAccount, error) {
			return &paymail.PublicAccount{}, nil
		},
		CreateFunc: func(context.Context, paymail.Account) error { return nil },
	}, "somedomain.com")
	a := NewAccount(svc)
	a.RegisterRoutes(e.Group(""))

	tests := map[string]struct {
		json string
		err  error
	}{
		"should return 400 when no payload": {
			err: errors.New("code=400, message=Request body can't be empty"),
		},
		"should return error when invalid email address": {
			json: `{"alias":"bob","name":"Bob Bobson","avatar_url":"","email":"invalid","mobile":"999999"}`,
			err:  errors.New("invalid email address"),
		},
		"should not return error when valid request": {
			json: `{"alias":"bob","name":"Bob Bobson","avatar_url":"","email":"bob@gmail.com","mobile":"999999"}`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, routeRegister, strings.NewReader(test.json))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			ctx := e.NewContext(req, rec)
			err := a.Account(ctx)
			if err != nil {
				assert.EqualError(t, err, test.err.Error())
			}
		})
	}

}
