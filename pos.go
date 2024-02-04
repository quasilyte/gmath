package gmath

// Pos represents a position with optional offset relative to its base.
type Pos struct {
	Base   *Vec
	Offset Vec
}

func MakePos(base Vec) Pos {
	return Pos{Base: &base}
}

func (p Pos) Resolve() Vec {
	if p.Base == nil {
		return p.Offset
	}
	return p.Base.Add(p.Offset)
}

func (p *Pos) SetBase(base Vec) {
	p.Base = &base
}

func (p *Pos) Set(base *Vec, offsetX, offsetY float64) {
	p.Base = base
	p.Offset.X = offsetX
	p.Offset.Y = offsetY
}

func (p Pos) WithOffset(offsetX, offsetY float64) Pos {
	return Pos{
		Base:   p.Base,
		Offset: Vec{X: p.Offset.X + offsetX, Y: p.Offset.Y + offsetY},
	}
}
