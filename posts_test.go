package thesrc

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/sourcegraph/thesrc/router"
)

func TestPostsService_Get(t *testing.T) {
	setup()
	defer teardown()

	want := &Post{ID: 1}

	var called bool
	mux.HandleFunc(urlPath(t, router.Post, map[string]string{"ID": "1"}), func(w http.ResponseWriter, r *http.Request) {
		called = true
		testMethod(t, r, "GET")

		writeJSON(w, want)
	})

	post, err := client.Posts.Get(1)
	if err != nil {
		t.Errorf("Posts.Get returned error: %v", err)
	}

	if !called {
		t.Fatal("!called")
	}

	normalizeTime(&want.SubmittedAt)
	if !reflect.DeepEqual(post, want) {
		t.Errorf("Posts.Get returned %+v, want %+v", post, want)
	}
}

func TestPostsService_List(t *testing.T) {
	setup()
	defer teardown()

	want := []*Post{{ID: 1}}

	var called bool
	mux.HandleFunc(urlPath(t, router.Posts, nil), func(w http.ResponseWriter, r *http.Request) {
		called = true
		testMethod(t, r, "GET")
		testFormValues(t, r, values{})

		writeJSON(w, want)
	})

	posts, err := client.Posts.List(nil)
	if err != nil {
		t.Errorf("Posts.List returned error: %v", err)
	}

	if !called {
		t.Fatal("!called")
	}

	for _, p := range want {
		normalizeTime(&p.SubmittedAt)
	}
	if !reflect.DeepEqual(posts, want) {
		t.Errorf("Posts.List returned %+v, want %+v", posts, want)
	}
}

func TestPostsService_Create(t *testing.T) {
	setup()
	defer teardown()

	want := &Post{Title: "t"}

	var called bool
	mux.HandleFunc(urlPath(t, router.CreatePost, nil), func(w http.ResponseWriter, r *http.Request) {
		called = true
		testMethod(t, r, "POST")
		testBody(t, r, `{"Title":"t","LinkURL":"","Body":"","SubmittedAt":"0001-01-01T00:00:00Z","AuthorUserID":0}`+"\n")

		writeJSON(w, want)
	})

	post := &Post{Title: "t"}
	err := client.Posts.Create(post)
	if err != nil {
		t.Errorf("Posts.Create returned error: %v", err)
	}

	if !called {
		t.Fatal("!called")
	}

	normalizeTime(&want.SubmittedAt)
	if !reflect.DeepEqual(post, want) {
		t.Errorf("Posts.Create returned %+v, want %+v", post, want)
	}
}
