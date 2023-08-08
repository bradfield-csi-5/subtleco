package vm

const (
	NOOP  = 0x00
	LOAD  = 0x01
	STORE = 0x02
	ADD   = 0x03
	SBT   = 0x04
	ADDI  = 0x05
	SUBI  = 0x06
	JUMP  = 0x07
	BEQZ  = 0x08
	HALT  = 0xff
)

func compute(memory []byte) {
	var registry [3]byte
	var currentInstruction, decodedInstruction [3]byte
	registry[0] = 8

	for {
		fetch(registry[:], memory, &currentInstruction)
		if currentInstruction[0] == HALT {
			return
		}
		decode(&currentInstruction, &registry, memory, &decodedInstruction)
		execute(memory, &decodedInstruction, &registry, currentInstruction[1])
	}
}

func fetch(registry []byte, memory []byte, currentInstruction *[3]byte) {
	currentInstruction[0] = memory[registry[0]]

	if currentInstruction[0] == HALT {
		return
	} else if currentInstruction[0] == JUMP {
		currentInstruction[1] = memory[registry[0]+1] // grab the value after the JMP command
	} else {
		currentInstruction[1] = memory[registry[0]+1]
		currentInstruction[2] = memory[registry[0]+2]
		registry[0] += 3
	}
}

func decode(instruction *[3]byte, registry *[3]byte, memory []byte, decodedInstruction *[3]byte) {
	decodedInstruction[0] = instruction[0]
	switch decodedInstruction[0] {
	case LOAD:
		decodedInstruction[1] = instruction[1]
		decodedInstruction[2] = memory[instruction[2]]
	case STORE:
		decodedInstruction[1] = registry[instruction[1]]
		decodedInstruction[2] = instruction[2]
	case ADD, SBT:
		decodedInstruction[1] = registry[instruction[1]]
		decodedInstruction[2] = registry[instruction[2]]
	case ADDI, SUBI:
		decodedInstruction[1] = registry[instruction[1]]
		decodedInstruction[2] = instruction[2]
	case JUMP:
		decodedInstruction[1] = instruction[1]
	case BEQZ:
		if registry[instruction[1]] == 0 {
			registry[0] += instruction[2]
		}
	default:
		return
	}
}

func execute(memory []byte, decodedInstruction *[3]byte, registry *[3]byte, registerAddress byte) {
	command := decodedInstruction[0]
	switch command {
	case LOAD:
		registry[decodedInstruction[1]] = decodedInstruction[2]
	case STORE:
		memory[decodedInstruction[2]] = decodedInstruction[1]
	case ADD, ADDI:
		registry[registerAddress] = decodedInstruction[1] + decodedInstruction[2]
	case SBT, SUBI:
		registry[registerAddress] = decodedInstruction[1] - decodedInstruction[2]
	case JUMP:
		registry[0] = decodedInstruction[1]
	default:
		return
	}
}
