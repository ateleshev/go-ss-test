package data

import "sync"

var emoticonPool = sync.Pool{ // {{{
	New: func() interface{} {
		return &Emoticon{}
	},
} // }}}

func getEmoticon() *Emoticon { // {{{
	if instance := emoticonPool.Get(); instance != nil {
		return instance.(*Emoticon)
	}

	return emoticonPool.New().(*Emoticon)
} // }}}

func putEmoticon(instance *Emoticon) { // {{{
	instance.Reset()
	emoticonPool.Put(instance)
} // }}}
