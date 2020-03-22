package bf

import (
	"testing"
)

var helloWorld = "++++++++8[>++++14[>++>+++>+++>+<<<<-33]>+>+>->>+43[<45]<-48]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."

func TestCompileAndRun(t *testing.T) {
	p := Program{stack: make([]int, 100)}
	p.Compile(helloWorld)
	expected := "Hello World!\n"
	if output := p.Run(); output != expected {
		t.Errorf("Expected %v, got: %v", expected, output)
	}
}
