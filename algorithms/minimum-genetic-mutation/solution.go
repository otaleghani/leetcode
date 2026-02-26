package main

func minMutation(startGene string, endGene string, bank []string) int {
	queue := []string{startGene}
	level := 0

	visited := make(map[string]bool)
	visited[startGene] = true

	for len(queue) > 0 {
		size := len(queue)

		for k := 0; k < size; k++ {
			curr := queue[0]
			queue = queue[1:]

			if curr == endGene {
				return level
			}

			for _, gene := range bank {
				// Skip if we've already used this gene
				if visited[gene] {
					continue
				}

				diff := 0
				for j := 0; j < 8; j++ {
					if curr[j] != gene[j] {
						diff++
					}
				}

				if diff == 1 {
					visited[gene] = true
					queue = append(queue, gene)
				}
			}
		}
		level++
	}

	return -1
}
