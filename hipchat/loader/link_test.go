package loader_test

import "testing"
import "github.com/ateleshev/go-ss-test/hipchat/data"
import "github.com/ateleshev/go-ss-test/hipchat/loader"

/**
 * ==[ Tests ]==
 *
 * go test -v -run=UpdateLinkTitle
 */

func TestUpdateLinkTitle(t *testing.T) {
	link := data.NewLink()
	link.Url = "https://golang.org"
	if err := loader.UpdateLinkTitle(link); err != nil {
		t.Fatal(err)
	}

	if link.Title != "The Go Programming Language" {
		t.Fatal("Incorrect link title, waiting 'The Go Programming Language', received:", link.Title)
	}
}

/**
 * ==[ Benchmarks ]==
 * go test -v -run=^$ -benchmem -bench=UpdateLinkTitle
 */

func BenchmarkUpdateLinkTitle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		link := data.NewLink()
		link.Url = "https://golang.org"

		if err := loader.UpdateLinkTitle(link); err != nil {
			b.Fatal(err)
		}

		if link.Title != "The Go Programming Language" {
			b.Fatal("Incorrect link title, waiting 'The Go Programming Language', received:", link.Title)
		}
	}
}
