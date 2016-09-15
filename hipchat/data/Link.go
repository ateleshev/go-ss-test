package data

import "github.com/ateleshev/go-bin/encoding/json"

type Link struct {
	Url    string
	Secure bool
	Title  string
}

func NewLink() *Link { // {{{
	return getLink()
} // }}}

func (this *Link) String() string { // {{{
	return this.Url
} // }}}

func (this *Link) Reset() { // {{{
	this.Url = ""
	this.Secure = false
	this.Title = ""
} // }}}

func (this *Link) Release() { // {{{
	putLink(this)
} // }}}

func (this *Link) JsonWriteTo(jw *json.JsonWriter) *json.JsonWriter { // {{{
	jw.ObjOpen()
	jw.StringValue("url").Sep().StringValue(this.Url)
	jw.Next().StringValue("title").Sep().StringValue(this.Title)
	return jw.ObjClose()
} // }}}
