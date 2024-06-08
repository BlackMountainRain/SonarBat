/**
 * Package biz
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/5/8
 */

package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewTaskUseCase)
