package main

import (
	"flag"

	"github.com/rs/zerolog/log"

	"github.com/Drinnn/aim-trainer-ms/highscore/server"
)

func main() {
	var address = flag.String("address", ":5051", "address where you can connect with highscore service")
	flag.Parse()

	server := server.NewGrpcServer(*address)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start gRPC server of highscore service")
	}
}