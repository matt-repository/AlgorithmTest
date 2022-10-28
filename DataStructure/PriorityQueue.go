package DataStructure

import "container/heap"

type Item struct {
	word  string
	count int
}

type ItemHeap []Item

func (h ItemHeap) Len() int { return len(h) }

func (h ItemHeap) Less(i, j int) bool {
	if h[i].count != h[j].count {
		return h[i].count < h[j].count
	} else {
		return h[i].word > h[j].word
	}
}

func (h ItemHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *ItemHeap) Push(val interface{}) {
	*h = append(*h, val.(Item))
}

func (h *ItemHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(words []string, k int) []string {
	countMap := make(map[string]int)
	for _, word := range words {
		countMap[word]++
	}

	h := &ItemHeap{}
	for w, v := range countMap {
		heap.Push(h, Item{
			word:  w,
			count: v,
		})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]string, h.Len())
	for i := h.Len() - 1; i >= 0; i-- {
		item := heap.Pop(h).(Item)
		res[i] = item.word
	}
	return res
}
