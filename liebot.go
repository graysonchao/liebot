package liebot

// Search returns an array of direct links to each comic URL returned by the given search query
func Search(query string) string {

	titles, err := searchTitles(query)
	if err == nil && len(titles) > 0 {
		return titles[0]
	}

	dialogue, err := searchDialogue(query)
	if err == nil && len(dialogue) > 0 {
		return dialogue[0]
	}

	return "I couldn't find anything."
}
