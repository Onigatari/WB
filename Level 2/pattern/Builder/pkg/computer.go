package pkg

import "fmt"

type Computer struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

func (comp *Computer) Print() {
	fmt.Printf("PC:%s Core:[%d] Memory:[%d] GraphicCard:[%d] Monitor:[%d]\n",
		comp.Brand, comp.Core, comp.Memory, comp.GraphicCard, comp.Monitor)
}
