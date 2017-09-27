package server

import (
    "fmt";
    "flag";
    "time";
    "strings";
    "path/filepath";

    "github.com/marekgalovic/photon/go/core/utils";
    "github.com/marekgalovic/photon/go/core/storage";
    "github.com/marekgalovic/photon/go/core/cluster";

    "github.com/BurntSushi/toml";
    log "github.com/Sirupsen/logrus"
)

type Config struct {
    Env string
    Root string
    ConfigPath string
    Server ServerConfig
    Mysql storage.MysqlConfig
    Cassandra storage.CassandraConfig
    Zookeeper storage.ZookeeperConfig
    Kubernetes cluster.KubernetesConfig
    DeploymentTypes map[string]*cluster.DeploymentType
}

func NewConfig() (*Config, error) {
    config := &Config{
        Env: utils.GetEnv("PHOTON_ENV", "development"),
        Root: utils.GetEnv("PHOTON_ROOT", "./"),
        Server: ServerConfig {
            Port: 5005,
        },
        Mysql: storage.MysqlConfig{
            User: "root",
            Host: "127.0.0.1",
            Port: 3306,
            Database: "photon",
        },
        Cassandra: storage.CassandraConfig{
            Nodes: []string{"127.0.0.1"},
            Keyspace: "photon",
        },
        Zookeeper: storage.ZookeeperConfig{
            Nodes: []string{"127.0.0.1:2181"},
            SessionTimeout: 1 * time.Second,
            BasePath: "/photon",
        },
        Kubernetes: cluster.KubernetesConfig{
            Namespace: "photon",
        },
        DeploymentTypes: map[string]*cluster.DeploymentType{
            "pmml": {RunnerImage: "marekgalovic/photon_pmml_runner", DeployerImage: "marekgalovic/photon_deployer"},
        },
    }

    if err := config.parse(); err != nil {
        return nil, err
    }

    return config, nil
}

func (c *Config) parse() error {
    c.ConfigPath = utils.GetEnv("PHOTON_CONF", filepath.Join(c.Root, "config", fmt.Sprintf("%s.tml", c.Env)))

    c.parseFlags()
    _, err := toml.DecodeFile(c.ConfigPath, &c)
    return err
}

func (c *Config) parseFlags() {
    flag.StringVar(&c.Env, "env", c.Env, "Server environment.")
    flag.StringVar(&c.Root, "root", c.Root, "Server root path.")
    flag.StringVar(&c.ConfigPath, "conf", c.ConfigPath, "Server config file.")
    // Listener
    flag.StringVar(&c.Server.Address, "address", c.Server.Address, "Server host.")
    flag.IntVar(&c.Server.Port, "port", c.Server.Port, "Server port.")
    // Mysql
    flag.StringVar(&c.Mysql.User, "mysql.user", c.Mysql.User, "Mysql username.")
    if c.Mysql.Password == "" {
        flag.StringVar(&c.Mysql.Password, "mysql.password", "", "Mysql password.")
    }
    flag.StringVar(&c.Mysql.Host, "mysql.host", c.Mysql.Host, "Mysql host.")
    flag.IntVar(&c.Mysql.Port, "mysql.port", c.Mysql.Port, "Mysql port.")
    flag.StringVar(&c.Mysql.Database, "mysql.database", c.Mysql.Database, "Mysql database.")
    // Cassandra
    c.Cassandra.Nodes = strings.Split(*flag.String("cassandra.nodes", strings.Join(c.Cassandra.Nodes, ","), "Cassandra nodes (comma delimited)."), ",")
    flag.StringVar(&c.Cassandra.Keyspace, "cassandra.keyspace", c.Cassandra.Keyspace, "Cassandra keyspace.")
    flag.IntVar(&c.Cassandra.ProtoVersion, "cassandra.protocol-version", c.Cassandra.ProtoVersion, "Cassandra protocol version.")
    flag.StringVar(&c.Cassandra.Username, "cassandra.user", c.Cassandra.Username, "Cassandra user.")
    if c.Cassandra.Password == "" {
        flag.StringVar(&c.Cassandra.Password, "cassandra.password", "", "Cassandra password.")
    }
    // Zookeeper
    c.Zookeeper.Nodes = strings.Split(*flag.String("zookeeper.nodes", strings.Join(c.Zookeeper.Nodes, ","), "Zookeeper nodes (comma delimited)."), ",")
    flag.StringVar(&c.Zookeeper.BasePath, "zookeeper.basepath", c.Zookeeper.BasePath, "Zookeeper base path.")
    // Kubernetes
    flag.StringVar(&c.Kubernetes.Host, "kubernetes.host", "", "Kubernetes host:port.")
    flag.StringVar(&c.Kubernetes.Namespace, "kubernetes.namespace", "photon", "Kubernetes namespace.")
    flag.BoolVar(&c.Kubernetes.Insecure, "kubernetes.insecure", false, "Access kubernetes master without verifying certificate (test only).")
    flag.StringVar(&c.Kubernetes.CertificateAuthority, "kubernetes.ca", "", "Kubernetes certificate authority.")
    flag.StringVar(&c.Kubernetes.CertificateFile, "kubernetes.cert", "", "Kubernetes certificate file.")
    flag.StringVar(&c.Kubernetes.KeyFile, "kubernetes.key", "", "Kubernetes key file.")

    flag.Parse()
}

func (c *Config) Print() {
    log.Info("Photon server")
    log.Infof("Environment: %s", c.Env)
    log.Infof("Server root: %s", c.Root)
    log.Infof("Loaded config: %s", c.ConfigPath)
}
