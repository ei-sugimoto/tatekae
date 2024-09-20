package middleware

import (
	"context"

	"github.com/ei-sugimoto/tatekae/api/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 特定のサービスにのみ認証を適用
	if info.FullMethod == "/protoUser.UserService/Register" || info.FullMethod == "/protoUser.UserService/Login" {
		return handler(ctx, req)

	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
	}

	// トークンの検証
	if !SetInfo(md["authorization"], ctx) {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}
	// 認証が成功した場合、次のハンドラを呼び出す
	return handler(ctx, req)
}

func SetInfo(authorization []string, ctx context.Context) bool {
	// トークンの検証ロジックを実装
	if len(authorization) < 1 {
		return false
	}
	token := authorization[1]
	claims, err := pkg.ParseToken(token)
	if err != nil {
		return false
	}
	if claims == nil {
		return false
	}
	id := claims["id"].(int)
	if id == 0 {
		return false
	}
	ctx = context.WithValue(ctx, "id", id)
	username := claims["username"].(string)
	if username == "" {
		return false
	}
	ctx = context.WithValue(ctx, "username", username)
	email := claims["email"].(string)
	if email == "" {
		return false
	}
	ctx = context.WithValue(ctx, "email", email)
	return true
}
