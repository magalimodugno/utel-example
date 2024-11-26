package main

import (
	"context"
	"create-adquirencia-pos-order-aws-lambda/pkg/middleware"
	"create-adquirencia-pos-order-aws-lambda/service"
	"errors"
	"github.com/Bancar/goala/ulog"
	"github.com/Bancar/goala/utel"
	"github.com/Bancar/lambda-go"
	"go.opentelemetry.io/otel/codes"
	"os"
)

const (
	_owner = "adquirencia"
	_flow  = "orders-v4"
)

var errUtel = errors.New("error when shutting down instrumentation")

func main() {
	os.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "otel-collector.otel-collector.local:4317")
	s := service.New(
		&service.Clients{})

	lambda.AsyncStart(s.CreateOrder, false, TelemetryMiddleware, middleware.RequestValidation)
}

func TelemetryMiddleware(next func(ctx context.Context, n *service.CreateOrderRequest) error) func(ctx context.Context, n *service.CreateOrderRequest) error {
	return func(ctx context.Context, n *service.CreateOrderRequest) error {

		//Create and set utel config
		utel.SetUtelConfig(_owner, _flow)

		//Enable instrumentation for basic configuration
		instrumentation, err := utel.EnableInstrumentation(ctx)
		if err != nil {
			return errUtel
		}
		defer func() {
			if err := instrumentation(ctx); err != nil {
				ulog.Error(err.Error())
			}
		}()

		//Create span to log traces
		name := os.Getenv("AWS_LAMBDA_FUNCTION_NAME")
		ctx, span := utel.NewSpan(ctx, name)
		defer span.End()

		if err := next(ctx, n); err != nil {
			span.SetStatus(codes.Error, "error")
			_ = utel.IncrementCounter(ctx, "error_count")
			return err
		}
		span.SetStatus(codes.Ok, "success")
		_ = utel.IncrementCounter(ctx, "success_count")
		return nil
	}
}
