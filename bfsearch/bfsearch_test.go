package bfsearch_test

import (
	"algorithms/bfsearch"
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	search := "maggie"
	items := &bfsearch.Item{
		Name: "root",
		Items: &[]bfsearch.Item{
			{
				Name: "homer",
				Items: &[]bfsearch.Item{
					{
						Name: "item_1",
					},
					{
						Name: "item_2",
					},
					{
						Name: "item_3",
					},
				},
			},
			{
				Name: "marge",
				Items: &[]bfsearch.Item{
					{
						Name: "item_4",
					},
					{
						Name: "item_2",
					},
					{
						Name: "item_5",
						Items: &[]bfsearch.Item{
							{
								Name: "item_6",
							},
							{
								Name: "maggie",
							},
						},
					},
				},
			},
		},
	}
	result := bfsearch.Search(items, search)
	fmt.Println(result)
}