/**
 * Package runtime
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/6/30
 */

package runtime

import "time"

type Info struct {
	ID        string
	Name      string
	Version   string
	BuildTime time.Duration
	Uptime    time.Time
}
