package database

import (
	"golang.org/x/net/dict"
	"myredis/face"
	"sync"
	"time"
)

type DB struct {
	index int
	// key -> DataEntity
	data dict.Dict
	// key -> expireTime (time.Time)
	ttlMap dict.Dict
	// key -> version(uint32)
	versionMap dict.Dict

	// dict.Dict will ensure concurrent-safety of its method
	// use this mutex for complicated command only, eg. rpush, incr ...
	locker *lock.Locks
	// stop all data access for execFlushDB
	stopWorld sync.WaitGroup
	addAof    func(CmdLine)
}

func (db *DB) Exec(client face.Connection, cmdLine [][]byte) face.Reply {
	//TODO implement me
	panic("implement me")
}

func (db *DB) AfterClientClose(c face.Connection) {
	//TODO implement me
	panic("implement me")
}

func (db *DB) Close() {
	//TODO implement me
	panic("implement me")
}

func (db *DB) ExecWithLock(conn face.Connection, cmdLine [][]byte) face.Reply {
	//TODO implement me
	panic("implement me")
}

func (db *DB) ExecMulti(conn face.Connection, watching map[string]uint32, cmdLines []face.CmdLine) face.Reply {
	//TODO implement me
	panic("implement me")
}

func (db *DB) GetUndoLogs(dbIndex int, cmdLine [][]byte) []face.CmdLine {
	//TODO implement me
	panic("implement me")
}

func (db *DB) ForEach(dbIndex int, cb func(key string, data *face.DataEntity, expiration *time.Time) bool) {
	//TODO implement me
	panic("implement me")
}

func (db *DB) RWLocks(dbIndex int, writeKeys []string, readKeys []string) {
	//TODO implement me
	panic("implement me")
}

func (db *DB) RWUnLocks(dbIndex int, writeKeys []string, readKeys []string) {
	//TODO implement me
	panic("implement me")
}

func (db *DB) GetDBSize(dbIndex int) (int, int) {
	//TODO implement me
	panic("implement me")
}

type ExecFunc func(db *DB, args [][]byte) redis.Reply

// PreFunc analyses command line when queued command to `multi`
// returns related write keys and read keys
type PreFunc func(args [][]byte) ([]string, []string)

// CmdLine is alias for [][]byte, represents a command line
type CmdLine = [][]byte

// UndoFunc returns undo logs for the given command line
// execute from head to tail when undo
type UndoFunc func(db *DB, args [][]byte) []CmdLine

func (db *DB) Exec(c redis.Connection, cmdLine [][]byte) redis.Reply {

func (D DB) Exec(client face.Connection, cmdLine [][]byte) face.Reply {
	 
}

func (D DB) AfterClientClose(c face.Connection) {
}

func (D DB) Close() {
}
