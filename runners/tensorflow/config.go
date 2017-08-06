package runner

import (
    "os";
    "fmt";
    "flag";
    "time";
    "strings";

    "github.com/marekgalovic/photon/go/core/utils";
    "github.com/marekgalovic/photon/go/core/storage"
)

type Config struct {
    Env string
    ModelsDir string
    Address string
    Port int
    Zookeeper storage.ZookeeperConfig
}

func getEnvDefault(key string, defaultValue string) string {
    if envValue, exists := os.LookupEnv(key); exists {
        return envValue
    }
    return defaultValue
}

func NewConfig() (*Config, error) {
    config := &Config{
        Env: utils.GetEnv("PHOTON_ENV", "development"),
        ModelsDir: utils.GetEnv("PHOTON_MODELS_DIR", "./"),
        Port: 5022,
        Zookeeper: storage.ZookeeperConfig{
            Nodes: []string{"127.0.0.1:2181"},
            SessionTimeout: 1 * time.Second,
            BasePath: "/photon",
        },
    }

    nodeIp, err := utils.NodeIp()
    if err != nil {
        return nil, err
    }
    config.Address = nodeIp
    config.parseFlags()

    return config, nil
}

func (c *Config) parseFlags() {
    flag.StringVar(&c.Env, "env", c.Env, "Server environment.")
    flag.StringVar(&c.ModelsDir, "models-dir", c.ModelsDir, "Models directory.")
    // Listener
    flag.StringVar(&c.Address, "address", c.Address, "Server address.")
    flag.IntVar(&c.Port, "port", 5022, "Server port.")
    // Zookeeper
    c.Zookeeper.Nodes = strings.Split(*flag.String("zookeeper-nodes", strings.Join(c.Zookeeper.Nodes, ","), "Zookeeper nodes (comma delimited)."), ",")
    flag.StringVar(&c.Zookeeper.BasePath, "zookeeper-basepath", c.Zookeeper.BasePath, "Zookeeper base path.")

    flag.Parse()
}

func (c *Config) BindAddress() string {
    return fmt.Sprintf("%s:%d", c.Address, c.Port)
}
