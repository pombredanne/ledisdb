package ledis

import (
	"errors"
)

const (
	kvType byte = iota + 1
	hashType
	hSizeType
	listType
	lMetaType
	zsetType
	zSizeType
	zScoreType

	kvExpType
	kvExpMetaType
	lExpType
	lExpMetaType
	hExpType
	hExpMetaType
	zExpType
	zExpMetaType
)

const (
	defaultScanCount int = 10
)

var (
	errKeySize        = errors.New("invalid key size")
	errValueSize      = errors.New("invalid value size")
	errHashFieldSize  = errors.New("invalid hash field size")
	errZSetMemberSize = errors.New("invalid zset member size")
	errExpireValue    = errors.New("invalid expire value")
)

const (
	//we don't support too many databases
	MaxDBNumber uint8 = 16

	//max key size
	MaxKeySize int = 1024

	//max hash field size
	MaxHashFieldSize int = 1024

	//max zset member size
	MaxZSetMemberSize int = 1024

	//max value size
	MaxValueSize int = 10 * 1024 * 1024
)

var (
	ErrKeySize        = errors.New("invalid key size")
	ErrHashFieldSize  = errors.New("invalid hash field size")
	ErrZSetMemberSize = errors.New("invalid zset member size")
	ErrScoreMiss      = errors.New("zset score miss")
)

const (
	BinLogTypeDeletion uint8 = 0x0
	BinLogTypePut      uint8 = 0x1
	BinLogTypeCommand  uint8 = 0x2
)
