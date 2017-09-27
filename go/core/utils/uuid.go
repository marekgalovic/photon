package utils

import (
    "fmt";

    "github.com/satori/go.uuid";
)

func UuidV1() string {
    return fmt.Sprintf("%s", uuid.NewV1())
}

func UuidV4() string {
    return fmt.Sprintf("%s", uuid.NewV4())
}
