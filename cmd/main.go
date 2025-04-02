package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"untitled5"
	"untitled5/internal/http"
	"untitled5/internal/store"
)

func main() {
	log.Info().Msg("Program start")

	var err error
	var config *untitled5.Config
	log.Info().Msg("Parse config")
	if config, err = untitled5.ParseConfig("./template/config.yaml"); err != nil {panic(err)}
	log.Info().Msg("With config good")

	var mainStore *store.Store
	log.Info().Msg("Connecting DB")
	if mainStore, err = store.ConnectDB(config.Store); err != nil {panic(err)}
	log.Info().Msg("DB good")

	network := &http.Network{Store: mainStore}

	router := gin.Default()
	router.GET("/get-meets", network.GetMeets)
	router.POST("/get-meet", network.GetMeetById)
	//router.POST("/create-meet", network.CreateMeet)

	router.Run()
}