package main

import (
    "github.com/marekgalovic/photon/go/client";

    log "github.com/Sirupsen/logrus"
)

func main() {
    client, err := photon.NewClient(&photon.Config{
        Address: "127.0.0.1",
        Port: 5005,
    })
    if err != nil {
        log.Fatal(err)
    }

    result, err := client.Evaluate("x000", map[string]interface{}{"foo": 1, "bar": 0.3})
    if err != nil {
        log.Fatal(err)
    }

    log.Info(result)
}
