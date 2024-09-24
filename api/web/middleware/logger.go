package middleware

import (
	"context"
	"log"

	"connectrpc.com/connect"
)

func LoggingInterceptor() connect.UnaryInterceptorFunc {
	intercepter := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			log.Println("protocol:", req.Peer().Protocol)
			log.Println("method:", req.Spec().Procedure)
			log.Println("query:", req.Any())
			return next(ctx, req)
		})
	}

	return connect.UnaryInterceptorFunc(intercepter)
}
