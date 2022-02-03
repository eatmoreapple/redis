package redis

type (
	LInsertOp int
	LOrderOp  int
	FlushOP   int
)

//go:generate stringer -type=LInsertOp,LOrderOp,FlushOP -linecomment=true -output=stringer.go

const (
	BEFORE LInsertOp = iota // BEFORE
	AFTER                   // AFTER

	ASC  LOrderOp = iota // ASC
	DESC                 // DESC

	ASYNC FlushOP = iota // ASYNC
	SYNC                 // SYNC
)
