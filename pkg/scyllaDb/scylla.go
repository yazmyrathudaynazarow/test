package scyllaDb

import (
	"fmt"
	"github.com/gocql/gocql"
	"test/project/pkg/config"
)

func NewScyllaDB(cfg *config.Configs) (*gocql.Session, error) {
	cluster := gocql.NewCluster(cfg.ScyllaDB.Host)
	cluster.Keyspace = cfg.ScyllaDB.KeySpace
	cluster.Consistency = gocql.Quorum
	cluster.Port = cfg.ScyllaDB.Port

	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: cfg.ScyllaDB.Username,
		Password: cfg.ScyllaDB.Password,
	}

	if cfg.ScyllaDB.SSL {
		cluster.SslOpts = &gocql.SslOptions{
			EnableHostVerification: false,
		}
	}

	session, err := cluster.CreateSession()
	if err != nil {
		fmt.Println("Error connecting to scyllaDB--->", err)
	}

	return session, err
}
