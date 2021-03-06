package opcode

import "Go-Nes/etai/pinckney/nintendo/cpu"

type OpCode struct {
	code uint16
	name string
}

func (o OpCode) Name() string {
	return o.name
}


func OpCodeTable() [256]OpCode {
	var opCodes = [256]OpCode{
		OpCode{name: "BRK"}, OpCode{name: "ORA"}, OpCode{name:"KIL"},
		OpCode{name: "SLO"}, OpCode{name:"NOP"}, OpCode{name:"ORA"},
		OpCode{name:"ASL"}, OpCode{name:"SLO"}, OpCode{name:"PHP"},
		OpCode{name:"ORA"}, OpCode{name:"ASL"}, OpCode{name:"ANC"},
		OpCode{name:"NOP"}, OpCode{name:"ORA"}, OpCode{name:"ASL"},
		OpCode{name:"SLO"}, OpCode{name:"BPL"}, OpCode{name:"ORA"},
		OpCode{name:"KIL"}, OpCode{name:"SLO"}, OpCode{name:"NOP"},
		OpCode{name:"ORA"}, OpCode{name:"ASL"}, OpCode{name:"SLO"},
		OpCode{name:"CLC"}, OpCode{name:"ORA"}, OpCode{name:"NOP"},
		OpCode{name:"SLO"}, OpCode{name:"NOP"}, OpCode{name:"ORA"},
		OpCode{name:"ASL"}, OpCode{name:"SLO"}, OpCode{name:"JSR"},
		OpCode{name:"AND"}, OpCode{name:"KIL"}, OpCode{name:"RLA"},
		OpCode{name:"BIT"}, OpCode{name:"AND"}, OpCode{name:"ROL"},
		OpCode{name:"RLA"}, OpCode{name:"PLP"}, OpCode{name:"AND"},
		OpCode{name:"ROL"}, OpCode{name:"ANC"}, OpCode{name:"BIT"},
		OpCode{name:"AND"}, OpCode{name:"ROL"}, OpCode{name:"RLA"},
		OpCode{name:"BMI"}, OpCode{name:"AND"}, OpCode{name:"KIL"},
		OpCode{name:"RLA"}, OpCode{name:"NOP"}, OpCode{name:"AND"},
		OpCode{name:"ROL"}, OpCode{name:"RLA"}, OpCode{name:"SEC"},
		OpCode{name:"AND"}, OpCode{name:"NOP"}, OpCode{name:"RLA"},
		OpCode{name:"NOP"}, OpCode{name:"AND"}, OpCode{name:"ROL"},
		OpCode{name:"RLA"}, OpCode{name:"RTI"}, OpCode{name:"EOR"},
		OpCode{name:"KIL"}, OpCode{name:"SRE"}, OpCode{name:"NOP"},
		OpCode{name:"EOR"}, OpCode{name:"LSR"}, OpCode{name:"SRE"},
		OpCode{name:"PHA"}, OpCode{name:"EOR"}, OpCode{name:"LSR"},
		OpCode{name:"ALR"}, OpCode{name:"JMP"}, OpCode{name:"EOR"},
		OpCode{name:"LSR"}, OpCode{name:"SRE"}, OpCode{name:"BVC"},
		OpCode{name:"EOR"}, OpCode{name:"KIL"}, OpCode{name:"SRE"},
		OpCode{name:"NOP"}, OpCode{name:"EOR"}, OpCode{name:"LSR"},
		OpCode{name:"SRE"}, OpCode{name:"CLI"}, OpCode{name:"EOR"},
		OpCode{name:"NOP"}, OpCode{name:"SRE"}, OpCode{name:"NOP"},
		OpCode{name:"EOR"}, OpCode{name:"LSR"}, OpCode{name:"SRE"},
		OpCode{name:"RTS"}, OpCode{name:"ADC"}, OpCode{name:"KIL"},
		OpCode{name:"RRA"}, OpCode{name:"NOP"}, OpCode{name:"ADC"},
		OpCode{name:"ROR"}, OpCode{name:"RRA"}, OpCode{name:"PLA"},
		OpCode{name:"ADC"}, OpCode{name:"ROR"}, OpCode{name:"ARR"},
		OpCode{name:"JMP"}, OpCode{name:"ADC"}, OpCode{name:"ROR"},
		OpCode{name:"RRA"}, OpCode{name:"BVS"}, OpCode{name:"ADC"},
		OpCode{name:"KIL"}, OpCode{name:"RRA"}, OpCode{name:"NOP"},
		OpCode{name:"ADC"}, OpCode{name:"ROR"}, OpCode{name:"RRA"},
		OpCode{name:"SEI"}, OpCode{name:"ADC"}, OpCode{name:"NOP"},
		OpCode{name:"RRA"}, OpCode{name:"NOP"}, OpCode{name:"ADC"},
		OpCode{name:"ROR"}, OpCode{name:"RRA"}, OpCode{name:"NOP"},
		OpCode{name:"STA"}, OpCode{name:"NOP"}, OpCode{name:"SAX"},
		OpCode{name:"STY"}, OpCode{name:"STA"}, OpCode{name:"STX"},
		OpCode{name:"SAX"}, OpCode{name:"DEY"}, OpCode{name:"NOP"},
		OpCode{name:"TXA"}, OpCode{name:"XAA"}, OpCode{name:"STY"},
		OpCode{name:"STA"}, OpCode{name:"STX"}, OpCode{name:"SAX"},
		OpCode{name:"BCC"}, OpCode{name:"STA"}, OpCode{name:"KIL"},
		OpCode{name:"AHX"}, OpCode{name:"STY"}, OpCode{name:"STA"},
		OpCode{name:"STX"}, OpCode{name:"SAX"}, OpCode{name:"TYA"},
		OpCode{name:"STA"}, OpCode{name:"TXS"}, OpCode{name:"TAS"},
		OpCode{name:"SHY"}, OpCode{name:"STA"}, OpCode{name:"SHX"},
		OpCode{name:"AHX"}, OpCode{name:"LDY"}, OpCode{name:"LDA"},
		OpCode{name:"LDX"}, OpCode{name:"LAX"}, OpCode{name:"LDY"},
		OpCode{name:"LDA"}, OpCode{name:"LDX"}, OpCode{name:"LAX"},
		OpCode{name:"TAY"}, OpCode{name:"LDA"}, OpCode{name:"TAX"},
		OpCode{name:"LAX"}, OpCode{name:"LDY"}, OpCode{name:"LDA"},
		OpCode{name:"LDX"}, OpCode{name:"LAX"}, OpCode{name:"BCS"},
		OpCode{name:"LDA"}, OpCode{name:"KIL"}, OpCode{name:"LAX"},
		OpCode{name:"LDY"}, OpCode{name:"LDA"}, OpCode{name:"LDX"},
		OpCode{name:"LAX"}, OpCode{name:"CLV"}, OpCode{name:"LDA"},
		OpCode{name:"TSX"}, OpCode{name:"LAS"}, OpCode{name:"LDY"},
		OpCode{name:"LDA"},OpCode{name:"LDX"},OpCode{name:"LAX"},
		OpCode{name:"CPY"},OpCode{name:"CMP"},OpCode{name:"NOP"},
		OpCode{name:"DCP"},OpCode{name:"CPY"},OpCode{name:"CMP"},
		OpCode{name:"DEC"},OpCode{name:"DCP"},OpCode{name:"INY"},
		OpCode{name:"CMP"},OpCode{name:"DEX"},OpCode{name:"AXS"},
		OpCode{name:"CPY"},OpCode{name:"CMP"},OpCode{name:"DEC"},
		OpCode{name:"DCP"},OpCode{name:"BNE"},OpCode{name:"CMP"},
		OpCode{name:"KIL"},OpCode{name:"DCP"},OpCode{name:"NOP"},
		OpCode{name:"CMP"},OpCode{name:"DEC"},OpCode{name:"DCP"},
		OpCode{name:"CLD"},OpCode{name:"CMP"},OpCode{name:"NOP"},
		OpCode{name:"DCP"},OpCode{name:"NOP"},OpCode{name:"CMP"},
		OpCode{name:"DEC"},OpCode{name:"DCP"},OpCode{name:"CPX"},
		OpCode{name:"SBC"},OpCode{name:"NOP"},OpCode{name:"ISC"},
		OpCode{name:"CPX"},OpCode{name:"SBC"},OpCode{name:"INC"},
		OpCode{name:"ISC"},OpCode{name:"INX"},OpCode{name:"SBC"},
		OpCode{name:"NOP"},OpCode{name:"SBC"},OpCode{name:"CPX"},
		OpCode{name:"SBC"},OpCode{name:"INC"},OpCode{name:"ISC"},
		OpCode{name:"BEQ"},OpCode{name:"SBC"},OpCode{name:"KIL"},
		OpCode{name:"ISC"},OpCode{name:"NOP"},OpCode{name:"SBC"},
		OpCode{name:"INC"},OpCode{name:"ISC"},OpCode{name:"SED"},
		OpCode{name:"SBC"},OpCode{name:"NOP"},OpCode{name:"ISC"},
		OpCode{name:"NOP"},OpCode{name:"SBC"},OpCode{name:"INC"},
		OpCode{name:"ISC"},
	}
	return opCodes
}

func (o OpCode) adc(cpu *cpu.CPU) {

}