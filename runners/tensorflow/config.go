package runner

import (
    "os";
    "fmt";
    "flag";
)

type Config struct {
    Env string
    ModelUid string
    ModelsDir string
    Address string
    Port int
}

func NewConfig() (*Config, error) {
    config := &Config{}

    config.Env = config.getEnvDefault("PHOTON_ENV", "development")
    config.ModelUid = config.getEnvDefault("PHOTON_MODEL_UID", "")
    config.ModelsDir = config.getEnvDefault("PHOTON_MODELS_DIR", "./")

    config.parseFlags()

    if err := config.validate(); err != nil {
        return nil, err
    }

    return config, nil
}

func (c *Config) parseFlags() {
    flag.StringVar(&c.Env, "env", c.Env, "Server environment.")
    flag.StringVar(&c.ModelUid, "model-uid", c.ModelUid, "Model uid.")
    flag.StringVar(&c.ModelsDir, "models-dir", c.ModelsDir, "Models directory.")
    flag.StringVar(&c.Address, "address", "0.0.0.0", "Server address.")
    flag.IntVar(&c.Port, "port", 5022, "Server port.")

    flag.Parse()
}

func (c *Config) validate() error {
    if c.ModelUid == "" {
        return fmt.Errorf("No model uid provided.")
    }
    return nil
}

func (c *Config) getEnvDefault(key string, defaultValue string) string {
    if envValue, exists := os.LookupEnv(key); exists {
        return envValue
    }
    return defaultValue
}

func (c *Config) NodeIp() (string, error) {
    return c.Address, nil
}
