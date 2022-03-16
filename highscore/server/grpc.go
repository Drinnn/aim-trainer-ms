package server

import (
	"context"
	"fmt"
	"math"

	"github.com/Drinnn/aim-trainer-ms/highscore/protos"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type Grpc struct {
	address string
	server *grpc.Server
}

var Highscore = math.MaxFloat64

func (g *Grpc) SetHighscore(ctx context.Context, args *protos.SetHighscoreRequest) (*protos.SetHighscoreResponse, error) {
	log.Info().Msg(fmt.Sprintf("SetHighscore called - %e", args.Highscore))
	
	Highscore = args.Highscore
	
	return &protos.SetHighscoreResponse{
		Set: true,
	}, nil
}

func (g *Grpc) GetHighscore(ctx context.Context) (*protos.GetHighscoreResponse, error) {
	log.Info().Msg("GetHighscore called")
	
	return &protos.GetHighscoreResponse{
		Highscore: Highscore,
	}, nil
}