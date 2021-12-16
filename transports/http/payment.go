package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/libsv/go-p4"
	"github.com/pkg/errors"
)

// paymentHandler is an http handler that supports BIP-270 requests.
type paymentHandler struct {
	svc p4.PaymentService
}

// NewPaymentHandler will create and return a new PaymentHandler.
func NewPaymentHandler(svc p4.PaymentService) *paymentHandler {
	return &paymentHandler{
		svc: svc,
	}
}

// RegisterRoutes will setup all routes with an echo group.
func (h *paymentHandler) RegisterRoutes(g *echo.Group) {
	g.POST(RouteV1Payment, h.createPayment)
}

// @Summary A user will submit an SpvEnvelope along with other information that is validated before being broadcast to the network.
// @Description Creates a payment based on a payment id (the identifier for an invoice).
// @Tags Payment
// @Accept json
// @Produce json
// @Param paymentID path string true "Payment ID"
// @Param body body p4.PaymentCreate true "payment message used in BIP270"
// @Success 201 {object} p4.PaymentACK "if success, error code will be empty, otherwise it will be filled in with reason"
// @Failure 404 {object} p4.ClientError "returned if the paymentID has not been found"
// @Failure 400 {object} p4.ClientError "returned if the user input is invalid, usually an issue with the paymentID"
// @Failure 500 {string} string "returned if there is an unexpected internal error"
// @Router /api/v1/payment/{paymentID} [POST].
func (h *paymentHandler) createPayment(e echo.Context) error {
	args := p4.PaymentCreateArgs{
		PaymentID: e.Param("paymentID"),
	}
	var req p4.Payment
	if err := e.Bind(&req); err != nil {
		return errors.WithStack(err)
	}
	resp, err := h.svc.PaymentCreate(e.Request().Context(), args, req)
	if err != nil {
		return errors.WithStack(err)
	}
	if resp.Error > 0 {
		return e.JSON(http.StatusUnprocessableEntity, resp)
	}
	return e.JSON(http.StatusCreated, resp)
}
