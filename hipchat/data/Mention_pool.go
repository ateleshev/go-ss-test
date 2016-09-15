package data

import "sync"

var mentionPool = sync.Pool{ // {{{
	New: func() interface{} {
		return &Mention{}
	},
} // }}}

func getMention() *Mention { // {{{
	if instance := mentionPool.Get(); instance != nil {
		return instance.(*Mention)
	}

	return mentionPool.New().(*Mention)
} // }}}

func putMention(instance *Mention) { // {{{
	instance.Reset()
	mentionPool.Put(instance)
} // }}}
