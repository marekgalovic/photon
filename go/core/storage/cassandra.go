package storage

import (
    "fmt";
    "time";
    
    "github.com/gocql/gocql"
)

type CassandraConfig struct {
    Nodes []string
    Keyspace string
    ProtoVersion int
    Username string
    Password string
}

type Cassandra struct {
    *gocql.Session
}

func NewCassandra(config CassandraConfig) (*Cassandra, error) {
    cluster := gocql.NewCluster(config.Nodes...)
    cluster.Timeout = 5 * time.Second
    cluster.Keyspace = config.Keyspace
    cluster.ProtoVersion = config.ProtoVersion
    cluster.Consistency = gocql.Quorum
    if config.Username != "" {
        cluster.Authenticator = gocql.PasswordAuthenticator{Username: config.Username, Password: config.Password}
    }

    session, err := cluster.CreateSession()
    if err != nil {
        return nil, fmt.Errorf("Failed to connect to Cassandra. %v", err)
    }
    
    return &Cassandra{session}, nil
}

func (cassandra *Cassandra) Count(tableName string) (int, error) {
    var count int
    if err := cassandra.Query(fmt.Sprintf(`SELECT count(1) FROM %s`, tableName)).Scan(&count); err != nil {
        return 0, err
    }
    return count, nil
}
