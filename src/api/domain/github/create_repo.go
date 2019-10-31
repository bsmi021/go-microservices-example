package github

// {
// 	"name": "Hello-World",
// 	"description": "This is your first repository",
// 	"homepage": "https://github.com",
// 	"private": false,
// 	"has_issues": true,
// 	"has_projects": true,
// 	"has_wiki": true
//   }

// CreateRepoRequest represents a request body to create a new Github Repository
type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

// CreateRepoResponse represents the response from the Github Repository Create action
type CreateRepoResponse struct {
	ID          int64           `json:"id"`
	Name        string          `json:"name"`
	FullName    string          `json:"full_name"`
	Owner       RepoOwner       `json:"owner"`
	Permissions RepoPermissions `json:"permissions"`
}

// RepoOwner represents the information about the owner of a Github repository
type RepoOwner struct {
	ID      int64  `json:"id"`
	Login   string `json:"login"`
	URL     string `json:"url"`
	HTMLURL string `json:"html_url"`
}

// RepoPermissions represent the permissions for the repo
type RepoPermissions struct {
	IsAdmin bool `json:"admin"`
	HasPull bool `json:"pull"`
	HasPush bool `json:"push"`
}

// func CreateRepo() {
// 	// One way to build up the request, not recommended as a lot of work will have to be done like:
// 	// private := request["private"].(bool) takes more time and doesn't scale well
// 	// request := map[string]interface(){
// 	// 	"name": "Hello-World",
// 	// 	"description": "This is your first repository",
// 	// 	"homepage": "https://github.com",
// 	// 	"private": false,
// 	// 	"has_issues": true,
// 	// 	"has_projects": true,
// 	// 	"has_wiki": true
// 	// }
// }
