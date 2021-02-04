package utils

import "errors"

type Pair struct {
	a interface{}
	b interface{}
}

func NewPair(a interface{}, b interface{}) (Pair, error) {
	if a == nil || b == nil {
		return Pair{}, errors.New("both elements of a pair must be present")
	}

	return Pair{a, b}, nil
}

func (p *Pair) A() interface{} {
	return p.a
}

func (p *Pair) B() interface{} {
	return p.b
}

func (p *Pair) ErrorB() error {
	b, done := p.b.(error)
	if done {
		return b
	}
	return nil
}
