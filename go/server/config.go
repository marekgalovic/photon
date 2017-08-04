package server

import (
    "os";
    "fmt";
    "flag";
    "time";
    "strings";
    "path/filepath";

    "github.com/marekgalovic/photon/go/core/storage";

    "github.com/BurntSushi/toml";
    log "github.com/Sirupsen/logrus"
)

type Config struct {
    Env string
    Root string
    ConfigPath string
    Address string
    Port int
    Mysql storage.MysqlConfig
    Cassandra storage.CassandraConfig
    Zookeeper storage.ZookeeperConfig
}

func NewConfig() (*Config, error) {
    config := &Config{
        Port: 5005,
        Root: "./",
        Env: "development",
        Mysql: storage.MysqlConfig{
            User: "root",
            Host: "127.0.0.1",
            Port: 3306,
            Database: "photon",
        },
        Cassandra: storage.CassandraConfig{
            Nodes: []string{"127.0.0.1"},
            Keyspace: "photon",
            Consistency: "quorum",
        },
        Zookeeper: storage.ZookeeperConfig{
            Nodes: []string{"127.0.0.1:2181"},
            SessionTimeout: 1 * time.Second,
            BasePath: "/photon",
        },
    }

    if err := config.parse(); err != nil {
        return nil, err
    }

    return config, nil
}

func (c *Config) parse() error {
    c.Env = c.getEnvDefault("PHOTON_ENV", c.Env)
    c.Root = c.getEnvDefault("PHOTON_ROOT", c.Root)
    c.ConfigPath = c.getEnvDefault("PHOTON_CONF", filepath.Join(c.Root, "config", fmt.Sprintf("%s.tml", c.Env)))

    _, err := toml.DecodeFile(c.ConfigPath, &c)
    if err != nil {
        return err
    }

    c.parseFlags()
    return nil
}

func (c *Config) parseFlags() {
    flag.StringVar(&c.Env, "env", c.Env, "Server environment.")
    flag.StringVar(&c.Root, "root", c.Root, "Server root path.")
    flag.StringVar(&c.ConfigPath, "conf", c.ConfigPath, "Server config file.")
    // Listener
    flag.StringVar(&c.Address, "address", c.Address, "Server host.")
    flag.IntVar(&c.Port, "port", c.Port, "Server port.")
    // Mysql
    flag.StringVar(&c.Mysql.User, "mysql-user", c.Mysql.User, "Mysql username.")
    if c.Mysql.Password == "" {
        flag.StringVar(&c.Mysql.Password, "mysql-password", "", "Mysql password.")
    }
    flag.StringVar(&c.Mysql.Host, "mysql-host", c.Mysql.Host, "Mysql host.")
    flag.IntVar(&c.Mysql.Port, "mysql-port", c.Mysql.Port, "Mysql port.")
    flag.StringVar(&c.Mysql.Database, "mysql-database", c.Mysql.Database, "Mysql database.")
    // Cassandra
    c.Cassandra.Nodes = strings.Split(*flag.String("cassandra-nodes", strings.Join(c.Cassandra.Nodes, ","), "Cassandra nodes (comma delimited)."), ",")
    flag.StringVar(&c.Cassandra.Keyspace, "cassandra-keyspace", c.Cassandra.Keyspace, "Cassandra keyspace.")
    flag.IntVar(&c.Cassandra.ProtoVersion, "cassandra-proto-version", c.Cassandra.ProtoVersion, "Cassandra protocol version.")
    flag.StringVar(&c.Cassandra.Consistency, "cassandra-consistency", c.Cassandra.Consistency, "Cassandra consistency (one, quorum).")
    flag.StringVar(&c.Cassandra.Username, "cassandra-user", c.Cassandra.Username, "Cassandra user.")
    if c.Cassandra.Password == "" {
        flag.StringVar(&c.Cassandra.Password, "cassandra-password", "", "Cassandra password.")
    }
    // Zookeeper
    c.Zookeeper.Nodes = strings.Split(*flag.String("zookeeper-nodes", strings.Join(c.Zookeeper.Nodes, ","), "Zookeeper nodes (comma delimited)."), ",")
    flag.StringVar(&c.Zookeeper.BasePath, "zookeeper-basepath", c.Zookeeper.BasePath, "Zookeeper base path.")

    flag.Parse()
}

func (c *Config) getEnvDefault(key string, defaultValue string) string {
    if envValue, exists := os.LookupEnv(key); exists {
        return envValue
    }
    return defaultValue
}

func (c *Config) Print() {
    log.Info("Photon server")
    log.Infof("Environment: %s", c.Env)
    log.Infof("Server root: %s", c.Root)
    log.Infof("Loaded config: %s", c.ConfigPath)
}

func (c *Config) BindAddress() string {
    return fmt.Sprintf("%s:%d", c.Address, c.Port)
}
