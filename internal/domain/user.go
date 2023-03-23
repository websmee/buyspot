package domain

import "context"

const CtxKeyUser = "user"

type User struct {
	ID string
}

func GetCtxUser(ctx context.Context) *User {
	if ctx.Value(CtxKeyUser) != nil {
		return ctx.Value(CtxKeyUser).(*User)
	}

	return nil
}
