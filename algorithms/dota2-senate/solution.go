package main

type Queue []byte

func (q *Queue) Peek() byte {
	if len(*q) == 0 {
		return 0
	}
	return (*q)[0]
}

func (q *Queue) Add(r byte) {
	*q = append(*q, r)
}

func (q *Queue) Pop() byte {
	if len(*q) == 0 {
		return 0
	}
	r := (*q)[0]
	*q = (*q)[1:]
	return r
}

// A not so great solution. Really slow and memory hungry
func predictPartyVictoryCrappy(senate string) string {
	q := Queue{}
	// Add to the queue only if they have
	majority := false
	for !majority {
		for i := 0; i < len(senate); i++ {
			if len(q) == 0 {
				q.Add(senate[i])
				continue
			}
			if senate[i] == q.Peek() {
				q.Add(senate[i])
			} else {
				// If they are different, pop one from the queue and delete it from the senate string
				senate = senate[:i] + senate[i+1:]
				i--
				_ = q.Pop()
			}
		}
		// Checks majority
		candidate := senate[0]
		majority = true
		for i := 0; i < len(senate); i++ {
			if senate[i] != candidate {
				majority = false
			}
		}
	}

	if senate[0] == 'R' {
		return "Radiant"
	}
	return "Dire"
}

// A much better approach using 2 queues
type QueueInt []int

func (q *QueueInt) Add(v int) {
	*q = append(*q, v)
}

func (q *QueueInt) Pop() int {
	val := (*q)[0]
	*q = (*q)[1:]
	return val
}

func (q *QueueInt) IsEmpty() bool {
	return len(*q) == 0
}

func predictPartyVictory(senate string) string {
	radiant := QueueInt{}
	dire := QueueInt{}
	n := len(senate)

	for i, senator := range senate {
		if senator == 'R' {
			radiant.Add(i)
		} else {
			dire.Add(i)
		}
	}

	for !radiant.IsEmpty() && !dire.IsEmpty() {
		radiantIndex := radiant.Pop()
		direIndex := dire.Pop()

		if radiantIndex < direIndex {
			radiant.Add(radiantIndex + n)
		} else {
			dire.Add(direIndex + n)
		}
	}

	if radiant.IsEmpty() {
		return "Dire"
	} else {
		return "Radiant"
	}
}
