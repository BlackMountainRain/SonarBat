/**
 * Package server
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/6/13
 */

package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	v1 "sonar-bat/api/auth/v1"
	"sonar-bat/internal/auth/service"
	"sonar-bat/internal/conf"
	"strings"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, auth *service.AuthService, _ log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
			handlers.ExposedHeaders([]string{"Content-Length"}),
			handlers.AllowCredentials(),
		)),
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
	if c.Http.EnableCors {
		corsOpts := []handlers.CORSOption{
			handlers.AllowedOrigins(strings.Split(c.Http.AllowedOrigins, ",")),
			handlers.AllowedMethods(c.Http.AllowedMethods),
			handlers.AllowedHeaders(c.Http.AllowedHeaders),
			handlers.ExposedHeaders(c.Http.ExposedHeaders),
		}
		if c.Http.AllowCredentials {
			corsOpts = append(corsOpts, handlers.AllowCredentials())
		}
		opts = append(opts, http.Filter(handlers.CORS(corsOpts...)))
	}
	srv := http.NewServer(opts...)
	v1.RegisterAuthHTTPServer(srv, auth)
	return srv
}
