package loader

import "errors"
import "net/http"
import "github.com/ateleshev/go-bin/html/dom"
import "github.com/ateleshev/go-ss-test/hipchat/data"

var (
	ErrResponseStatusIsNotOK = errors.New("Response status is not OK")
)

func UpdateLinkTitle(link *data.Link) error {
	resp, err := http.Get(link.Url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ErrResponseStatusIsNotOK
	}

	finder := dom.NewElmFinder(resp.Body)
	defer finder.Release()
	finder.SetMaxLoadSize(100 * 1024) // 100 kB

	title, err := finder.Find("html.head.title")
	if err != nil {
		return err
	}

	link.Title = string(title)
	return nil
}
