package data

type Mention struct {
	Name string
}

func NewMention() *Mention { // {{{
	return getMention()
} // }}}

func (this *Mention) String() string { // {{{
	return this.Name
} // }}}

func (this *Mention) Reset() { // {{{
	this.Name = ""
} // }}}

func (this *Mention) Release() { // {{{
	putMention(this)
} // }}}
