package types

import "underwater/utils"

type Forward [utils.MAX_LVL]*Node

type Node struct {
	Key     []byte
	Value   []byte
	Forward Forward
}
