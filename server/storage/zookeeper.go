package storage

import (
    "time";

    "github.com/samuel/go-zookeeper/zk"
)

type ZookeeperConfig struct {
    Nodes []string
    SessionTimeout time.Duration
}

type Zookeeper struct {
    *zk.Conn
}

func NewZookeeper(config ZookeeperConfig) (*Zookeeper, error) {
    conn, _, err := zk.Connect(config.Nodes, config.SessionTimeout)
    if err != nil {
        return nil, err
    }

    return &Zookeeper{conn}, nil
}
