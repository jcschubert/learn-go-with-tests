package maps

// Search returns the map[word] from the given map
func Search(dict map[string]string, word string) string {
	return dict[word]
}
