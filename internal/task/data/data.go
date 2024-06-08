/**
 * Package data
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/5/8
 */

package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	_ "github.com/lib/pq"
	"sonar-bat/ent"
	"sonar-bat/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewTaskRepo)

// Data .
type Data struct {
	cli *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	client, err := ent.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		log.NewHelper(logger).Error("failed opening connection to postgres: ", err)
		return nil, nil, err
	}

	// Run the auto migration tool.
	err = client.Schema.Create(context.Background())
	if err != nil {
		log.NewHelper(logger).Error("failed creating schema resources: ", err)
		return nil, nil, err
	}

	cleanup := func() {
		err = client.Close()
		if err != nil {
			log.NewHelper(logger).Error("failed closing connection: ", err)
		} else {
			log.NewHelper(logger).Info("closing the data resources")
		}
	}

	return &Data{
		cli: client,
	}, cleanup, nil
}
