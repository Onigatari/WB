package main

import (
	"main/pkg"
)

func main() {
	asusCollector := pkg.GetCollector("Asus")
	hpCollector := pkg.GetCollector("HP")
	dellCollector := pkg.GetCollector("Dell")

	factory := pkg.NewFactory(asusCollector)

	asusComputer := factory.CreateComputer()
	asusComputer.Print()

	factory.SetCollector(hpCollector)

	hpComputer := factory.CreateComputer()
	hpComputer.Print()

	factory.SetCollector(dellCollector)

	dellComputer := factory.CreateComputer()
	dellComputer.Print()
}
