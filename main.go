package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type Heap struct {
	data []*ListNode
}

// Len heap is 1 indexed
func (h *Heap) Len() int {
	return len(h.data) - 1
}

// Append
// BubbleUp
func (h *Heap) Insert(x *ListNode) {
	(*h).data = append(h.data, x)
	h.BubbleUp(h.Len())
}

// working with index here
// bubble up the last
func (h *Heap) BubbleUp(k int) {
	p, ok := parent(k)
	if !ok {
		return
	}

	if h.data[p].Val > h.data[k].Val {
		h.data[k], h.data[p] = h.data[p], h.data[k]
		h.BubbleUp(p)
	}
}

func (h *Heap) PopMin() (*ListNode, bool) {
	if h.Len() == 0 {
		return nil, false
	}
	min := h.data[1]
	h.data[1] = h.data[h.Len()] // TODO
	(*h).data = h.data[:h.Len()]
	h.BubbleDown(1)
	return min, true
}

func (h *Heap) BubbleDown(index int) {
	min := index
	leftOfIndex := left(index)

	for i := 0; i < 2; i++ {
		if (leftOfIndex + i) <= h.Len() {
			if h.data[min].Val > h.data[leftOfIndex+i].Val {
				min = leftOfIndex + i
			}
		}
	}

	if min != index {
		h.data[index], h.data[min] = h.data[min], h.data[index]
		h.BubbleDown(min)
	}
}

func left(k int) int {
	return 2 * k
}

func right(k int) int {
	return 2*k + 1
}

func parent(k int) (int, bool) {
	if k == 1 {
		return 0, false
	}
	return k / 2, true
}

func mergeKLists(lists []*ListNode) *ListNode {
	ln := make([]*ListNode, 1)
	heap := Heap{ln}
	var ret *ListNode
	var cur *ListNode

	if len(lists) == 0 {
		return nil
	}

	// insert everything into the heap
	for _, node := range lists {
		if node != nil {
			heap.Insert(node)
		}
	}

	for heap.Len() > 0 {
		node, _ := heap.PopMin()
		fmt.Println(node)

		// if there is a node following the current node
		// put it back in the heap so we can prioritize it
		if node.Next != nil {
			heap.Insert(node.Next)
		}

		// only executed once
		// we return the first node returned from the heap
		// and we set the current for the following loop
		if ret == nil {
			ret = node
			cur = node
		} else {
			// set the next node returned
			// and set the current
			cur.Next = node
			cur = node
		}
	}
	return ret
}
