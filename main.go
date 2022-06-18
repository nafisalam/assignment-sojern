package main

import (
	"Users/nafisalam/go/projects/assignment-sojern/config"
	"Users/nafisalam/go/projects/assignment-sojern/server"
	"Users/nafisalam/go/projects/assignment-sojern/util"
	"log"
	"net/http"
)

func main() {
	cfg := config.InitConfig("./.conf")
	util.CreateBanner(cfg.BannerPath)
	err := util.CreateCache(cfg.ImagePath)
	if err != nil {
		log.Println("error to create the cache ..", err)
	}
	r := server.NewRouter(cfg)
	log.Fatal(http.ListenAndServe(":8080", r))
}
