package runner

import (
    "os";
    "fmt";
    "flag";
    "time";
    "strings";
    "net";

    "github.com/marekgalovic/photon/server/storage"
)

type Config struct {
    Env string
    ModelUid string
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
        Env: getEnvDefault("PHOTON_ENV", "development"),
        ModelUid: getEnvDefault("PHOTON_MODEL_UID", ""),
        ModelsDir: getEnvDefault("PHOTON_MODELS_DIR", "./"),
        Port: 5022,
        Zookeeper: storage.ZookeeperConfig{
            Nodes: []string{"127.0.0.1:2181"},
            SessionTimeout: 1 * time.Second,
            BasePath: "/photon",
        },
    }

    nodeIp, err := config.NodeIp()
    if err != nil {
        return nil, err
    }
    config.Address = nodeIp

    if err := config.parseFlags(); err != nil {
        return nil, err
    }

    return config, nil
}

func (c *Config) parseFlags() error {
    flag.StringVar(&c.Env, "env", c.Env, "Server environment.")
    flag.StringVar(&c.ModelUid, "model-uid", c.ModelUid, "Model uid.")
    flag.StringVar(&c.ModelsDir, "models-dir", c.ModelsDir, "Models directory.")
    // Listener
    flag.StringVar(&c.Address, "address", c.Address, "Server address.")
    flag.IntVar(&c.Port, "port", 5022, "Server port.")
    // Zookeeper
    c.Zookeeper.Nodes = strings.Split(*flag.String("zookeeper-nodes", strings.Join(c.Zookeeper.Nodes, ","), "Zookeeper nodes (comma delimited)."), ",")
    flag.StringVar(&c.Zookeeper.BasePath, "zookeeper-basepath", c.Zookeeper.BasePath, "Zookeeper base path.")

    flag.Parse()
    return c.validate()
}

func (c *Config) validate() error {
    // if c.ModelUid == "" {
    //     return fmt.Errorf("No model uid provided.")
    // }
    return nil
}

func (c *Config) NodeIp() (string, error) {
    interfaces, err := net.Interfaces()
    if err != nil {
        return "", err
    }
    for _, iface := range interfaces {
        if iface.Flags & net.FlagUp == 0 {
            continue
        }
        if iface.Flags & net.FlagLoopback != 0 {
            continue
        }

        addrs, err := iface.Addrs()
        if err != nil {
            return "", err
        }

        for _, addr := range addrs {
            var ip net.IP
            switch v := addr.(type) {
                case *net.IPNet:
                    ip = v.IP
                case *net.IPAddr:
                    ip = v.IP
            }

            if ip == nil || ip.IsLoopback() {
                continue
            }
            ip = ip.To4()
            if ip == nil {
                continue
            }

            return ip.String(), nil
        }
    }
    return "", fmt.Errorf("Unable to find node IP.")
}
