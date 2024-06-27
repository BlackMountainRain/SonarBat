/**
 * Package biz
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/6/27
 */

package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewHostManager)
