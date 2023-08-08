package vm

const (
	LOAD  = 0x01
	STORE = 0x02
	ADD   = 0x03
	SBT   = 0x04
	HALT  = 0xff
)

const (
	DATA  = 0
	INSTR = 8
)

func compute(memory []byte) {
	var registry [2]byte
	var currentInstruction [3]byte
	PC := 0

	for true {
		currentInstruction = fetch(&PC, memory)
		if currentInstruction[0] == HALT {
			return
		}
		decodedInstruction := decode(currentInstruction, registry, memory)
		registry = execute(memory, decodedInstruction, registry, currentInstruction[1])
	}
}

func fetch(PC *int, memory []byte) [3]byte {
	var instructions [3]byte
	instructions[0] = memory[*PC+INSTR]

	if instructions[0] == HALT {
		instructions[1], instructions[2] = 0, 0
		*PC++
	} else {
		instructions[1] = memory[*PC+INSTR+1]
		instructions[2] = memory[*PC+INSTR+2]
		*PC += 3
	}
	return instructions
}

func decode(instruction [3]byte, registry [2]byte, memory []byte) [3]byte {
	var decodedInstruction [3]byte
	decodedInstruction[0] = instruction[0]
	switch decodedInstruction[0] {
	case LOAD:
		decodedInstruction[1] = instruction[1]
		decodedInstruction[2] = memory[instruction[2]]
	case STORE:
		decodedInstruction[1] = registry[instruction[1]-1]
	case ADD, SBT:
		decodedInstruction[1] = registry[instruction[1]-1]
		decodedInstruction[2] = registry[instruction[2]-1]
	default:
	}
	return decodedInstruction
}

func execute(memory []byte, decodedInstruction [3]byte, registry [2]byte, registerAddress byte) [2]byte {
	command := decodedInstruction[0]
	switch command {
	case LOAD:
		registry[decodedInstruction[1]-1] = decodedInstruction[2]
	case STORE:
		memory[decodedInstruction[2]] = decodedInstruction[1]
	case ADD:
		registry[registerAddress-1] = decodedInstruction[1] + decodedInstruction[2]
	case SBT:
		registry[registerAddress-1] = decodedInstruction[1] - decodedInstruction[2]
	}
	return registry
}
