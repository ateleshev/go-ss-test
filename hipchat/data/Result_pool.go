package data

import "sync"

var resultPool = sync.Pool{ // {{{
	New: func() interface{} {
		return &Result{}
	},
} // }}}

func getResult() *Result { // {{{
	if instance := resultPool.Get(); instance != nil {
		return instance.(*Result)
	}

	return resultPool.New().(*Result)
} // }}}

func putResult(instance *Result) { // {{{
	instance.Reset()
	resultPool.Put(instance)
} // }}}
