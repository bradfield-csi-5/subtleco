package interfaces

type DB interface {
	Get(key []byte) (inx int, value []byte, err error)

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
