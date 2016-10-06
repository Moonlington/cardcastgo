package cardcastgo

var (
	EndpointCardcast = "https://api.cardcastgame.com/v1/"
	EndpointDecks    = EndpointCardcast + "decks/"

	EndpointDeck      = func(dID string) string { return EndpointDecks + dID }
	EndpointCalls     = func(dID string) string { return EndpointDecks + dID + "/calls" }
	EndpointResponses = func(dID string) string { return EndpointDecks + dID + "/responses" }
	EndpointCall      = func(dID string, cID string) string { return EndpointDecks + dID + "/calls/" + cID }
	EndpointResponse  = func(dID string, cID string) string { return EndpointDecks + dID + "/responses/" + cID }
)
