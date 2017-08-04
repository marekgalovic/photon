package storage

import (
    "fmt";
    "time";

    log "github.com/Sirupsen/logrus"
)

func NewTestMysql() *Mysql {
    mysql, err := NewMysql(MysqlConfig{Host: "127.0.0.1", Port: 3306, User: "root", Database: "photon_test"})
    if err != nil {
        log.Fatal(err)
    }
    return mysql
}

func NewTestCassandra() *Cassandra {
    cassandra, err := NewCassandra(CassandraConfig{Nodes: []string{"127.0.0.1"}, Keyspace: "photon_test", Username: "cassandra", Password: "cassandra"})
    if err != nil {
        log.Fatal(err)
    }
    return cassandra
}

func NewTestZookeeper() *Zookeeper {
    zookeeper, err := NewZookeeper(ZookeeperConfig{Nodes: []string{"127.0.0.1:2181"}, SessionTimeout: 100 * time.Millisecond, BasePath: "photon_test"})
    if err != nil {
        log.Fatal(err)
    }
    return zookeeper
}

func CleanupMysql(db *Mysql, tables ...string) {
    for _, table := range tables {
        if _, err := db.Exec(fmt.Sprintf("DELETE FROM %s", table)); err != nil {
            log.Fatal(err)
        }
    }
}

func CleanupCassandra(db *Cassandra, tables ...string) {
    for _, table := range tables {
        if err := db.Query(fmt.Sprintf("DELETE FROM %s", table)).Exec(); err != nil {
            log.Fatal(err)
        }
    }
}

func AssertCountChanged(db Countable, tableName string, diff int, callable func()) {
    expected, err := db.Count(tableName)
    if err != nil {
        log.Fatal(err)
    }
    expected += diff

    callable()
    
    actual, _ := db.Count(tableName)

    if expected != actual {
        log.Fatalf("Expected count %d != %d. Table: %s", expected, actual, tableName)
    }
}
