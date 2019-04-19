package global

import (
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
}{
	Log: &log.Log{
		Level: "DEBUG",
	},
	Server: transport_http.ServeHTTP{
		WithCORS: true,
		Port:     8000,
	},
}
