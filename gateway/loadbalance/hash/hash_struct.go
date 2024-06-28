package hash

import "sync"

type ConsistentHashBalance struct {
	mux      sync.RWMutex
	hash     Hash              //hash函数
	replicas int               //复制因子
	keys     Uint32Slice       //已排序的hash节点切片
	hashmap  map[uint32]string //key为hash值val为节点
}
