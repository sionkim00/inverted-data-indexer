package utils

type Index map[string][]int

func (i Index) Add(docs []document) {
	for _, doc := range docs {
		for _, token := range analyze(doc.Text) {
			ids := i[token]
		}
	}
}
