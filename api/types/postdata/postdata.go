package postdata

type PostData struct {
	Query     string                 `json:"query"`
	Operation string                 `json:"operation"`
	Variables map[string]interface{} `json:"variables"`
}