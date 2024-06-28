package weight

import (
	"errors"
	"strconv"
)

func (r *WeightRoundLoadBalance) Add(params ...string) error {
	if len(params) != 2 {
		return errors.New("params need more 2")
	}
	parseInt, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		return errors.New("convert to int error")
	}
	newNode := &WeightNode{
		Addr:            params[0],
		Weight:          parseInt,
		EffectiveWeight: parseInt,
		CurrentWeight:   parseInt,
	}
	r.list = append(r.list, newNode)
	return nil
}
func (r *WeightRoundLoadBalance) Next() string {
	var total int64 = 0
	var best *WeightNode
	for index, v := range r.list {
		total += v.CurrentWeight
		r.list[index].CurrentWeight += r.list[index].EffectiveWeight
		if v.EffectiveWeight < v.CurrentWeight {
			v.EffectiveWeight++
		}
		if best == nil || best.CurrentWeight < v.CurrentWeight {
			best = v
		}
	}
	if best == nil {
		return ""
	}
	best.CurrentWeight -= total
	return best.Addr
}

func (r *WeightRoundLoadBalance) Get() (string, error) {
	return r.Next(), nil
}
