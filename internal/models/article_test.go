// article_test.go

package models

import "testing"

func TestGetAllArticles(t *testing.T) {
	alist := GetAllArticles()

	// Check that the length of the list of articles returned is the
	// same as the length of the global variable holding the list
	if len(alist) != len(articleList) {
		t.Fail()
	}

	// Check that each member is identical
	for i, v := range alist {
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title {

			t.Fail()
			break
		}
	}
}

func TestGetArticleByID(t *testing.T) {
	alist := GetAllArticles()

	// Check that a valid item was received from the list
	test := alist[0]
	result, err := GetArticleByID(test.ID)

	if err != nil ||
		test.Content != result.Content ||
		test.ID != result.ID ||
		test.Title != result.Title {

		t.Fail()
	}

	// Try to query an invalid ID to make sure the operation fails
	_, err = GetArticleByID(999999)
	if err == nil {
		t.Fail()
	}
}
