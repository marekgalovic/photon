package main

import (
    "os";

    "github.com/marekgalovic/photon/cli/commands";

    "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
    app := kingpin.New("photon", "Photon CLI.")

    commands.NewModelsCommand().Register(app)

    app.Parse(os.Args[1:])
}
