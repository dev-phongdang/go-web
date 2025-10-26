package main

import (
	"log"

	"aidanwoods.dev/go-paseto"
)

func main() {
	private_key := paseto.NewV4AsymmetricSecretKey()
	private_byte := private_key.ExportHex()
	log.Println(private_byte)
}
