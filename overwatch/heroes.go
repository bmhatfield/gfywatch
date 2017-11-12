package overwatch

// TagsForHero accepts title-cased keywords and returns an array of appropriate tags
func TagsForHero(keywords []string) []string {
	tags := make([]string, 0)

	for _, word := range keywords {
		switch word {
		// Offense
		case "Doomfist", "Genji", "Reaper", "Sombra":
			tags = append(tags, word, "DPS")

		case "McCree", "Mccree":
			tags = append(tags, "McCree", "DPS")

		case "Soldier: 76", "Soldier":
			tags = append(tags, "Soldier: 76", "Soldier", "DPS")

		case "Tracer":
			tags = append(tags, word, "DPS", "Annoying", "Gotta Go Fast")

		case "Pharah":
			tags = append(tags, word, "DPS", "Attac", "Protec", "Respec", "Justice Rains From Above")

		// Defense
		case "Bastion", "Hanzo", "Junkrat", "Mei":
			tags = append(tags, word, "Defense")

		case "Torbjörn", "Torb", "Torbjorn":
			tags = append(tags, "Torbjörn", "Torb", "Defense")

		case "Widowmaker", "Widow":
			tags = append(tags, "Widowmaker", "Defense")

		// Tank
		case "D.Va", "Dva":
			tags = append(tags, "D.Va", "Tank")

		case "Orisa", "Winston", "Zarya":
			tags = append(tags, word, "Tank")

		case "Reinhardt", "Rein":
			tags = append(tags, "Reinhardt", "Rein", "Tank")

		case "Roadhog", "Hog":
			tags = append(tags, "Roadhog", "Hog", "Tank")

		// Support
		case "Ana":
			tags = append(tags, word, "Support", "Granny")

		case "Lucio", "Lúcio":
			tags = append(tags, "Lucio", "Lúcio", "Support")

		case "Mercy":
			tags = append(tags, word, "Support", "Angel", "Valkryie")

		case "Moira":
			tags = append(tags, word, "Support")

		case "Symmetra", "Symm":
			tags = append(tags, word, "Support", "Turrets", "Microwave")

		case "Zenyatta", "Zen":
			tags = append(tags, word, "Support", "Orbs")
		}
	}

	return tags
}
