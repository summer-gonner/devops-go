package random

import (
	"errors"
	"math/rand"
)

func (r *RandomBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("params need more than 0")
	}
	addr := params[0]
	r.res = append(r.res, addr)
	return nil
}
func (r *RandomBalance) Next() string {
	if len(r.res) == 0 {
		return ""
	}
	r.curIndex = rand.Intn(len(r.res))
	return r.res[r.curIndex]
}

func (r *RandomBalance) Get() (string, error) {
	return r.Next(), nil
}
