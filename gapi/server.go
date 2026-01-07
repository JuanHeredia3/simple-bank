package gapi

import (
	db "github.com/JuanHeredia3/simple-bank/db/sqlc"
	"github.com/JuanHeredia3/simple-bank/pb"
	"github.com/JuanHeredia3/simple-bank/token"
	"github.com/JuanHeredia3/simple-bank/util"
	"github.com/JuanHeredia3/simple-bank/worker"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
