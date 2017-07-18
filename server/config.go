package server

import (
    "fmt";
)

type Config struct {
    Address string
    Port int
    Env string
    Root string
    Mysql *MysqlConfig
}

type MysqlConfig struct {
    User string
    Password string
    Host string
    Port int
    Database string
}

func NewConfig() *Config {
    return &Config{
        Port: 5005,
        Root: "./",
        Env: "development",
        Mysql: &MysqlConfig{
            User: "root",
            Host: "127.0.0.1",
            Port: 3306,
            Database: "serving_development",
        },
    }
}

func (c *Config) BindAddress() string {
    return fmt.Sprintf("%s:%d", c.Address, c.Port)
}

func (c *MysqlConfig) ConnectionUrl() string {
    return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", c.User, c.Password, c.Host, c.Port, c.Database)
}
