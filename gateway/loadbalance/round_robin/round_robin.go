package round_robin

import "errors"

func (r *RoundRobinBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("params need more than 0")
	}
	addr := params[0]
	r.res = append(r.res, addr)
	return nil
}
func (r *RoundRobinBalance) Next() string {
	if len(r.res) == 0 {
		return ""
	}
	res := r.res[r.curIndex]
	r.curIndex = (r.curIndex + 1) % len(r.res)
	return res
}

func (r *RoundRobinBalance) Get() (string, error) {
	return r.Next(), nil
}
