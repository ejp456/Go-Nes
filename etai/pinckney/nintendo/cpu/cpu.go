package cpu

type CPU struct {
	//program counter
	programCounter uint16
	stackPointer uint16
	accumulator uint16
	// Index Register
	registerX uint16
	// Index Register
	registerY uint16
	processorFlags uint16

	// Memory here
}

func run()  {
	//cpu := CPU{}
	//for !quit() {
		//currentOpcode := readMemory(cpu)
		//runOpcode(currentOpcode)
		//checkTickCount()
	//}
}

func runOpcode(currentOpcode uint16)  {

}

func readMemory(cpu *CPU) uint16 {
				return 9
}