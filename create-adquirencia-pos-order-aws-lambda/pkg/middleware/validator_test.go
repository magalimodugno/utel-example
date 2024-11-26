package middleware

import (
	"create-adquirencia-pos-order-aws-lambda/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_validation(t *testing.T) {
	type args struct {
		req *service.CreateOrderRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "given a correct input, should not throw any error",
			args: args{
				req: &service.CreateOrderRequest{
					OrderID:   "OrderID",
					AccountID: "AccountID",
					Currency:  "Currency",
					Amount:    100,
					RefNumber: "RefNumber",
					Origin:    "Origin",
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "given an incorrect input (empty account id), should throw an error",
			args: args{
				req: &service.CreateOrderRequest{
					OrderID:   "OrderID",
					AccountID: "",
					Currency:  "Currency",
					Amount:    100,
					RefNumber: "RefNumber",
					Origin:    "Origin",
				},
			},
			wantErr: func(t assert.TestingT, got error, i2 ...interface{}) bool {
				value := "account_id"
				return assert.Containsf(t, got.Error(), value, "error message %s", "formatted")
			},
		},
		{
			name: "given an incorrect input (empty order id), should throw an error",
			args: args{
				req: &service.CreateOrderRequest{
					OrderID:   "",
					AccountID: "AccountID",
					Currency:  "Currency",
					Amount:    100,
					RefNumber: "RefNumber",
					Origin:    "Origin",
				},
			},
			wantErr: func(t assert.TestingT, got error, i2 ...interface{}) bool {
				value := "order_id"
				return assert.Containsf(t, got.Error(), value, "error message %s", "formatted")
			},
		},
		{
			name: "given an incorrect input (empty currency), should throw an error",
			args: args{
				req: &service.CreateOrderRequest{
					OrderID:   "OrderID",
					AccountID: "AccountID",
					Currency:  "",
					Amount:    100,
					RefNumber: "RefNumber",
					Origin:    "Origin",
				},
			},
			wantErr: func(t assert.TestingT, got error, i2 ...interface{}) bool {
				value := "currency"
				return assert.Containsf(t, got.Error(), value, "error message %s", "formatted")
			},
		},
		{
			name: "given an incorrect input (empty reference number), should throw an error",
			args: args{
				req: &service.CreateOrderRequest{
					OrderID:   "OrderID",
					AccountID: "AccountID",
					Currency:  "Currency",
					Amount:    100,
					RefNumber: "",
					Origin:    "Origin",
				},
			},
			wantErr: func(t assert.TestingT, got error, i2 ...interface{}) bool {
				value := "ref_number"
				return assert.Containsf(t, got.Error(), value, "error message %s", "formatted")
			},
		},
		{
			name: "given an incorrect input (wrong amount), should throw an error",
			args: args{
				req: &service.CreateOrderRequest{
					OrderID:   "OrderID",
					AccountID: "AccountID",
					Currency:  "Currency",
					Amount:    -100,
					RefNumber: "RefNumber",
					Origin:    "Origin",
				},
			},
			wantErr: func(t assert.TestingT, got error, i2 ...interface{}) bool {
				value := "invalid"
				return assert.Containsf(t, got.Error(), value, "error message %s", "formatted")
			},
		},
		{
			name: "given an incorrect input (empty amount), should throw an error",
			args: args{
				req: &service.CreateOrderRequest{
					OrderID:   "OrderID",
					AccountID: "AccountID",
					Currency:  "Currency",
					RefNumber: "RefNumber",
					Origin:    "Origin",
				},
			},
			wantErr: func(t assert.TestingT, got error, i2 ...interface{}) bool {
				value := "empty value"
				return assert.Containsf(t, got.Error(), value, "error message %s", "formatted")
			},
		},
	}
	for _, tt := range tests {
		t.Setenv("ENVIRONMENT", "dev")
		t.Run(tt.name, func(t *testing.T) {
			if err := validation(tt.args.req); err != nil {
				tt.wantErr(t, err)
				if err != nil {
					t.Log(err.Error())
				}
			}
		})
	}
}
