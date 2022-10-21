package pkg

type Factory struct {
	Collector Collector
}

func NewFactory(collector Collector) *Factory {
	return &Factory{Collector: collector}
}

func (fact *Factory) SetCollector(collector Collector) {
	fact.Collector = collector
}

func (fact *Factory) CreateComputer() Computer {
	fact.Collector.SetCore()
	fact.Collector.SetMemory()
	fact.Collector.SetBrand()
	fact.Collector.SetGraphicCard()
	fact.Collector.SetMonitor()

	return fact.Collector.GetComputer()
}
