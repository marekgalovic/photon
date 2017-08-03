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

type ZNode struct {
    Name string
    FullPath string
    Data []byte
}

func (n *ZNode) Scan(v interface{}) error {
    return json.Unmarshal(n.Data, &v)
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

func (z *Zookeeper) Watch(path string) (<-chan zk.Event, error) {
    _, _, event, err := z.conn.GetW(z.fullPath(path))

    return event, err
}

func (z *Zookeeper) Children(path string) ([]string, error) {
    children, _, err := z.conn.Children(z.fullPath(path))

    return children, err
}

func (z *Zookeeper) ChildrenW(path string) ([]string, <-chan zk.Event, error) {
    children, _, event, err := z.conn.ChildrenW(z.fullPath(path))

    return children, event, err
}

func (z *Zookeeper) ChildrenData(path string) (map[string]*ZNode, error) {
    children, err := z.Children(path)
    if err != nil {
        return nil, err
    }

    return z.getChildrenData(path, children)
}

func (z *Zookeeper) ChildrenDataW(path string) (map[string]*ZNode, <-chan zk.Event, error) {
    children, event, err := z.ChildrenW(path)
    if err != nil {
        return nil, nil, err
    }

    childrenWithData, err := z.getChildrenData(path, children)
    return childrenWithData, event, err
}

func (z *Zookeeper) Get(path string) (*ZNode, error) {
    data, _, err := z.conn.Get(z.fullPath(path))
    if err != nil {
        return nil, err
    }

    return &ZNode {
        Name: filepath.Base(path),
        FullPath: z.fullPath(path),
        Data: data,
    }, nil
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

func (z *Zookeeper) getChildrenData(path string, children []string) (map[string]*ZNode, error) {
    result := make(map[string]*ZNode)

    for _, child := range children {
        data, err := z.Get(filepath.Join(path, child))
        if err != nil {
            return nil, err
        }
        result[child] = data
    }

    return result, nil
}

func (z *Zookeeper) fullPath(path string) string {
    return filepath.Join("/", z.config.BasePath, path)
}
