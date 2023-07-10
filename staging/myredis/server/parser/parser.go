package parser

import (
	"io"
	"myredis/face"
)

// Payload stores redis.Reply or error
type Payload struct {
	Data face.Reply
	Err  error
}

func ParseBytes(data []byte) ([]face.Reply, error) {

}

func ParseStream(reader io.Reader) <-chan *Payload {

}
