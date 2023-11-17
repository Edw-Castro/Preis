package utils

import "os"

var SecretKey string = os.Getenv("secret-key")
