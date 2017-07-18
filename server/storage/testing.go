package storage

import (
    // "fmt";
    // "testing";

    log "github.com/Sirupsen/logrus"
)

func NewTestMysql() *Mysql {
    mysql, _ := NewMysql("root:@tcp(127.0.0.1:3306)/serving_test?parseTime=True")
    return mysql
}

func CleanupMysql(db *Mysql) {
    db.Exec(`DELETE FROM models`)
    db.Exec(`DELETE FROM feature_sets`)
}

func AssertCountChanged(db *Mysql, tableName string, diff int, callable func()) {
    expected, _ := db.Count(tableName)
    expected += diff

    callable()
    
    actual, _ := db.Count(tableName)

    if expected != actual {
        log.Fatalf("Expected count %d != %d", expected, actual)
    }
}
