package storage

import (
    "time";
    "strings";
    "path/filepath";
    "encoding/json";

    "github.com/samuel/go-zookeeper/zk";
)

type ZookeeperConfig struct {
    Nodes []string
    SessionTimeout time.Duration
    BasePath string
}

type Zookeeper struct {
    conn *zk.Conn
    config ZookeeperConfig
}

func NewZookeeper(config ZookeeperConfig) (*Zookeeper, error) {
    conn, _, err := zk.Connect(config.Nodes, config.SessionTimeout)
    if err != nil {
        return nil, err
    }

    return &Zookeeper{
        conn: conn,
        config: config,
    }, nil
}

func (z *Zookeeper) Close() {
    z.conn.Close()
}

func (z *Zookeeper) Exists(path string) (bool, error) {
    exists, _, err := z.conn.Exists(z.fullPath(path))

    return exists, err
}

func (z *Zookeeper) Children(path string) ([]string, error) {
    children, _, err := z.conn.Children(z.fullPath(path))

    return children, err
}

func (z *Zookeeper) Get(path string) (interface{}, error) {
    data, _, err := z.conn.Get(z.fullPath(path))
    if err != nil {
        return nil, err
    }

    var unmarshaledData interface{}
    err = json.Unmarshal(data, &unmarshaledData)

    return unmarshaledData, err
}

func (z *Zookeeper) Set(path string, data interface{}, version int32) error {
    marshaledData, err := json.Marshal(data)
    if err != nil {
        return err
    }

    _, err = z.conn.Set(z.fullPath(path), marshaledData, version)
    return err
}

func (z *Zookeeper) Create(path string, data interface{}, flags int32, acl []zk.ACL) (string, error) {
    marshaledData, err := json.Marshal(data)
    if err != nil {
        return "", err
    }

    pathParts := strings.Split(z.fullPath(path), "/")
    for i, _ := range pathParts {
        zNodePath := strings.Join(pathParts[:i], "/")
        if len(zNodePath) == 0 {
            continue
        }

        exists, _, err := z.conn.Exists(zNodePath)
        if err != nil {
            return "", err
        }
        if exists {
            continue
        }

        if _, err = z.conn.Create(zNodePath, nil, flags, zk.WorldACL(zk.PermAll)); err != nil {
            return "", err
        }
    }

    return z.conn.Create(z.fullPath(path), marshaledData, flags, acl)
}

func (z *Zookeeper) Delete(path string, version int32) error {
    return z.conn.Delete(z.fullPath(path), version)
}

func (z *Zookeeper) fullPath(path string) string {
    return filepath.Join("/", z.config.BasePath, path)
}
