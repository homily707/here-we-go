package database

import "myredis/face"

func init() {
	RegisterCommand("ZAdd", execZAdd, writeFirstKey, undoZAdd, -4)
	RegisterCommand("ZScore", execZScore, readFirstKey, nil, 3)
	RegisterCommand("ZIncrBy", execZIncrBy, writeFirstKey, undoZIncr, 4)
	RegisterCommand("ZRank", execZRank, readFirstKey, nil, 3)
	RegisterCommand("ZCount", execZCount, readFirstKey, nil, 4)
	RegisterCommand("ZRevRank", execZRevRank, readFirstKey, nil, 3)
	RegisterCommand("ZCard", execZCard, readFirstKey, nil, 2)
	RegisterCommand("ZRange", execZRange, readFirstKey, nil, -4)
	RegisterCommand("ZRangeByScore", execZRangeByScore, readFirstKey, nil, -4)
	RegisterCommand("ZRevRange", execZRevRange, readFirstKey, nil, -4)
	RegisterCommand("ZRevRangeByScore", execZRevRangeByScore, readFirstKey, nil, -4)
	RegisterCommand("ZRem", execZRem, writeFirstKey, undoZRem, -3)
	RegisterCommand("ZRemRangeByScore", execZRemRangeByScore, writeFirstKey, rollbackFirstKey, 4)
	RegisterCommand("ZRemRangeByRank", execZRemRangeByRank, writeFirstKey, rollbackFirstKey, 4)
}

func execZAdd(db *DB, args [][]byte) face.Reply {}
