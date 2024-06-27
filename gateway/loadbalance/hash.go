package loadbalance

import (
	"errors"
	"hash/crc32"
	"log"
	"sort"
	"strconv"
	"sync"
)

type Hash func([]byte) uint32

type Uint32Slice []uint32

func (s Uint32Slice) Len() int {
	return len(s)
}
func (s Uint32Slice) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Uint32Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ConsistentHashBalance struct {
	mux      sync.RWMutex
	hash     Hash              //hash函数
	replicas int               //复制因子
	keys     Uint32Slice       //已排序的hash节点切片
	hashmap  map[uint32]string //key为hash值val为节点
}

func NewConsistenceHashBalance(replicas int, hash Hash) *ConsistentHashBalance {

	log.Println("--------------进入hash算法负载均衡---------------------")
	c := &ConsistentHashBalance{
		hash:     hash,
		replicas: replicas,
		hashmap:  make(map[uint32]string),
		keys:     make([]uint32, 0, 100),
	}

	if c.hash == nil {
		//保证是一个2^32 - 1的一个环
		c.hash = crc32.ChecksumIEEE
	}
	return c
}

func (c *ConsistentHashBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("param need more than one")
	}
	addr := params[0]
	//计算虚拟节点hash值
	for i := 0; i < c.replicas; i++ {
		hash := c.hash([]byte(strconv.Itoa(i) + addr))
		//实现了排序接口
		c.keys = append(c.keys, hash)
		c.hashmap[hash] = addr
	}
	sort.Sort(c.keys)
	return nil
}

// 得到取缓存的服务器
func (c *ConsistentHashBalance) Next(key string) (string, error) {
	if len(c.keys) == 0 {
		return "", nil
	}
	hash := c.hash([]byte(key))
	//通过二分查询到最优节点（第一个hash大于资源hash的服务器）
	idx := sort.Search(len(c.keys), func(i int) bool {
		return c.keys[i] > hash
	})
	if idx == len(c.keys) {
		//没有找到服务器，说明此时处于环的尾部，那么直接用第0台服务器
		idx = 0
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	val := c.hashmap[c.keys[idx]]
	return val, nil
}
