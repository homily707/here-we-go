package protocol

// BulkReply stores a binary-safe string
type BulkReply struct {
	Arg []byte
}

func MakeMultiBulkReply(args [][]byte) *MultiBulkReply {
	return &MultiBulkReply{
		Args: args,
	}
}
