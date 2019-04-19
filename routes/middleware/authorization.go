package middleware

import (
	"context"
	"github.com/in2store/service-in2-gateway/clients/client_in2_user"
	"github.com/in2store/service-in2-gateway/constants/errors"
	"github.com/in2store/service-in2-gateway/modules/auth"
	"github.com/johnnyeven/libtools/courier"
	"github.com/sirupsen/logrus"
	"strings"
)

const contextKeyAuthorization = "MIDDLEWARE_AUTH"

var MiddlewareAuth = Authorization{}

// 认证中间件
type Authorization struct {
	Token string `name:"X-Token" in:"header"`
}

func (req Authorization) ContextKey() string {
	return contextKeyAuthorization
}

func (req Authorization) Output(ctx context.Context) (result interface{}, err error) {
	tokens := strings.Split(req.Token, ":")
	if len(tokens) < 2 {
		return nil, errors.BadAuthToken
	}
	c, exist := auth.AuthChannelMgr.GetChannel(strings.ToUpper(tokens[0]))
	if !exist {
		return nil, errors.BadAuthChannal
	}
	resp, err := c.GetEntityByToken(strings.Join(tokens[1:], ":"))
	if err != nil {
		logrus.Errorf("Authorization.GetEntityByToken err: %v, request: %v", err, req.Token)
		return nil, err
	}
	return resp, nil
}

func GetAuthUserFromContext(ctx context.Context) *client_in2_user.User {
	value := courier.GetContextValue(ctx, &Authorization{})
	s, ok := value.(*client_in2_user.User)
	if !ok {
		logrus.Panicf("GetAuthUserFromContext format err: %+v", s)
	}
	return s
}
