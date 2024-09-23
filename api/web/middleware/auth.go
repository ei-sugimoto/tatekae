package middleware

import (
	"context"
	"strings"

	"connectrpc.com/connect"
	"github.com/ei-sugimoto/tatekae/api/pkg"
)

// func NewAuthInterceptor(issuer string, keyPath string) connect.UnaryInterceptorFunc {
// 	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
// 		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
// 			// リクエストヘッダーからJWTを取得する
// 			token := req.Header().Get("Authorization")
// 			if token == "" {
// 				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("error: invalid token"))
// 			}
// 			token = strings.TrimPrefix(token, "Bearer")
// 			token = strings.TrimSpace(token)

// 			// トークンを検証しUserIDを取得する
// 			tm, err := auth.NewTokenManager(issuer, keyPath)
// 			if err != nil {
// 				return nil, connect.NewError(connect.CodeAborted, err)
// 			}
// 			uid, err := tm.GetUserID(token)
// 			if err != nil {
// 				return nil, connect.NewError(connect.CodeUnauthenticated, err)
// 			}

// 			// コンテキストにUserIDをセットする
// 			cw := contextkey.NewContextWriter()
// 			ctx = cw.SetUserID(ctx, uid)

// 			return next(ctx, req)
// 		})
// 	}
// 	return connect.UnaryInterceptorFunc(interceptor)
// }

// func AuthUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
// 	// 特定のサービスにのみ認証を適用
// 	if info.FullMethod == "/protoUser.UserService/Register" || info.FullMethod == "/protoUser.UserService/Login" {
// 		return handler(ctx, req)

// 	}
// 	md, ok := metadata.FromIncomingContext(ctx)
// 	if !ok {
// 		return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
// 	}
// 	BearerToken := md["authorization"][0]
// 	token, ok := ExtractToken(BearerToken)
// 	if !ok {
// 		return nil, status.Errorf(codes.Unauthenticated, "missing token")
// 	}

// 	me, err := GetMe(token)
// 	if err != nil {
// 		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
// 	}

// 	ctx = setMeContext(ctx, me)

// 	return handler(ctx, req)
// }

func AuthUnaryInterceptor() connect.UnaryInterceptorFunc {
	intercepter := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			token := req.Header().Get("Authorization")

			if token == "" {
				return nil, connect.NewError(connect.CodeUnauthenticated, nil)
			}

			extractToken, ok := ExtractToken(token)
			if !ok {
				return nil, connect.NewError(connect.CodeUnauthenticated, nil)
			}

			me, err := GetMe(extractToken)
			if err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, nil)
			}

			ctx = setMeContext(ctx, me)

			return next(ctx, req)
		})
	}

	return connect.UnaryInterceptorFunc(intercepter)
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
