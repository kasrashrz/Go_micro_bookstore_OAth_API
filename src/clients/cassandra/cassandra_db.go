package cassandra

import (
	"github.com/gocql/gocql"
	"time"
)

var (
	session *gocql.Session
)

func init() {
	// connect to the cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oath"
	cluster.Consistency = gocql.Quorum
	cluster.Timeout = 50 * time.Second

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}
