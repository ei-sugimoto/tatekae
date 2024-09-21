package middleware

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// UnaryLoggingInterceptor は、ステータスとメッセージをログに記録するカスタムインターセプターです。
func UnaryLoggingInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		start := time.Now()
		resp, err := handler(ctx, req)
		code := status.Code(err)
		log.Printf("method: %s, status: %s, Errmessage: %v, duration: %s, response: %v",
			info.FullMethod, code, err, time.Since(start), resp)
		return resp, err
	}
}
