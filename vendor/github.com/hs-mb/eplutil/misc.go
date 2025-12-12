package eplutil

import "fmt"

func (b *EPLBuilder) Label() {
	b.WriteString("N")
}

func (b *EPLBuilder) Print(n int) {
	b.WriteString(fmt.Sprintf("P%d", n))
}

func (b *EPLBuilder) Density(t int) {
	b.WriteString(fmt.Sprintf("D%d", t))
}

func (b *EPLBuilder) Speed(v int) {
	b.WriteString(fmt.Sprintf("S%d", v))
}
