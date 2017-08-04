package photon

import (
    "fmt";
)

type Config struct {
    Address string
    Port int
}

func (c *Config) serverAddr() string {
    return fmt.Sprintf("%s:%d", c.Address, c.Port)
}
