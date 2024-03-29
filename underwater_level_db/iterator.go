package main

type RSIterator struct {
	entries []Entry
	index   int
}

func (i *RSIterator) Next() bool {
	if i.index+1 < len(i.entries) {
		i.index++
		return true
	}
	return false
}

func (i *RSIterator) Error() error {
	return nil
}

func (i *RSIterator) Key() []byte {
	return i.entries[i.index].key
}

func (i *RSIterator) Value() []byte {
	return i.entries[i.index].value
}
