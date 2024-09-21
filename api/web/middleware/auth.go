package middleware

import (
	"context"
	"log"
	"strings"

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
	token := authorization[0]
	token = strings.Split(token, "Bearer ")[1]
	claims, err := pkg.ParseToken(token)
	if err != nil {
		log.Println("err1:", err)
		return false
	}
	me, err := pkg.GetUserInfo(claims)
	if err != nil {
		log.Println("err2:", err)
		return false
	}
	ctx = context.WithValue(ctx, "id", me.ID)
	ctx = context.WithValue(ctx, "username", me.Username)
	ctx = context.WithValue(ctx, "email", me.Email)
	return true
}
