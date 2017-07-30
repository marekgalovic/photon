package storage

import (
    "github.com/gocql/gocql"
)

type CassandraConfig struct {
    Nodes []string
    Keyspace string
    Username string
    Password string
}

type Cassandra struct {
    *gocql.Session
}

func NewCassandra(config CassandraConfig) (*Cassandra, error) {
    cluster := gocql.NewCluster(config.Nodes...)
    cluster.Keyspace = config.Keyspace
    cluster.Consistency = gocql.Quorum
    cluster.Authenticator = gocql.PasswordAuthenticator{Username: config.Username, Password: config.Password}

    session, err := cluster.CreateSession()
    if err != nil {
        return nil, err
    }
    
    return &Cassandra{session}, nil
}
