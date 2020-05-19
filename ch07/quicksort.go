package ch07

func quicksort(a []int, p int, r int) {
	if p < r {
		q := partition(a, p, r)
		quicksort(a, p, q-1)
		quicksort(a, q+1, r)
	}
}

func partition(a []int, p int, r int) int {
	x := a[r]
	i := p - 1
	for j := p; j < r; j++ {
		if a[j] < x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}
