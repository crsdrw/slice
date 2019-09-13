package slice

type SmallFloat64 struct {
	local [12]float64
	data  []float64
}

func (sm SmallFloat64) Data() []float64 {
	if sm.data == nil {
		sm.data = sm.local[:0]
	}
	return sm.data
}

func (sm SmallFloat64) Append(ss ...[]float64) SmallFloat64 {
	return SmallFloat64{data: append(sm.data, ss...)}
}
