package pcgo

// a node in the relationship graph
type RelationshipNode struct {
	Data struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	} `json:"data"`
}
