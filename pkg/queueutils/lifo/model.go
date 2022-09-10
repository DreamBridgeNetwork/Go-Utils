package lifo

import "github.com/DreamBridgeNetwork/Go-Utils/pkg/queueutils/roundqueue"

type Lifo struct {
	queue *roundqueue.RoundQueue
}
