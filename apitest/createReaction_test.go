package apitest

import (
	"testing"

	"github.com/romshark/dgraph_graphql_go/apitest/setup"
	"github.com/romshark/dgraph_graphql_go/store/enum/emotion"
)

// TestCreateReaction tests reaction creation
func TestCreateReaction(t *testing.T) {
	ts := setup.New(t, tcx)
	defer ts.Teardown()

	debug := ts.Debug()

	// User 1
	firstP := debug.Help.OK.CreateUser("first", "1@test.test", "testpass")

	// User 2
	secondP := debug.Help.OK.CreateUser("second", "2@test.test", "testpass")
	second, _ := ts.Client("2@test.test", "testpass")

	// User 3
	thirdP := debug.Help.OK.CreateUser("third", "3@test.test", "testpass")
	third, _ := ts.Client("3@test.test", "testpass")

	// Post
	post := debug.Help.OK.CreatePost(*firstP.ID, "test post", "test content")

	// Reaction -> Post
	reaction1 := second.Help.OK.CreateReaction(
		*secondP.ID,
		*post.ID,
		emotion.Happy,
		"nice!",
	)

	// Reaction -> Reaction
	third.Help.OK.CreateReaction(
		*thirdP.ID,
		*reaction1.ID,
		emotion.Happy,
		"me too!",
	)
}
