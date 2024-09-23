package middleware

import (
	"context"
	"log"

	"connectrpc.com/connect"
)

// UnaryLoggingInterceptor は、ステータスとメッセージをログに記録するカスタムインターセプターです。
// func UnaryLoggingInterceptor() grpc.UnaryServerInterceptor {
// 	return func(
// 		ctx context.Context,
// 		req interface{},
// 		info *grpc.UnaryServerInfo,
// 		handler grpc.UnaryHandler,
// 	) (interface{}, error) {
// 		start := time.Now()
// 		resp, err := handler(ctx, req)
// 		code := status.Code(err)
// 		log.Printf("method: %s, status: %s, Errmessage: %v, duration: %s, response: %v",
// 			info.FullMethod, code, err, time.Since(start), resp)
// 		return resp, err
// 	}
// }

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
