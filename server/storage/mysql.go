package storage

import (
    "fmt";
    "database/sql";

    _ "github.com/go-sql-driver/mysql"
)

type MysqlConfig struct {
    User string
    Password string
    Host string
    Port int
    Database string
}

func (c *MysqlConfig) ConnectionUrl() string {
    return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True", c.User, c.Password, c.Host, c.Port, c.Database)
}

type Mysql struct {
    *sql.DB
}

type Scannable interface {
    Scan(...interface{}) error
}

type Countable interface {
    Count(string) (int, error)
}

func NewMysql(config MysqlConfig) (*Mysql, error) {
    conn, err := sql.Open("mysql", config.ConnectionUrl())
    if err != nil {
        return nil, err
    }

    if err = conn.Ping(); err != nil {
        return nil, err
    }

    return &Mysql{conn}, nil
}

func (mysql *Mysql) ExecPrepared(query string, args ...interface{}) (sql.Result, error) {
    stmt, err := mysql.Prepare(query)
    if err != nil {
        return nil, err
    }
    defer stmt.Close()

    return stmt.Exec(args...)
}

func (mysql *Mysql) QueryPrepared(query string, args ...interface{}) (*sql.Rows, error) {
    stmt, err := mysql.Prepare(query)
    if err != nil {
        return nil, err
    }
    defer stmt.Close()

    return stmt.Query(args...)
}

func (mysql *Mysql) QueryRowPrepared(query string, args ...interface{}) (*sql.Row, error) {
    stmt, err := mysql.Prepare(query)
    if err != nil {
        return nil, err
    }
    defer stmt.Close()

    return stmt.QueryRow(args...), nil
}

func (mysql *Mysql) Count(tableName string) (int, error) {
    var count int
    if err := mysql.QueryRow(fmt.Sprintf(`SELECT count(1) FROM %s`, tableName)).Scan(&count); err != nil {
        return 0, err
    }
    return count, nil
}
