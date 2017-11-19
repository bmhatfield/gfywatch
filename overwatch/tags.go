package overwatch

import "strings"

// TagsFromTitle parses a gfy title and provides a set of tags related to it
func TagsFromTitle(title string) []string {
	keywords := strings.Split(strings.Title(title), " ")
	tags := []string{"Overwatch", "Awesome"}

	heroTags := TagsForHero(keywords)
	tags = append(tags, heroTags...)

	gameplayTags := TagsForGameplay(keywords)
	tags = append(tags, gameplayTags...)

	return tags
}
