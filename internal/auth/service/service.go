/**
 * Package service
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/6/13
 */

package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAuthService)
