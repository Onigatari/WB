package point

import "math"

func absInt(n int) int {
	return int(math.Abs(float64(n)))
}

// Point Доступты только внутри пакета
type Point struct {
	x int
	y int
}

// CreatePoint Конструктор Point
func CreatePoint(x, y int) Point {
	return Point{x, y}
}

// GetDistance Растояние между точками
func GetDistance(firstPoint, secondPoint Point) float64 {
	dx := absInt(secondPoint.x - firstPoint.x)
	dy := absInt(secondPoint.y - firstPoint.y)

	dist := math.Sqrt(float64(dx*dx + dy*dy))
	return dist
}
