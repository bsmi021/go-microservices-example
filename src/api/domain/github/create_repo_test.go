package github

import (
	"encoding/json"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJson (t *testing.T){
	request := CreateRepoRequest {
		Name: "test-repo",
		Description: "this is a test",
		Homepage: "google.com",
		Private: true,
		HasIssues: true,
		HasProjects: true,
		HasWiki: true,
	}

	// Marshal takes an input interface and attempts to create a valid json string
	// works like json.dumps (Python) but returns a byte array
	bytes, err := json.Marshal(request)
	
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	var target CreateRepoRequest
	
	// Unmarshal takes an input byte array and a *pointer* that we're tyring to fill using this json
	// works like json.loads (Python) but requires a pointer to the variable which will receive the data
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)

	assert.EqualValues(t, target.Name, request.Name)
	assert.EqualValues(t, target.HasIssues, request.HasIssues)

}