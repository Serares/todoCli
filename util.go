package todo

// mutate the existing list in memory
// TODO is this the best solution?
func filter(tl *List) {
	filteredList := make(List, 0)

	for _, s := range *tl {
		if !s.Done {
			filteredList = append(filteredList, s)
		}
	}
	*tl = filteredList
}
