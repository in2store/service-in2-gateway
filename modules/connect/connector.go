package connect

import (
	"context"
	"fmt"
	"github.com/in2store/service-in2-gateway/clients/client_in2_auth"
	"github.com/in2store/service-in2-gateway/modules/connect/channels"
	"github.com/in2store/service-in2-gateway/modules/connect/constants"
	"github.com/sirupsen/logrus"
	"sync"
)

var RepoConnector = Connector{}

func init() {
	RepoConnector.RegisterChannelCreator(channels.GithubChannel{}.ChannelID(), channels.NewGithubChannel)
}

type RepoChannelCreator func(ctx context.Context, token client_in2_auth.Token) (constants.RepoChannel, error)

type Connector struct {
	channels sync.Map
}

func (connector *Connector) RegisterChannelCreator(channelID uint64, c RepoChannelCreator) error {
	if _, exist := connector.channels.Load(channelID); exist {
		return fmt.Errorf("repo channel already exist: %d", channelID)
	}
	connector.channels.Store(channelID, c)

	logrus.Infof("%d registered", channelID)
	return nil
}

func (connector *Connector) GetChannelCreator(channelID uint64) (RepoChannelCreator, bool) {
	v, exist := connector.channels.Load(channelID)
	if exist {
		if ch, ok := v.(RepoChannelCreator); ok {
			return ch, exist
		}
	}

	return nil, false
}
