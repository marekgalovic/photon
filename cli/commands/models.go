package commands

import (
    "gopkg.in/alecthomas/kingpin.v2"
)

type ModelsCommand struct {}

func NewModelsCommand() *ModelsCommand {
    return &ModelsCommand{}
}

func (mc *ModelsCommand) Register(app *kingpin.Application) {
    models := app.Command("models", "Manage models.")
    // List
    models.Command("list", "List models.")
    // Show
    show := models.Command("show", "Show model.")
    show.Arg("uid", "Uid of the model.").Required().String()
    // Delete
    delete := models.Command("delete", "Delete model.")
    delete.Arg("uid", "Uid of the model to delete.").Required().String()
}
