package main

type Pair struct {
	pos int
	val int
}

type StockSpanner struct {
	stack []Pair
	count int // Total numbers of calls to next
}

func Constructor() StockSpanner {
	return StockSpanner{
		stack: []Pair{},
	}
}

func (this *StockSpanner) Next(price int) int {
	(*this).count += 1
	if len((*this).stack) == 0 {
		(*this).stack = append((*this).stack, Pair{val: price, pos: this.count})
		return 1
	}

	last := (*this).stack[len((*this).stack)-1]
	if last.val > price {
		(*this).stack = append((*this).stack, Pair{val: price, pos: this.count})
		return 1
	} else {
		for len((*this).stack) > 0 && last.val <= price {
			(*this).stack = (*this).stack[:len((*this).stack)-1]
			if len((*this).stack) == 0 {
				last = Pair{val: 0}
			} else {
				last = (*this).stack[len((*this).stack)-1]
			}
		}
		(*this).stack = append((*this).stack, Pair{val: price, pos: this.count})
		return this.count - last.pos
	}
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
