/**
 * Package server
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/6/27
 */

package server

import (
	v1 "sonar-bat/api/resource/v1"
	"sonar-bat/internal/conf"
	"sonar-bat/internal/resource/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, resource *service.ResourceService, _ log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterResourceHTTPServer(srv, resource)
	return srv
}
