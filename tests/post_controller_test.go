package tests

import (
	"collagen/tests/test_helper"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func TestPostController(t *testing.T) {
	router := test_helper.SetupServerBeforeTest()

	t.Run("Create Post Without Body", func(t *testing.T) {
		request := ""
		requestJSON, _ := json.Marshal(request)

		w := test_helper.RequestApiForTest(router, "POST", "/api/v1/post/", strings.NewReader(string(requestJSON)))

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	//t.Run("Create Valid Post", func(t *testing.T) {
	//
	//	// Prepare a sample form data
	//	post := database_model.Post{
	//		Title:   "Test Post",
	//		Content: "Content",
	//	}
	//
	//	requestJSON, _ := json.Marshal(post)
	//
	//	w := test_helper.RequestApiForTest(router, "POST", "/api/v1/post", strings.NewReader(string(requestJSON)))
	//
	//	// Assertions and validations for the response
	//	assert.Equal(t, http.StatusOK, w.Code)
	//})

}
