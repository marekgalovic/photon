package main

import (
    "os";
    "fmt";

    "gopkg.in/urfave/cli.v1"
)

func main() {
    app := cli.NewApp()
    app.Name = "photon"
    app.Usage = ""
    app.Description = "Photon CLI."

    app.Commands = []cli.Command{
        {
            Name: "models",
            Usage: "Manage models.",
            Subcommands: []cli.Command{
                {
                    Name: "list",
                    Usage: "List models.",
                },
                {
                    Name: "create",
                    Usage: "Create model.",
                },
                {
                    Name: "delete",
                    Usage: "Delete model.",
                },
                {
                    Name: "versions",
                    Usage: "Show model versions.",
                    ArgsUsage: "[model-uid]",
                },
            },
        },
    }


    app.Run(os.Args)
}
