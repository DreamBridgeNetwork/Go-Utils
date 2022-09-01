package roundqueue

import (
	"sync"

	"github.com/DreamBridgeNetwork/Go-Utils/queueutils"
)

type RoundQueue struct {
	mu      sync.Mutex
	pointer *queueutils.Block
	size    int
	maxSize int
}
