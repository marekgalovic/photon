package server

import (
    "fmt";

    "github.com/marekgalovic/photon/server/storage"
)

type Config struct {
    Address string
    Port int
    Env string
    Root string
    Mysql storage.MysqlConfig
    Cassandra storage.CassandraConfig
}

func NewConfig() *Config {
    return &Config{
        Port: 5005,
        Root: "./",
        Env: "development",
        Mysql: storage.MysqlConfig{
            User: "root",
            Host: "127.0.0.1",
            Port: 3306,
            Database: "serving_development",
        },
        Cassandra: storage.CassandraConfig{
            Nodes: []string{"127.0.0.1"},
            Keyspace: "development",
            Username: "cassandra",
            Password: "cassandra",
        },
    }
}

func (c *Config) BindAddress() string {
    return fmt.Sprintf("%s:%d", c.Address, c.Port)
}
