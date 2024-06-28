package weight

type WeightNode struct {
	Addr            string
	Weight          int64
	EffectiveWeight int64
	CurrentWeight   int64
}

type WeightRoundLoadBalance struct {
	list []*WeightNode
}
