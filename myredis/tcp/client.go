package tcp

import (
	"myredis/face"
	"net"
	"sync"
	"time"
)

type Client struct {
	conn        net.Conn
	pendingReqs chan *request // wait to send
	waitingReqs chan *request // waiting response
	ticker      *time.Ticker
	addr        string

	working *sync.WaitGroup // its counter presents unfinished requests(pending and waiting)
}

// request is a message sends to redis server
type request struct {
	id        uint64
	args      [][]byte
	reply     face.Reply
	heartbeat bool
	//waiting   *wait.Wait
	err error
}

func MakeClient(addr string) (*Client, error) {
	return nil, nil
}

func (client *Client) Start() {

}

func (client *Client) Close() {}

func (client *Client) Send(args [][]byte) face.Reply {
	return nil
}
