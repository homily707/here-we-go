package server

import (
	"context"
	"net"
	"sync"
)

type Handler struct {
	activeConn sync.Map // *client -> placeholder
	db         database.DB
	closing    atomic.Boolean // refusing new client and new request
}

func (h Handler) Handle(ctx context.Context, conn net.Conn) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) Close(err error) {
	//TODO implement me
	panic("implement me")
}
