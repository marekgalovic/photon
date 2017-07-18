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

func (mysql *Mysql) QueryDynamic(query string, args ...interface{}) ([]map[string]interface{}, error) {
    stmt, err := mysql.Prepare(query)
    if err != nil {
        return nil, err
    }
    defer stmt.Close()

    rows, err := stmt.Query(args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    columns, err := rows.ColumnTypes()
    if err != nil {
        return nil, err
    }

    result := make([]map[string]interface{}, 0)
    placeholders := mysql.blankRow(len(columns))
    for rows.Next() {
        rows.Scan(placeholders...)

        resultRow := map[string]interface{}{}
        for i, column := range columns {
            resultRow[column.Name()] = *(placeholders[i].(*interface{}))
        }
        result = append(result, resultRow)
    }

    return result, nil
}

func (mysql *Mysql) Count(tableName string) (int, error) {
    var count int
    if err := mysql.QueryRow(fmt.Sprintf(`SELECT count(1) FROM %s`, tableName)).Scan(&count); err != nil {
        return 0, err
    }
    return count, nil
}

func (mysql *Mysql) ShowTables() ([]string, error) {
    rows, err := mysql.Query(`SHOW TABLES`)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    tables := make([]string, 0)
    for rows.Next() {
        var tableName string
        if err = rows.Scan(&tableName); err != nil {
            return nil, err
        }
        tables = append(tables, tableName)
    }
    return tables, nil
}

func (mysql *Mysql) blankRow(size int) []interface{} {
    placeholders := make([]interface{}, size)
    values := make([]interface{}, size)
    for i := 0; i < size; i++ {
       placeholders[i] = &values[i]
    }
    return placeholders
}
