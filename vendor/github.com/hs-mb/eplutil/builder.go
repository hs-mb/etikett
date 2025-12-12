package eplutil

type EPLBuilder struct {
	Width int
	Height int

	commands [][]byte
	lineBreak string
}

func NewEPLBuilder() *EPLBuilder {
	b := new(EPLBuilder)
	b.Width = 448
	b.Height = 224
	b.lineBreak = "\r\n"
	b.Clear()
	return b
}

func (b *EPLBuilder) Clear() {
	b.commands = make([][]byte, 0)
}

func (b *EPLBuilder) Write(command []byte) (int, error) {
	b.commands = append(b.commands, command)
	return len(command), nil
}

func (b *EPLBuilder) WriteString(command string) (int, error) {
	b.commands = append(b.commands, []byte(command))
	return len(command), nil
}

func (b *EPLBuilder) String() string {
	out := ""
	for _, command := range b.commands {
		out += b.lineBreak + string(command)
	}
	out += b.lineBreak
	return out
}
