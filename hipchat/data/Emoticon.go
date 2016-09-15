package data

const (
	EmoticonMaxLen = 15
)

type Emoticon struct {
	Name string
}

func NewEmoticon() *Emoticon { // {{{
	return getEmoticon()
} // }}}

func (this *Emoticon) String() string { // {{{
	return this.Name
} // }}}

func (this *Emoticon) Reset() { // {{{
	this.Name = ""
} // }}}

func (this *Emoticon) Release() { // {{{
	putEmoticon(this)
} // }}}
