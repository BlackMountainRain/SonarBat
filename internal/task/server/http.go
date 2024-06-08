/**
 * Package server
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/5/6
 */

package server

import (
	v1 "sonar-bat/api/task/v1"
	"sonar-bat/internal/conf"
	"sonar-bat/internal/task/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, task *service.TaskService, subtask *service.SubtaskService, logger log.Logger) *http.Server {
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
	v1.RegisterTaskHTTPServer(srv, task)
	v1.RegisterSubtaskHTTPServer(srv, subtask)
	return srv
}
