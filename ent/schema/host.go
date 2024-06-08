package schema

import (
	"database/sql/driver"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"fmt"
	"github.com/google/uuid"
	"net"
)

// Host holds the schema definition for the Host entity.
type Host struct {
	ent.Schema
}

// Fields of the Host.
func (Host) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.Bool("status").Default(true),

		field.String("name").NotEmpty().MaxLen(250),
		field.Time("live_at").Optional(),
		field.Bool("is_agent_installed").Default(false),
		field.String("agent_version").MaxLen(10).Optional(),

		field.JSON("ips", []IpWithInfo{}).Optional(),
		field.Int16("net_type").Optional(),

		field.JSON("additions", map[string]interface{}{}).Optional(),
	}
}

// Edges of the Host.
func (Host) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("host_blacklist", HostBlacklist.Type),
	}
}

func (Host) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (Host) Index() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
	}
}

type IpWithInfo struct {
	IP        net.IP
	IsPrivate bool
	IsIPv6    bool
}

// Value implements the driver.Valuer interface.
func (i IpWithInfo) Value() (driver.Value, error) {
	return fmt.Sprintf("(%s,%t,%t)", i.IP.String(), i.IsPrivate, i.IsIPv6), nil
}

// Scan implements the sql.Scanner interface.
func (i IpWithInfo) Scan(value interface{}) error {
	var ipStr string
	var isPrivate, isIPv6 bool
	_, err := fmt.Sscanf(value.(string), "(%s,%t,%t)", &ipStr, &isPrivate, &isIPv6)
	if err != nil {
		return err
	}
	i.IP = net.ParseIP(ipStr)
	i.IsPrivate = isPrivate
	i.IsIPv6 = isIPv6
	return nil
}
