package main

import (
    "github.com/marekgalovic/serving/client_go";

    log "github.com/Sirupsen/logrus"
)

func main() {
    client, err := serving.NewClient(&serving.Config{
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

    // c := 0
    // for i := 0; i < 1000; i++ {
        

    //     if result["score"] == "B" {
    //         c += 1
    //     }
    // }
    // log.Info(c)
}
