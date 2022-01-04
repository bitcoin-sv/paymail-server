package service

import (
	"context"
	"encoding/json"

	"github.com/libsv/go-bk/envelope"
	"github.com/pkg/errors"
	validator "github.com/theflyingcodr/govalidator"

	"github.com/libsv/go-p4"
)

// proof enforces business rules.
type proof struct {
	store p4.ProofsWriter
}

// NewProof will setup a new proof service.
func NewProof(store p4.ProofsWriter) *proof {
	return &proof{
		store: store,
	}
}

// Create will add an object to the data store, rejecting the request
// if it fails to match required validation params.
func (s *proof) Create(ctx context.Context, args p4.ProofCreateArgs, req envelope.JSONEnvelope) error {
	var proof *p4.ProofWrapper
	if err := json.Unmarshal([]byte(req.Payload), &proof); err != nil {
		return errors.Wrap(err, "failed to unmarshall JSONEnvelope")
	}
	if err := validator.New().Validate("jsonEnvelope", func() error {
		if ok, err := req.IsValid(); !ok || err != nil {
			return errors.Wrap(err, "invalid merkleProof envelope")
		}
		return nil
	}).Err(); err != nil {
		return err
	}
	if err := proof.Validate(args); err != nil {
		return err
	}
	if err := s.store.ProofCreate(ctx, args, req); err != nil {
		return errors.Wrapf(err, "failed to add proof with txid '%s' and invoiceID '%s'", args.TxID, args.PaymentReference)
	}
	return nil
}
