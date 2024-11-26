package middleware

import (
	"context"
	"create-adquirencia-pos-order-aws-lambda/service"
	"errors"
	"github.com/Bancar/uala-bis-go-dependencies/validator"
)

var (
	errEmpty   = errors.New("is an empty value")
	errInvalid = errors.New("is invalid")
)

func RequestValidation(next func(context.Context, *service.CreateOrderRequest) error) func(context.Context, *service.CreateOrderRequest) error {
	return func(ctx context.Context, req *service.CreateOrderRequest) error {
		if err := validation(req); err != nil {
			return err
		}
		return next(ctx, req)
	}
}

func validation(req *service.CreateOrderRequest) error {
	validations := map[string]func(s string) error{
		"order_id":   validator.ValidateEmpty(req.OrderID),
		"account_id": validator.ValidateEmpty(req.AccountID),
		"currency":   validator.ValidateEmpty(req.Currency),
		"amount":     ValidatePositiveFloat64(req.Amount),
		"ref_number": validator.ValidateEmpty(req.RefNumber),
	}
	return validator.ValidateValues(validations)
}

func ValidatePositiveFloat64(num float64) func(string) error {
	return func(_ string) error {
		if num == 0 {
			return errEmpty
		}
		if num < 0 {
			return errInvalid
		}
		return nil
	}
}
