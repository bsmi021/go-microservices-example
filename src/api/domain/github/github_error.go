package github

/*
{
	"message": "Repository creation failed",
	"errors": [
		{
			"resource": "Repository",
			"code": "custom",
			"field": "name",
			"message": "name already exists on this account"
		}
	],
	"documentation_url": "https://developer.github.com/v3/repos/#create"
}
*/

// ErrorResponse represents an error returned from the Github API
type ErrorResponse struct {
	StatusCode		 int		 `json:"status_code"` // added for clarity
	Message          string      `json:"message"`
	Errors           []ErrorItem `json:"errors"`
	DocumentationURL string      `json:"documentation_url"`
}

func (r ErrorResponse) Error() string {
	return r.Message
}

// ErrorItem represents an error item returned by a GithubError
type ErrorItem struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}
