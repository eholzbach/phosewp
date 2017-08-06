package events

func Help(query string) string {

	var response string

	switch query {
	case "acronym":
		response = "string ; V.E.R.A. -- Virtual Entity of Relevant Acronyms"
	case "astronomy":
		response = "zip code ; Returns the moon phase, sunrise and sunset times"
	case "drama":
		response = "string ; In lulz we trust"
	case "dict":
		response = "string ; Queries WordNet, a large lexical database of English"
	case "fu":
		response = "nil or string ; FoaaS"
	case "quote":
		response = "add string to save ; get [id] to fetch quote"
	case "stock":
		response = "string ; Stock price at previous day closing"
	case "tide":
		response = "zip code ; Tidal information"
	case "trump":
		response = "string ; Tronald Dump"

	case "urban":
		response = "string ; Urban Dictionary"
	case "weather":
		response = "zip code ; Returns the current temperature, weather condition, humidity, wind, 'feels like' temperature, barometric pressure, and visibility"
	case "wiki":
		response = "string ; Wikipedia"
	default:
		response = "Commands are: acronym, astronomy, drama, dict, fu, stock, tide, trump, urban, weather, wiki"
	}

	return response

}
