package hash

//一致性hash算法
import (
	"errors"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

type Ring struct {
	replicas int
	Nodes    Nodes
	sync.Mutex
}

func NewRing(replicas int) *Ring {
	return &Ring{Nodes: Nodes{}, replicas: replicas}
}

func (r *Ring) AddNodes(ids []string) {
	r.Lock()
	defer r.Unlock()

	r.buildNodeIds(ids, func(realNodeId, virtualNodeId string) {
		node := NewNode(realNodeId, virtualNodeId)
		r.Nodes = append(r.Nodes, node)
	})

	sort.Sort(r.Nodes)
}

func (r *Ring) RemoveNodes(ids []string) error {
	r.Lock()
	defer r.Unlock()

	ok := true
	r.buildNodeIds(ids, func(realNodeId, virtualNodeId string) {
		nodeIndex := r.search(virtualNodeId)
		if nodeIndex >= r.Nodes.Len() || r.Nodes[nodeIndex].Id != realNodeId {
			ok = false
		} else {
			r.Nodes = append(r.Nodes[:nodeIndex], r.Nodes[nodeIndex+1:]...)
		}
	})

	if !ok {
		return NodeNotFoundErr
	}
	return nil
}

func (r *Ring) GetNode(id string) *Node {
	i := r.search(id)
	if i >= r.Nodes.Len() {
		i = 0
	}

	return r.Nodes[i]
}

func (r *Ring) search(id string) int {
	return sort.Search(r.Nodes.Len(), func(i int) bool {
		return r.Nodes[i].hashId >= buildHashId(id)
	})
}

func (r *Ring) buildNodeIds(ids []string, fn func(realNodeId, virtualNodeId string)) {
	for _, id := range ids {
		for i := 0; i < r.replicas; i++ {
			fn(id, strconv.Itoa(i)+id)
		}
	}
}

type Node struct {
	Id     string
	hashId uint32
}

func NewNode(realNodeId, virtualNodeId string) *Node {
	return &Node{
		Id:     realNodeId,
		hashId: buildHashId(virtualNodeId),
	}
}

type Nodes []*Node

func (n Nodes) Len() int           { return len(n) }
func (n Nodes) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n Nodes) Less(i, j int) bool { return n[i].hashId < n[j].hashId }

var NodeNotFoundErr = errors.New("node not found")

func buildHashId(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}
