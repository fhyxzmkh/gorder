package main

import (
	"github.com/fhyxzmkh/gorder/common/config"
	"github.com/spf13/viper"
	"log"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Println(viper.Get("order"))
}
