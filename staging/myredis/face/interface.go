package face

import (
	"context"
	"net"
	"time"
)

type HandleFunc func(ctx context.Context, conn net.Conn)

type Handler interface {
	Handle(ctx context.Context, conn net.Conn)
	Close(error)
}

type CmdLine = [][]byte

// DB is the face for redis style storage engine
type DB interface {
	Exec(client Connection, cmdLine [][]byte) Reply
	AfterClientClose(c Connection)
	Close()
}

// EmbedDB is the embedding storage engine exposing more methods for complex application
type EmbedDB interface {
	DB
	ExecWithLock(conn Connection, cmdLine [][]byte) Reply
	ExecMulti(conn Connection, watching map[string]uint32, cmdLines []CmdLine) Reply
	GetUndoLogs(dbIndex int, cmdLine [][]byte) []CmdLine
	ForEach(dbIndex int, cb func(key string, data *DataEntity, expiration *time.Time) bool)
	RWLocks(dbIndex int, writeKeys []string, readKeys []string)
	RWUnLocks(dbIndex int, writeKeys []string, readKeys []string)
	GetDBSize(dbIndex int) (int, int)
}

// DataEntity stores data bound to a key, including a string, list, hash, set and so on
type DataEntity struct {
	Data interface{}
}

type Connection interface {
	Write([]byte) error
	SetPassword(string)
	GetPassword() string

	// server should keep its subscribing channels
	Subscribe(channel string)
	UnSubscribe(channel string)
	SubsCount() int
	GetChannels() []string

	// used for `Multi` command
	InMultiState() bool
	SetMultiState(bool)
	GetQueuedCmdLine() [][][]byte
	EnqueueCmd([][]byte)
	ClearQueuedCmds()
	GetWatching() map[string]uint32

	// used for multi database
	GetDBIndex() int
	SelectDB(int)
}

type Reply interface {
	ToBytes() []byte
}
