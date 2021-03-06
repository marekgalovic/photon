package main

import (
    "os";
    "fmt";
    "time";
    "path/filepath";

    "github.com/marekgalovic/photon/go/server";

    "github.com/mattes/migrate";
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/mattes/migrate/database/mysql"
    _ "github.com/mattes/migrate/source/file"
    log "github.com/Sirupsen/logrus"
)

const (
    migrationsPath = "./go/core/db/migrations"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatal("No command provided.")
    }

    config, err := server.NewConfig()
    if err != nil {
        log.Fatal()
    }

    migrator, err := migrate.New(fmt.Sprintf("file://%s", migrationsPath), fmt.Sprintf("mysql://%s", config.Mysql.ConnectionUrl()))
    if err != nil {
        log.Fatalf("Blah %v", err)
    }

    switch os.Args[1] {
    case "up":
        if err = migrator.Up(); err != nil {
            log.Fatal(err)
        }
    case "down":
        if err = migrator.Down(); err != nil {
            log.Fatal(err)
        }
    case "drop":
        if err = migrator.Drop(); err != nil {
            log.Fatal(err)
        }
    case "create":
        if len(os.Args) < 3 {
            log.Fatal("No migration name provided.")
        }

        name := fmt.Sprintf("%d_%s", time.Now().Unix(), os.Args[2])
        upFile, _ := filepath.Abs(filepath.Join(migrationsPath, fmt.Sprintf("%s.up.sql", name)))
        downFile, _ := filepath.Abs(filepath.Join(migrationsPath, fmt.Sprintf("%s.down.sql", name)))

        if _, err := os.Create(upFile); err != nil {
            log.Fatal(err)
        }
        log.Infof("Created file: %s", upFile)
        if _, err := os.Create(downFile); err != nil {
            log.Fatal(err)
        }
        log.Infof("Created file: %s", downFile)    
    default:
        log.Fatalf("Unknown command '%s'", os.Args[1])
    }
}
