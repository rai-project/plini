package localqueue

import (
	queue "github.com/jakewins/4fq/pkg/queue"
)

// Create a new queue that safely handles multiple producers publishing items,
// and multiple consumer receiving them.
//
// Options are, as implied, optional. The queue defaults to 64 slots fixed size,
// and initializes the Val on each slot to nil.
func NewMultiProducerMultiConsumer(opts queue.Options) (queue.Queue, error) {
	return queue.NewMultiProducerMultiConsumer(opts)
}

// Create a new queue that safely handles multiple producers publishing items,
// and one consumer receiving them. Note that the onus is on you to ensure there
// is just one consumer - the queue will do crazy things if multiple consumers
// run concurrently.
//
// Options are, as implied, optional. The queue defaults to 64 slots fixed size,
// and initializes the Val on each slot to nil.
func NewMultiProducerSingleConsumer(opts queue.Options) (queue.Queue, error) {
	return queue.NewMultiProducerSingleConsumer(opts)
}

// Create a new queue that safely handles one producer publishing items,
// and multiple consumer receiving them.
//
// Options are, as implied, optional. The queue defaults to 64 slots fixed size,
// and initializes the Val on each slot to nil.
func NewSingleProducerMultiConsumer(opts queue.Options) (queue.Queue, error) {
	return queue.NewSingleProducerMultiConsumer(opts)
}

// Create a new queue that safely handles one producer publishing items,
// and one consumer receiving them. Note that the onus is on you to ensure there
// is just one consumer and one producer - the queue will do crazy things if this
// rule is broken.
//
// Options are, as implied, optional. The queue defaults to 64 slots fixed size,
// and initializes the Val on each slot to nil.
func NewSingleProducerSingleConsumer(opts queue.Options) (queue.Queue, error) {
	return queue.NewSingleProducerSingleConsumer(opts)
}
