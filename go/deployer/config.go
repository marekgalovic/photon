package deployer

import (
    "fmt";
    "flag";
    "time";
    "strings";
    "strconv";

    "github.com/marekgalovic/photon/go/core/utils";
    "github.com/marekgalovic/photon/go/core/storage"
)

type Config struct {
    Env string
    ModelsDir string
    ModelId int64
    Zookeeper storage.ZookeeperConfig
}

func NewConfig() (*Config, error) {
    envModelId, err := strconv.ParseInt(utils.GetEnv("PHOTON_MODEL_ID", "0"), 10, 64)
    if err != nil {
        return nil, err
    }

    config := &Config {
        Env: utils.GetEnv("PHOTON_ENV", "development"),
        ModelsDir: utils.GetEnv("PHOTON_MODELS_DIR", "./models"),
        ModelId: envModelId,
        Zookeeper: storage.ZookeeperConfig{
            Nodes: []string{"127.0.0.1:2181"},
            SessionTimeout: 1 * time.Second,
            BasePath: "/photon",
        },
    }

    config.parseFlags()

    if err := config.validate(); err != nil {
        return nil, err       
    }

    return config, nil
}

func (c *Config) parseFlags() {
    flag.StringVar(&c.Env, "env", c.Env, "Deployer environment.")
    flag.StringVar(&c.ModelsDir, "models-dir", c.ModelsDir, "Models directory.")
    flag.Int64Var(&c.ModelId, "model-id", c.ModelId, "Model id.")
    // Zookeeper
    c.Zookeeper.Nodes = strings.Split(*flag.String("zookeeper.nodes", strings.Join(c.Zookeeper.Nodes, ","), "Zookeeper nodes (comma delimited)."), ",")
    flag.StringVar(&c.Zookeeper.BasePath, "zookeeper.basepath", c.Zookeeper.BasePath, "Zookeeper base path.")

    flag.Parse()
}

func (c *Config) validate() error {
    if c.ModelId == 0 {
        return fmt.Errorf("Model id cannot be empty.")
    }
    if c.ModelsDir == "" {
        return fmt.Errorf("Models dir cannot be empty.")
    }
    return nil
}
