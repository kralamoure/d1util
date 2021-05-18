package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/pflag"

	"github.com/kralamoure/retroutil"
)

var (
	data string
	key  string
)

func main() {
	pflag.StringVarP(&data, "data", "d", "", "The encrypted data")
	pflag.StringVarP(&key, "key", "k", "", "The key")

	pflag.Parse()

	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	result, err := retroutil.DecipherGameMap(data, key)
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}
