package global

import (
	"github.com/in2store/service-in2-gateway/clients/client_in2_auth"
	"github.com/in2store/service-in2-gateway/clients/client_in2_user"
	"github.com/johnnyeven/libtools/courier/client"
	"github.com/johnnyeven/libtools/courier/transport_http"
	"github.com/johnnyeven/libtools/log"
	"github.com/johnnyeven/libtools/servicex"
)

func init() {
	servicex.SetServiceName("service-in2-gateway")
	servicex.ConfP(&Config)
}

var Config = struct {
	Log    *log.Log
	Server transport_http.ServeHTTP

	GithubChannelID uint64 `conf:"env"`

	ClientAuth *client_in2_auth.ClientIn2Auth
	ClientUser *client_in2_user.ClientIn2User
}{
	Log: &log.Log{
		Level: "DEBUG",
	},
	Server: transport_http.ServeHTTP{
		WithCORS: true,
		Port:     8000,
	},

	GithubChannelID: 565955135633625088,

	ClientUser: &client_in2_user.ClientIn2User{
		Client: client.Client{
			Host: "service-in2-user.in2store.service.profzone.net",
		},
	},
	ClientAuth: &client_in2_auth.ClientIn2Auth{
		Client: client.Client{
			Host: "service-in2-auth.in2store.service.profzone.net",
		},
	},
}
