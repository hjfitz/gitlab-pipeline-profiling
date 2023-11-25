package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hjfitz/gitlab-pipeline-profiling/controllers"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	r := gin.Default()

	r.Use(loggerMiddleware)
	r.Use(cors.Default())

	apiControllers := controllers.BuildControllers()

	for _, controller := range apiControllers {
		controller.ApplyRoutes(r)
	}

	if err := r.Run(); err != nil {
		panic(err)
	} else {
		log.Info().Msg("BFF started")
	}
}
