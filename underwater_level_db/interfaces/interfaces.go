package interfaces

type DB interface {
	// returns value for given key, error if not found
	Get(key []byte) (value []byte, err error)

	Has(key []byte) (ret bool, err error)

	Put(key, value []byte) error

	Delete(key []byte) error

	RangeScan(start, limit []byte) (Iterator, error)
}

type Iterator interface {
	// Moves iterator to next k/v pair. Returns false if exhausted
	Next() bool

	// Error returns any accumulated Error. Exhaustion is not an error
	Error() error

	// Returns key of current pair, or nil if done
	Key() []byte

	// Returns value of current pair, or nil if done
	Value() []byte
}

type ImmutableDB interface {
	// returns value for given key, error if not found
	Get(key []byte) (value []byte, err error)

	Has(key []byte) (ret bool, err error)

	RangeScan(start, limit []byte) (Iterator, error)
}
