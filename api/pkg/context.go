package pkg

import "context"

var (
	ID       string = "id"
	Username string = "username"
	Email    string = "email"
)

func SetID(ctx context.Context, id int) context.Context {
	return context.WithValue(ctx, ID, id)
}

func SetUsername(ctx context.Context, username string) context.Context {
	return context.WithValue(ctx, Username, username)
}

func SetEmail(ctx context.Context, email string) context.Context {
	return context.WithValue(ctx, Email, email)
}
