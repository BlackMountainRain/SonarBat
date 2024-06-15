/**
 * Package server
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/6/13
 */

package server

import "github.com/google/wire"

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer)
