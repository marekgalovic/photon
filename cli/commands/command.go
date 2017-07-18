package commands

import (
    "gopkg.in/alecthomas/kingpin.v2"
)

type CommandInterface struct {
    Register(*kingpin.Application)
}
