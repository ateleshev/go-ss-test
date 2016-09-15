package message_test

import "testing"
import "github.com/ateleshev/go-ss-test/hipchat/data"
import "github.com/ateleshev/go-ss-test/hipchat/message"

var parserData = [][]byte{
	[]byte("@chris you around?"),
	[]byte("Good morning! (megusta) (coffee)"),
	[]byte("Olympics are starting soon; http://www.nbcolympics.com"),
	[]byte("@bob @john (success) such a cool feature; https://twitter.com/jdorfman/status/430511497475670016"),
}

var parser = message.NewParser()

/**
 * ==[ Tests ]==
 *
 * go test -v -run=Parser_
 */

/**
 * go test -v -run=Parser_Mentions
 */
func TestParser_Mentions(t *testing.T) { // {{{
	var r *data.Result
	var err error

	if r, err = parser.Do(parserData[0]); err != nil {
		t.Fatal(err)
	}

	if len(r.Mentions) != 1 {
		t.Fatal("Incorrect parse result for Mentions, waiting 1 element, received:", len(r.Mentions))
	}

	if r.Mentions[0].Name != "chris" {
		t.Fatal("Incorrect parse result for Mentions, waiting 'chris', received:", r.Mentions[0].Name)
	}
} // }}}

/**
 * go test -v -run=Parser_Emoticons
 */
func TestParser_Emoticons(t *testing.T) { // {{{
	var r *data.Result
	var err error

	if r, err = parser.Do(parserData[1]); err != nil {
		t.Fatal(err)
	}

	if len(r.Emoticons) != 2 {
		t.Fatal("Incorrect parse result for Emoticons, waiting 2 elements, received:", len(r.Emoticons))
	}

	for i := range r.Emoticons {
		switch r.Emoticons[i].Name {
		case "megusta", "coffee":
			continue
		default:
			t.Fatal("Incorrect parse result for Emoticons, waiting 'megusta' and 'coffee', received:", r.Emoticons[i])
			break
		}
	}
} // }}}

/**
 * go test -v -run=Parser_Links
 */
func TestParser_Links(t *testing.T) { // {{{
	var r *data.Result
	var err error

	if r, err = parser.Do(parserData[2]); err != nil {
		t.Fatal(err)
	}

	if len(r.Links) != 1 {
		t.Fatal("Incorrect parse result for Links, waiting 1 element, received:", len(r.Links))
	}

	if r.Links[0].Url != "http://www.nbcolympics.com" {
		t.Fatal("Incorrect parse result for Links, waiting 'http://www.nbcolympics.com', received:", r.Links[0].Url)
	}

	if r.Links[0].Secure {
		t.Fatal("Incorrect parse result for Links, waiting non secure connect (http), received: https")
	}
} // }}}

/**
 * go test -v -run=Parser_Combined
 */
func TestParser_Combined(t *testing.T) { // {{{
	var r *data.Result
	var err error

	if r, err = parser.Do(parserData[3]); err != nil {
		t.Fatal(err)
	}

	if len(r.Mentions) != 2 {
		t.Fatal("Incorrect parse result for Mentions, waiting 2 elements, received:", len(r.Mentions))
	}

	for i := range r.Mentions {
		switch r.Mentions[i].Name {
		case "bob", "john":
			continue
		default:
			t.Fatal("Incorrect parse result for Mentions, waiting 'bob' and 'john', received:", r.Mentions[i].Name)
			break
		}
	}

	if len(r.Emoticons) != 1 {
		t.Fatal("Incorrect parse result for Emoticons, waiting 1 element, received:", len(r.Emoticons))
	}

	if r.Emoticons[0].Name != "success" {
		t.Fatal("Incorrect parse result for Emoticons, waiting 'success', received:", r.Emoticons[0])
	}

	if len(r.Links) != 1 {
		t.Fatal("Incorrect parse result for Links, waiting 1 element, received:", len(r.Links))
	}

	if r.Links[0].Url != "https://twitter.com/jdorfman/status/430511497475670016" {
		t.Fatal("Incorrect parse result for Links, waiting 'https://twitter.com/jdorfman/status/430511497475670016', received:", r.Links[0].Url)
	}

	if !r.Links[0].Secure {
		t.Fatal("Incorrect parse result for Links, waiting secure connect (https), received: http")
	}
} // }}}

/**
 * ==[ Benchmarks ]==
 *
 * go test -v -run=^$ -benchmem -bench=Parser_ -memprofile=mem.1.out -benchtime=3s | tee mem.1.profile
 * go tool pprof --alloc_objects message.test mem.1.out
 * go test -v -run=^$ -benchmem -bench=Parser_ -memprofile=mem.2.out -benchtime=3s | tee mem.2.profile
 * benchcmp mem.1.profile mem.2.profile
 *
 * go test -v -run=^$ -bench=Parser_ -cpuprofile=cpu.1.out -benchtime=3s | tee cpu.1.profile
 * go tool pprof message.test cpu.1.out
 * go test -v -run=^$ -bench=Parser_ -cpuprofile=cpu.2.out -benchtime=3s | tee cpu.2.profile
 * benchcmp cpu.1.profile cpu.2.profile
 *
 * go test -v -run=^$ -benchmem -bench=Parser_
 */

/**
 * go test -v -run=^$ -benchmem -bench=Parser_Mentions
 */
func BenchmarkParser_Mentions(b *testing.B) { // {{{
	for i := 0; i < b.N; i++ {
		if _, err := parser.Do(parserData[0]); err != nil {
			b.Fatal(err)
		}
	}
} // }}}

/**
 * go test -v -run=^$ -benchmem -bench=Parser_Emoticons
 */
func BenchmarkParser_Emoticons(b *testing.B) { // {{{
	for i := 0; i < b.N; i++ {
		if _, err := parser.Do(parserData[1]); err != nil {
			b.Fatal(err)
		}
	}
} // }}}

/**
 * go test -v -run=^$ -benchmem -bench=Parser_Links
 */
func BenchmarkParser_Links(b *testing.B) { // {{{
	for i := 0; i < b.N; i++ {
		if _, err := parser.Do(parserData[2]); err != nil {
			b.Fatal(err)
		}
	}
} // }}}

/**
 * go test -v -run=^$ -benchmem -bench=Parser_Combined
 */
func BenchmarkParser_Combined(b *testing.B) { // {{{
	for i := 0; i < b.N; i++ {
		if _, err := parser.Do(parserData[3]); err != nil {
			b.Fatal(err)
		}
	}
} // }}}

/**
 * go test -v -run=^$ -benchmem -bench=Parser_IterateAll
 */
func BenchmarkParser_IterateAll(b *testing.B) { // {{{
	for i := 0; i < b.N; i++ {
		for j := 0; j < len(parserData); j++ {
			if _, err := parser.Do(parserData[j]); err != nil {
				b.Fatal(err)
			}
		}
	}
} // }}}
