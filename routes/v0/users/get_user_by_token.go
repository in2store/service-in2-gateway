package users

import (
	"context"
	"github.com/in2store/service-in2-gateway/routes/middleware"
	"github.com/johnnyeven/libtools/courier"
	"github.com/johnnyeven/libtools/courier/httpx"
)

func init() {
	Router.Register(courier.NewRouter(middleware.MiddlewareAuth, GetUserByToken{}))
}

// 通过token获取用户信息
type GetUserByToken struct {
	httpx.MethodGet
}

func (req GetUserByToken) Path() string {
	return ""
}

func (req GetUserByToken) Output(ctx context.Context) (result interface{}, err error) {
	user := middleware.GetAuthUserFromContext(ctx)
	return user, nil
}
