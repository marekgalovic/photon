package main

import (
    // "os";

    "github.com/marekgalovic/photon/go/client";
    // pb "github.com/marekgalovic/photon/go/core/protos";

    log "github.com/Sirupsen/logrus"
)

func main() {
    client, err := photon.NewClient(&photon.Config{Address: "127.0.0.1", Port: 5005}, photon.NewCredentials("my-key", "my-secret"))
    if err != nil {
        log.Fatal(err)
    }

    // versions, err := client.ListModels()
    // if err != nil {
    //     log.Fatal(err)
    // }
    // log.Info(versions)

    // modelFile, err := os.Open("./../../../../runners/pmml/models/test.xml")
    // if err != nil {
    //     log.Fatal(err)
    // }

    // versionId, err := client.CreateModelVersion(1, "my-version", true, false, modelFile)
    // if err != nil {
    //     log.Fatal(err)
    // }

    // log.Info(versionId)

    // modelId, err := client.CreateModel("rf25", "pmml", 1, []*pb.ModelFeature{{Name: "x1", Required: true}}, nil)
    // if err != nil {
    //     log.Fatal(err)
    // }
    // log.Info(modelId)

    log.Info(client.DeleteModel(2))

    // result, err := client.Evaluate("x000", map[string]interface{}{"foo": 1, "bar": 0.3})
    // if err != nil {
    //     log.Fatal(err)
    // }

    // log.Info(result)
}
