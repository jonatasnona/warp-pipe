package postgres

import (
	"fmt"
	"strings"
	"time"
)

const (
	// MinConnectTimeout min timeout necessary for connect
	MinConnectTimeout = 1 * time.Second
	// DefaultPort for postgres
	DefaultPort = 5432
	// DefaultUser for postgres
	DefaultUser = "postgres"
)

// ReplicateConfig object
type ReplicateConfig struct {
	Plugin string
	Slot   string

	IgnoreDuplicateObjectError bool
}

// SQLConfig object
type SQLConfig struct {
	Driver string

	CreateDatabaseIfNotExist bool

	ConnectTimeout time.Duration
}

// StreamingReplicateProtocolConfig config object
type StreamingReplicateProtocolConfig struct {
	SendStandByStatusPeriod time.Duration
	WaitMessageTimeout      time.Duration
}

// Config object
type Config struct {
	User     string
	Password string
	Host     string
	Database string

	SQL       SQLConfig
	Replicate ReplicateConfig

	Port uint16

	Streaming StreamingReplicateProtocolConfig
}

// DSN return dsn (data source name)
func (c *Config) DSN(setDatabase, setPassword bool) (dsn string, missing []string) {

	var url []string

	if len(c.User) > 0 {
		url = append(url, fmt.Sprintf("user=%s", c.User))
	} else {
		missing = append(missing, "user")
	}

	if len(c.Password) > 0 && setPassword {
		url = append(url, fmt.Sprintf("password=%s", c.Password))
	} else {
		missing = append(missing, "password")
	}

	if len(c.Host) > 0 {
		url = append(url, fmt.Sprintf("host=%s", c.Host))
	} else {
		missing = append(missing, "host")
	}

	if c.Port > 0 {
		url = append(url, fmt.Sprintf("port=%d", c.Port))
	} else {
		missing = append(missing, "port")
	}

	if len(c.Database) > 0 && setDatabase {
		url = append(url, fmt.Sprintf("database=%s", c.Database))
	} else {
		missing = append(missing, "database")
	}

	dsn = strings.Join(url, " ")
	return dsn, missing
}
