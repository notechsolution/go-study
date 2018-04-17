package shaper

type Shaper interface {
	Area() float32;
}

type Square struct {
	Size float32;
}

func (square Square) Area() float32 {
	return square.Size * square.Size;
}

type Retangle struct {
	Length, Width float32;
}

func (retangle *Retangle) Area() float32 {
	return retangle.Length * retangle.Width;
}

