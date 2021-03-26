package bfsearch_test

import (
	"algorithms/bfsearch"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	search := "Dolly"

	maggie := &bfsearch.Node{
		Name: "Maggie",
		Children: []*bfsearch.Node{&bfsearch.Node{
			Name: search,
		}},
	}
	bart := &bfsearch.Node{
		Name: "Bart",
	}
	liza := &bfsearch.Node{
		Name: "Liza",
	}
	homer := &bfsearch.Node{
		Name:     "Homer",
		Children: []*bfsearch.Node{bart, liza},
	}
	marge := &bfsearch.Node{
		Name:     "Marge",
		Children: []*bfsearch.Node{bart, liza, maggie},
	}
	root := &bfsearch.Node{
		Name:     "Abraham Simpson",
		Children: []*bfsearch.Node{homer, marge}, // Sweet home alabama
	}
	result := bfsearch.Search(root, search)
	assert.Equal(t, search, result.Name)

	result = bfsearch.Search(nil, search)
	assert.Nil(t, result)

	result = bfsearch.Search(root, "fake")
	assert.Nil(t, result)
}
