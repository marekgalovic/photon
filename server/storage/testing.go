package storage

import (
    log "github.com/Sirupsen/logrus"
)

func NewTestMysql() *Mysql {
    mysql, err := NewMysql(MysqlConfig{Host: "127.0.0.1", Port: 3306, User: "root", Database: "serving_test"})
    if err != nil {
        log.Fatal(err)
    }
    return mysql
}

func CleanupMysql(db *Mysql) {
    db.Exec(`DELETE FROM models`)
    db.Exec(`DELETE FROM feature_sets`)
}

func AssertCountChanged(db *Mysql, tableName string, diff int, callable func()) {
    expected, err := db.Count(tableName)
    if err != nil {
        log.Fatal(err)
    }
    expected += diff

    callable()
    
    actual, _ := db.Count(tableName)

    if expected != actual {
        log.Fatalf("Expected count %d != %d", expected, actual)
    }
}
