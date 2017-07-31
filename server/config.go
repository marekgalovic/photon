package server

import (
    "os";
    "fmt";
    "flag";
    "strings";
    "path/filepath";

    "github.com/marekgalovic/photon/server/storage";

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
            Database: "serving_development",
        },
        Cassandra: storage.CassandraConfig{
            Nodes: []string{"127.0.0.1"},
            Keyspace: "development",
            Username: "cassandra",
            Password: "cassandra",
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

    c.parseFlags()

    _, err := toml.DecodeFile(c.ConfigPath, &c)
    return err
}

func (c *Config) parseFlags() {
    flag.StringVar(&c.Env, "env", c.Env, "Server environment.")
    flag.StringVar(&c.Root, "root", c.Root, "Server root path.")
    flag.StringVar(&c.ConfigPath, "conf", c.ConfigPath, "Server config file.")
    flag.StringVar(&c.Address, "address", "0.0.0.0", "Server host.")
    flag.IntVar(&c.Port, "port", 5021, "Server port.")
    flag.StringVar(&c.Mysql.User, "mysql-user", "root", "Mysql username.")
    flag.StringVar(&c.Mysql.Password, "mysql-password", "", "Mysql password.")
    flag.StringVar(&c.Mysql.Host, "mysql-host", "127.0.0.1", "Mysql host.")
    flag.IntVar(&c.Mysql.Port, "mysql-port", 3306, "Mysql port.")
    flag.StringVar(&c.Mysql.Database, "mysql-database", "photon", "Mysql database.")
    c.Cassandra.Nodes = strings.Split(*flag.String("cassandra-nodes", "127.0.0.1", "Cassandra nodes (comma delimited)."), ",")
    flag.StringVar(&c.Cassandra.Keyspace, "cassandra-keyspace", "photon", "Cassandra keyspace.")
    flag.StringVar(&c.Cassandra.Username, "cassandra-user", "cassandra", "Cassandra user.")
    flag.StringVar(&c.Cassandra.Password, "cassandra-password", "cassandra", "Cassandra password.")

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
