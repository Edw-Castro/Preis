package utils

import "os"

var SecretKey string = os.Getenv("secret-key")

var SecretKeyClaims string = os.Getenv("secret_key")
