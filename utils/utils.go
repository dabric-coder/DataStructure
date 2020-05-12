package utils


type Interface interface {
	Len() int
	Swap(i, j int)
	Comparator(i, j int) bool
}

type IntSlice []int
type StringSlice []string

func (p IntSlice) Len() int {
	return len(p)
}

func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p IntSlice) Comparator(i, j int) bool {
	return p[i] > p[j]
}


func (s StringSlice) Len() int {
	return len(s)
}

func (s StringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s StringSlice) Comparator(i, j int) bool {
	return s[i] > s[j]
}
