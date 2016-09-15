package message

import "github.com/ateleshev/go-bin/bytes/byteutil"

func IsAllowedForMention(v byte) bool { // {{{
	return byteutil.IsWordCharacter(v)
} // }}}

func IsAllowedForEmoticon(v byte) bool { // {{{
	return byteutil.IsAlphanumeric(v)
} // }}}
