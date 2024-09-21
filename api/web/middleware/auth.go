package middleware

import (
	"context"
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
	BearerToken := md["authorization"][0]
	token, ok := ExtractToken(BearerToken)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "missing token")
	}

	me, err := GetMe(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	ctx = setMeContext(ctx, me)

	return handler(ctx, req)
}

func ExtractToken(BearerToken string) (string, bool) {
	splitToken := strings.Split(BearerToken, "Bearer ")
	if len(splitToken) != 2 {
		return "", false
	}
	return splitToken[1], true
}

type MeType struct {
	ID       int
	Username string
	Email    string
}

func GetMe(token string) (*MeType, error) {
	claims, err := pkg.ParseToken(token)
	if err != nil {
		return nil, err
	}
	me, err := pkg.GetUserInfo(claims)

	return &MeType{ID: me.ID, Username: me.Username, Email: me.Email}, err
}

func setMeContext(ctx context.Context, me *MeType) context.Context {
	ctx = pkg.SetID(ctx, me.ID)
	ctx = pkg.SetUsername(ctx, me.Username)
	ctx = pkg.SetEmail(ctx, me.Email)

	return ctx
}
