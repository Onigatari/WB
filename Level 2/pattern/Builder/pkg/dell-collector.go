package pkg

type DellCollector struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

func (collector *DellCollector) SetCore() {
	collector.Core = 16
}

func (collector *DellCollector) SetBrand() {
	collector.Brand = "Dell"
}

func (collector *DellCollector) SetMemory() {
	collector.Memory = 32
}

func (collector *DellCollector) SetMonitor() {
	collector.Monitor = 3
}

func (collector *DellCollector) SetGraphicCard() {
	collector.GraphicCard = 2
}

func (collector *DellCollector) GetComputer() Computer {
	return Computer{
		Core:        collector.Core,
		Brand:       collector.Brand,
		Memory:      collector.Memory,
		Monitor:     collector.Monitor,
		GraphicCard: collector.GraphicCard,
	}
}
