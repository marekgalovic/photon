package utils

import (
    "os";
)

func GetEnv(key string, defaultValue string) string {
    if envValue, exists := os.LookupEnv(key); exists {
        return envValue
    }
    return defaultValue
}
