package utils

type Index map[string][]int

func (i Index) Add(docs []document) {
	for _, doc := range docs {
		for _, token := range analyze(doc.Text) {
			ids := i[token]

			if ids != nil && i[len(ids)-1] == doc.ID {
				// Ignore duplicate ID
				continue
			}
			i[token] = append(ids, doc.ID)
		}
	}
}

func (i Index) Search(text string) []int {
	var r []int

	for _, token := range analyze(text) {
		if ids, ok := i[token]; ok {
			if r == nil {
				r = ids
			} else {
				r = Intersection(r, ids)
			}
		} else {
			// Token does not exist
			return nil
		}
	}

	return r
}

func Intersection(a []int, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}

	r := make([]int, 0, maxLen)

	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}

	return r
}
