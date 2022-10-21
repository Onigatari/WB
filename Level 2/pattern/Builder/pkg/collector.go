package pkg

const (
	AsusCollectorType = "Asus"
	HpCollectorType   = "HP"
	DellCollectorType = "Dell"
)

type Collector interface {
	SetCore()
	SetBrand()
	SetMemory()
	SetMonitor()
	SetGraphicCard()
	GetComputer() Computer
}

func GetCollector(collectorType string) Collector {
	switch collectorType {
	default:
		return nil
	case AsusCollectorType:
		return &AsusCollector{}
	case HpCollectorType:
		return &HpCollector{}
	case DellCollectorType:
		return &DellCollector{}
	}

}
