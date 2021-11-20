package enum

type RestClient string

const (
	// Resty for instance resty
	Resty RestClient = "resty"
)

type ConetentType string

const (
	ContentTypeJson ConetentType = "application/json"
)
