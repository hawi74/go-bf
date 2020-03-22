package bf

type Operator uint8

const (
	opModifyValue Operator = iota
	opModifyPointer
	opOut
	opIn
	opJumpForward
	opJumpBackward
)

type Instruction struct {
	Operator
	Operand int
}

type Program struct {
	instructions []Instruction
	stack        []int
	pointer      int
	opPointer    int
}

func (p *Program) Compile(code string) {
	jumpStack := make([]int, 1)
	for i, c := range code {
		var instruction Instruction
		switch c {
		case '+':
			instruction = Instruction{
				Operator: opModifyValue,
				Operand:  1,
			}
		case '-':
			instruction = Instruction{
				Operator: opModifyValue,
				Operand:  -1,
			}
		case '<':
			instruction = Instruction{
				Operator: opModifyPointer,
				Operand:  -1,
			}
		case '>':
			instruction = Instruction{
				Operator: opModifyPointer,
				Operand:  1,
			}
		case '[':
			instruction = Instruction{
				Operator: opJumpForward,
			}
			jumpStack = append(jumpStack, i)
		case ']':
			jumpTo := jumpStack[len(jumpStack)-1]
			jumpStack = jumpStack[:len(jumpStack)-1]
			instruction = Instruction{
				Operator: opJumpBackward,
				Operand:  jumpTo,
			}
			p.instructions[jumpTo].Operand = i
		case '.':
			instruction = Instruction{
				Operator: opOut,
			}
		}
		p.instructions = append(p.instructions, instruction)
	}
}

func (p *Program) Run() string {
	output := ""
	for p.opPointer < len(p.instructions) {
		opPtr := p.opPointer
		p.opPointer++
		instruction := p.instructions[opPtr]
		switch instruction.Operator {
		case opModifyValue:
			p.stack[p.pointer] += instruction.Operand
		case opModifyPointer:
			p.pointer += instruction.Operand
		case opJumpForward:
			if p.stack[p.pointer] == 0 {
				p.opPointer = instruction.Operand
			}
		case opJumpBackward:
			if p.stack[p.pointer] != 0 {
				p.opPointer = instruction.Operand
			}
		case opOut:
			output += string(p.stack[p.pointer])
		}
	}
	return output
}
