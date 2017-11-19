package overwatch

import "fmt"

// TagsForGameplay analyzes keywords and returns tags based upon them
func TagsForGameplay(keywords []string) []string {
	tags := make([]string, 0)

	for index, word := range keywords {
		switch word {
		case "Solo", "Single", "Double", "Quadruple", "Quintuple", "Sextuple":
			tags = append(tags, fmt.Sprintf("%s %s", word, keywords[index+1]))

		case "Triple", "3x":
			tags = append(tags, fmt.Sprintf("%s %s", word, keywords[index+1]), "Three's a Crowd")

		case "1x", "2x", "4x", "5x", "6x":
			tags = append(tags, fmt.Sprintf("%s %s", word, keywords[index+1]))

		case "Potg":
			tags = append(tags, "POTG")

		case "On", "The", "Go", "No":
			tags = append(tags, fmt.Sprintf("%s %s", word, keywords[index+1]))

		case "Boop":
			tags = append(tags, "Boop", "Satisfying", "Have a Nice Trip", "See You Next Fall")

		case "Bomb":
			tags = append(tags, "Go Boom")

		case "Backfills", "Backfill":
			tags = append(tags, "Better Late Than Never")
		}
	}

	if !contains(tags, "POTG") && !contains(tags, "Highlight") {
		tags = append(tags, "Highlight")
	}

	return tags
}
