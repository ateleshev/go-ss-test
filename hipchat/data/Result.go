package data

import "github.com/ateleshev/go-bin/encoding/json"

type Result struct {
	Mentions  []*Mention
	Emoticons []*Emoticon
	Links     []*Link
}

func NewResult() *Result { // {{{
	return getResult()
} // }}}

func (this *Result) Reset() { // {{{
	this.Mentions = this.Mentions[:0]
	this.Emoticons = this.Emoticons[:0]
	this.Links = this.Links[:0]
} // }}}

func (this *Result) Release() { // {{{
	for _, mention := range this.Mentions {
		mention.Release()
	}

	for _, emoticon := range this.Emoticons {
		emoticon.Release()
	}

	for _, link := range this.Links {
		link.Release()
	}

	putResult(this)
} // }}}

func (this *Result) JsonWriteTo(jw *json.JsonWriter) *json.JsonWriter { // {{{
	var n int
	jw.ObjOpen()

	// -- Mentions --

	if len(this.Mentions) > 0 {
		jw.StringValue("mentions").Sep().ArrOpen()
		for i, v := range this.Mentions {
			if i > 0 {
				jw.Next()
			}
			jw.StringValue(v.Name)
		}
		jw.ArrClose()
		n++
	}

	// -- Emoticons --

	if len(this.Emoticons) > 0 {
		jw.NextIf(n > 0)
		jw.StringValue("emoticons").Sep().ArrOpen()
		for i, v := range this.Emoticons {
			if i > 0 {
				jw.Next()
			}
			jw.StringValue(v.Name)
		}
		jw.ArrClose()
		n++
	}

	// -- Links --

	if len(this.Links) > 0 {
		jw.NextIf(n > 0)

		jw.StringValue("links").Sep().ArrOpen()
		for i, v := range this.Links {
			if i > 0 {
				jw.Next()
			}
			v.JsonWriteTo(jw)
		}
		jw.ArrClose()
		n++
	}

	return jw.ObjClose()
} // }}}
