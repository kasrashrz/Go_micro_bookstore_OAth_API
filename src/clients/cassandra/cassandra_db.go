package cassandra

import (
	"github.com/gocql/gocql"
)

var(
	cluster *gocql.ClusterConfig
)

func init(){
	// connect to the cluster
	cluster = gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oath"
	cluster.Consistency = gocql.Quorum
}

func GetSession() (*gocql.Session, error){
	return cluster.CreateSession()
}