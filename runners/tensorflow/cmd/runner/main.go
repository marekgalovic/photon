package main

import (
    "github.com/marekgalovic/photon/runners/tensorflow";

    log "github.com/Sirupsen/logrus"
)

func main() {
    config, err := runner.NewConfig()
    if err != nil {
        log.Fatal(err)
    }

    log.Info(config)
}
