package connection

import (
	"net"
	"sync"
)

// implent the face connection

type Connection struct {
	conn net.Conn

	// waiting until protocol finished
	waitingReply wait.Wait

	// lock while server sending response
	mu sync.Mutex

	// subscribing channels
	subs map[string]bool

	// password may be changed by CONFIG command during runtime, so store the password
	password string

	// queued commands for `multi`
	multiState bool
	queue      [][][]byte
	watching   map[string]uint32

	// selected db
	selectedDB int
}

func (c Connection) Write(bytes []byte) error {
	//TODO implement me
	panic("implement me")
}

func (c Connection) SetPassword(s string) {
	//TODO implement me
	panic("implement me")
}

func (c Connection) GetPassword() string {
	//TODO implement me
	panic("implement me")
}

func (c Connection) Subscribe(channel string) {
	//TODO implement me
	panic("implement me")
}

func (c Connection) UnSubscribe(channel string) {
	//TODO implement me
	panic("implement me")
}

func (c Connection) SubsCount() int {
	//TODO implement me
	panic("implement me")
}

func (c Connection) GetChannels() []string {
	//TODO implement me
	panic("implement me")
}

func (c Connection) InMultiState() bool {
	//TODO implement me
	panic("implement me")
}

func (c Connection) SetMultiState(b bool) {
	//TODO implement me
	panic("implement me")
}

func (c Connection) GetQueuedCmdLine() [][][]byte {
	//TODO implement me
	panic("implement me")
}

func (c Connection) EnqueueCmd(i [][]byte) {
	//TODO implement me
	panic("implement me")
}

func (c Connection) ClearQueuedCmds() {
	//TODO implement me
	panic("implement me")
}

func (c Connection) GetWatching() map[string]uint32 {
	//TODO implement me
	panic("implement me")
}

func (c Connection) GetDBIndex() int {
	//TODO implement me
	panic("implement me")
}

func (c Connection) SelectDB(i int) {
	//TODO implement me
	panic("implement me")
}
