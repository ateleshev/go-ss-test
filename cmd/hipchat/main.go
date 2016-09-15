package main

import (
	"bufio"
	"os"
	"sync"
)

import "github.com/ateleshev/go-bin/encoding/json"

import "github.com/ateleshev/go-ss-test/hipchat/data"
import "github.com/ateleshev/go-ss-test/hipchat/message"
import "github.com/ateleshev/go-ss-test/hipchat/loader"

var wg sync.WaitGroup
var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)
var jw = json.NewJsonWriter(w)

func UpdateLinkTitle(l *data.Link, wg *sync.WaitGroup) {
	defer wg.Done()

	if err := loader.UpdateLinkTitle(l); err != nil {
		w.WriteString("[UpdateLinkTitle] Error: ")
		w.WriteString(err.Error())
		w.WriteByte('\n')
		w.Flush()
	}
}

func main() {
	for {
		// Prompt and read
		w.WriteString("Press [ Ctrl ] + [ C ] to exit. \n")
		w.WriteString("Enter text: ")
		w.Flush()

		buf, _ := r.ReadBytes('\n')
		parser := message.NewParser()

		res, err := parser.Do(buf[:len(buf)-1])
		if err != nil {
			w.WriteString("Error: ")
			w.WriteString(err.Error())
			goto end_iteration
		}

		// TODO: Add timeout
		for _, l := range res.Links {
			wg.Add(1)
			go UpdateLinkTitle(l, &wg)
		}
		wg.Wait()

		w.WriteString("Json: ")
		res.JsonWriteTo(jw)
		if jw.HasErrors() {
			w.WriteString("[JsonWriter] Error: ")
			w.WriteString(jw.LastError().Error())
		}

	end_iteration:

		w.WriteByte('\n')
		w.Flush()
	}
}
