// Copyright 2021 MIKAMAI s.r.l
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
