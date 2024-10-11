package tracer

type Interval struct {
	Min, Max float64
}

func NewInterval(min, max float64) Interval {
	return Interval{Min: min, Max: max}
}

func (i Interval) Contains(x float64) bool {
	return i.Min <= x && x <= i.Max
}

func (i Interval) Surrounds(x float64) bool {
	return i.Min < x && x < i.Max
}

func (i Interval) Clamp(x float64) float64 {
	if x < i.Min {
		return i.Min
	} else if x > i.Max {
		return i.Max
	}
	return x
}

func (i Interval) Expand(δ float64) Interval {
	padding := δ / 2
	return Interval{Min: i.Min - padding, Max: i.Max + padding}
}
