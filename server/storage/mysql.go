package storage

import (
    "fmt";
    "database/sql";

    _ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
    *sql.DB
}

type Scannable interface {
    Scan(...interface{}) error
}

func NewMysql(connUrl string) (*Mysql, error) {
    conn, err := sql.Open("mysql", connUrl)
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
