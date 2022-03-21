package server

import (
	"context"
	"fmt"
	"math"
	"net"

	"github.com/Drinnn/aim-trainer-ms/highscore/protos"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Grpc struct {
	address string
	server *grpc.Server
}

func NewGrpcServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

var Highscore = math.MaxFloat64

func (g *Grpc) SetHighscore(ctx context.Context, args *protos.SetHighscoreRequest) (*protos.SetHighscoreResponse, error) {
	log.Info().Msg(fmt.Sprintf("SetHighscore called - %e", args.Highscore))
	
	Highscore = args.Highscore
	
	return &protos.SetHighscoreResponse{
		Set: true,
	}, nil
}

func (g *Grpc) GetHighscore(context.Context, *emptypb.Empty) (*protos.GetHighscoreResponse, error) {
	log.Info().Msg("GetHighscore called")
	
	return &protos.GetHighscoreResponse{
		Highscore: Highscore,
	}, nil
}

func (g *Grpc) ListenAndServe() error {
	listener, err := net.Listen("tcp", g.address)
	if err != nil {
		return errors.Wrap(err, "failed to open tcp port")
	}

	serverOpts := []grpc.ServerOption{}
	g.server = grpc.NewServer(serverOpts...)

	protos.RegisterHighscoreServer(g.server, g)

	log.Info().Str("address", g.address).Msg("starting gRPC server for higscore service")

	err = g.server.Serve(listener)
	if err != nil {
		return errors.Wrap(err, "failed to start gRPC server for highscore service")
	}
	
	return nil
}