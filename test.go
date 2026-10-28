func moveCheapest(a, b *Stack, ops *[]string) {

	bestCost := 1 << 30
	bestA, bestB := 0, 0

	for i, val := range *a {
		tb := targetB(val, *b)
		j := indexOf(*b, tb)

		cost := abs(costToTop(i, len(*a))) + abs(costToTop(j, len(*b)))

		if cost < bestCost {
			bestCost = cost
			bestA = i
			bestB = j
		}
	}

	rotateBoth(a, b, bestA, bestB, ops)

	pb(a, b, ops)
}

func targetA(bVal int, a Stack) int {
	target := -1
	best := 1 << 30

	for _, v := range a {
		if v > bVal && v < best {
			best = v
			target = v
		}
	}

	if target == -1 {
		// prendre min A
		target = a[0]
		for _, v := range a {
			if v < target {
				target = v
			}
		}
	}
	return target
}

func targetB(aVal int, b Stack) int {
	target := -1
	best := -1

	for _, v := range b {
		if v < aVal && v > best {
			best = v
			target = v
		}
	}

	if target == -1 {
		// prendre max B
		target = b[0]
		for _, v := range b {
			if v > target {
				target = v
			}
		}
	}
	return target
}

func costToTop(i, size int) int {
	if i <= size/2 {
		return i
	}
	return i - size
}

func isSorted(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i-1] > a[i] {
			return false
		}
	}
	return true
}

func minIndex(a []int) int {
	min := 0
	for i := 1; i < len(a); i++ {
		if a[i] < a[min] {
			min = i
		}
	}
	return min
}

func indexOf(a []int, val int) int {
	for i, v := range a {
		if v == val {
			return i
		}
	}
	return -1
}