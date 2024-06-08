/**
 * Package server
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/5/8
 */

package server

import "github.com/google/wire"

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer)
