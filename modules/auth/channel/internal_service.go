package channel

import (
	"github.com/in2store/service-in2-gateway/clients/client_in2_auth"
	"github.com/in2store/service-in2-gateway/clients/client_in2_user"
)

type InternalService struct {
	clientUser *client_in2_user.ClientIn2User
	clientAuth *client_in2_auth.ClientIn2Auth
}

func NewInternalService(clientUser *client_in2_user.ClientIn2User, clientAuth *client_in2_auth.ClientIn2Auth) *InternalService {
	return &InternalService{
		clientUser,
		clientAuth,
	}
}

func (c *InternalService) GetChannelName() string {
	return "INNER"
}

func (c *InternalService) GetEntityByToken(token string) (*client_in2_user.User, error) {
	tokenRequest := client_in2_auth.GetSessionBySessionIDRequest{
		SessionID: token,
	}
	t, err := c.clientAuth.GetSessionBySessionID(tokenRequest)
	if err != nil {
		return nil, err
	}
	userRequest := client_in2_user.GetUserByUserIDRequest{
		UserID: t.Body.UserID,
	}
	user, err := c.clientUser.GetUserByUserID(userRequest)
	if err != nil {
		return nil, err
	}
	return &user.Body, nil
}
