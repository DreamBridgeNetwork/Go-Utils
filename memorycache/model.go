package memorycache

import "github.com/DreamBridgeNetwork/Go-Utils/queueutils/roundqueue"

// MemoryCache - Struct to manage memory cache
type MemoryCache struct {
	DataType         map[string]*roundqueue.RoundQueue
	maxTypeCacheSize int
}
