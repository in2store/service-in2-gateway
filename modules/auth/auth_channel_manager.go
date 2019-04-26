package auth

import (
	"fmt"
	"github.com/in2store/service-in2-gateway/clients/client_in2_user"
	"github.com/in2store/service-in2-gateway/global"
	"github.com/in2store/service-in2-gateway/modules/auth/channel"
	"github.com/sirupsen/logrus"
	"sync"
)

var AuthChannelMgr = &AuthChannelManager{}

func init() {
	AuthChannelMgr.RegisterChannel(channel.NewInternalService(global.Config.ClientUser, global.Config.ClientAuth))
}

type AuthTokenChannel interface {
	GetChannelName() string
	GetEntityByToken(token string) (*client_in2_user.User, error)
	GetEntriesByEntity(entityID uint64) (client_in2_user.UserEntryList, error)
}

type AuthChannelManager struct {
	channels sync.Map
}

func (am *AuthChannelManager) RegisterChannel(c AuthTokenChannel) error {
	if _, exist := am.channels.Load(c.GetChannelName()); exist {
		return fmt.Errorf("auth token channel already exist: %s", c.GetChannelName())
	}
	am.channels.Store(c.GetChannelName(), c)

	logrus.Infof("%s registered", c.GetChannelName())
	return nil
}

func (am *AuthChannelManager) GetChannel(name string) (AuthTokenChannel, bool) {
	v, exist := am.channels.Load(name)
	if exist {
		if ch, ok := v.(AuthTokenChannel); ok {
			return ch, exist
		}
	}

	return nil, false
}
