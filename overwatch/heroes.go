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

		case "Soldier76", "Soldier":
			tags = append(tags, "Soldier: 76", "Soldier", "DPS")

		case "Tracer":
			tags = append(tags, word, "DPS", "Annoying", "Gotta Go Fast", "Here You Go")

		case "Pharah":
			tags = append(tags, word, "DPS", "Attac", "Protec", "Respec", "Justice Rains From Above", "Rocket Lady")

		// Defense
		case "Bastion", "Hanzo", "Junkrat", "Mei":
			tags = append(tags, word, "Defense")

		case "Torbjörn", "Torb", "Torbjorn":
			tags = append(tags, "Torbjörn", "Torb", "Defense", "Turrets", "Hammers")

		case "Widowmaker", "Widow":
			tags = append(tags, "Widowmaker", "Defense")

		// Tank
		case "D.Va", "Dva":
			tags = append(tags, "D.Va", "Tank")

		case "Orisa", "Winston", "Zarya":
			tags = append(tags, word, "Tank")

		case "Reinhardt", "Rein":
			tags = append(tags, "Reinhardt", "Rein", "Tank", "Hammer", "I Live For This")

		case "Roadhog", "Hog":
			tags = append(tags, "Roadhog", "Hog", "Tank", "Hook")

		// Support
		case "Ana":
			tags = append(tags, word, "Support", "Granny", "Healer")

		case "Lucio", "Lúcio":
			tags = append(tags, "Lucio", "Lúcio", "Support", "Turn It Up", "Healer")

		case "Mercy":
			tags = append(tags, word, "Support", "Angel", "Valkryie", "Healer")

		case "Moira":
			tags = append(tags, word, "Support", "Healer")

		case "Symmetra", "Symm":
			tags = append(tags, word, "Support", "Turrets", "Microwave")

		case "Zenyatta", "Zen":
			tags = append(tags, "Zenyatta", "Zen", "Support", "Orbs", "Healer")
		}
	}

	return tags
}
