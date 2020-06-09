package market

type Bucket struct {
	Width, Height int
	Wheels        []wheel
	number        int
}

func NewWheel() wheel {
	return wheel{}
}

type wheel struct {
	R float64
	s float64
}
