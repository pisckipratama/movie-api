package config

import (
	jwtware "github.com/gofiber/jwt/v3"
)

var Token []byte = []byte("secret")

var JWTConfig jwtware.Config = jwtware.Config{
	SigningKey: Token,
}
