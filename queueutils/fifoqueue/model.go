package fifoqueue

import (
	"sync"

	"github.com/DreamBridgeNetwork/Go-Utils/queueutils"
)

type FifoQueue struct {
	mu          sync.Mutex
	firstBlock  *queueutils.Block
	lastBlock   *queueutils.Block
	freePointer *queueutils.Block
	size        uint
	maxSize     uint
}
