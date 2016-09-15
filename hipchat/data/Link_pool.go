package data

import "sync"

var linkPool = sync.Pool{ // {{{
	New: func() interface{} {
		return &Link{}
	},
} // }}}

func getLink() *Link { // {{{
	if instance := linkPool.Get(); instance != nil {
		return instance.(*Link)
	}

	return linkPool.New().(*Link)
} // }}}

func putLink(instance *Link) { // {{{
	instance.Reset()
	linkPool.Put(instance)
} // }}}
