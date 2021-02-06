package utils

type Pair struct {
	a interface{}
	b interface{}
}

func NewPair(a interface{}, b interface{}) Pair {
	return Pair{a, b}
}

func (p *Pair) A() interface{} {
	return p.a
}

func (p *Pair) B() interface{} {
	return p.b
}

func (p *Pair) StringA() string {
	s, done := p.a.(string)
	if done {
		return s
	}
	return ""
}

func (p *Pair) ErrorB() error {
	b, done := p.b.(error)
	if done {
		return b
	}
	return nil
}
