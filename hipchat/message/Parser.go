package message

import "errors"
import "unicode/utf8"
import "github.com/ateleshev/go-ss-test/hipchat/data"
import "github.com/ateleshev/go-bin/net/http/url"

const (
	MentionInitByte   = '@'
	EmoticonStartByte = '('
	EmoticonStopByte  = ')'
	UrlStartByteLow   = 'h'
	UrlStartByteUpp   = 'H'
)

var (
	ErrValueIsEmpty  = errors.New("Value is empty")
	ErrIsNotMention  = errors.New("Is not mention")
	ErrIsNotEmoticon = errors.New("Is not emoticon")
	ErrIsNotUrl      = errors.New("Is not url")
)

type Parser interface {
	Do([]byte) (*data.Result, error)
}

func NewParser() Parser {
	return &parser{}
}

type parser struct {
}

func (this *parser) Do(v []byte) (*data.Result, error) { // {{{
	if len(v) == 0 {
		return nil, ErrValueIsEmpty
	}

	r := data.NewResult()
	for i := 0; i < len(v); i++ {
		switch v[i] {
		case utf8.RuneSelf:
			_, size := utf8.DecodeRune(v[i:])
			i += size
			continue
		case MentionInitByte:
			if n, err := this.fetchMention(v[i+1:]); err == nil {
				mention := data.NewMention()
				mention.Name = string(v[i+1 : i+n+1])
				r.Mentions = append(r.Mentions, mention)
				i += n + 1
			}
			break
		case EmoticonStartByte:
			if n, err := this.fetchEmoticon(v[i+1:]); err == nil {
				emoticon := data.NewEmoticon()
				emoticon.Name = string(v[i+1 : i+n+1])
				r.Emoticons = append(r.Emoticons, emoticon)
				i += n + 2
			}
			break
		case UrlStartByteLow, UrlStartByteUpp:
			if n, secure, err := this.fetchUrl(v[i:]); err == nil {
				link := data.NewLink()
				link.Url = string(v[i : i+n])
				link.Secure = secure
				r.Links = append(r.Links, link)
				i += n
			}
			break
		}
	}

	return r, nil
} // }}}

func (this *parser) fetchMention(v []byte) (int, error) { // {{{
	n := 0
	for i := n; i < len(v); i++ {
		switch {
		case IsAllowedForMention(v[i]):
			n++
			continue
		default:
			goto end_mention
		}
	}

end_mention:
	if n == 0 {
		return 0, ErrIsNotMention
	}

	return n, nil
} // }}}

func (this *parser) fetchEmoticon(v []byte) (int, error) { // {{{
	n, stop := 0, false
	for i := n; i < len(v); i++ {
		switch {
		case IsAllowedForEmoticon(v[i]) && n <= data.EmoticonMaxLen:
			n++
			continue
		case v[i] == EmoticonStopByte:
			if n <= data.EmoticonMaxLen {
				stop = true
			}
			goto end_emoticon
		default:
			goto end_emoticon
		}
	}

end_emoticon:
	if n == 0 || stop == false {
		return 0, ErrIsNotEmoticon
	}

	return n, nil
} // }}}

func (this *parser) fetchUrl(v []byte) (int, bool, error) { // {{{
	if len(v) <= url.MinLen {
		return 0, false, ErrIsNotUrl
	}

	var i, n int
	var secure bool
	var err error

	if i, secure, err = url.SchemeIndex(v); err != nil {
		return 0, false, ErrIsNotUrl
	}
	n += i

	if n >= len(v) {
		return 0, false, ErrIsNotUrl
	}

	if i, err = url.AuthorityIndex(v[n:]); err != nil {
		return 0, false, ErrIsNotUrl
	}
	n += i

	if n >= len(v) {
		goto end_url
	}

	if i, err = url.PathIndex(v[n:]); err != nil {
		goto end_url
	}
	n += i

	if n >= len(v) {
		goto end_url
	}

	if i, err = url.QueryIndex(v[n:]); err != nil {
		goto end_url
	}
	n += i

	if n >= len(v) {
		goto end_url
	}

	if i, err = url.FragmentIndex(v[n:]); err != nil {
		goto end_url
	}
	n += i

end_url:
	return n, secure, nil
} // }}}
