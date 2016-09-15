PKG=github.com/ateleshev/go-ss-test
GOPATH:=$(PWD)/bin:$(GOPATH)
export GOPATH

all: test

$(PWD):
	  mkdir $(PWD)/bin

root: $(PWD)

clean:
	  rm -rf $(PWD)/bin

build:
	  git submodule update --init --recursive --remote && echo "Vendors was successfully updated"
	  go build -i -o bin/hipchat $(PKG)/cmd/hipchat

generate: root build

test: generate root
		go test $(PKG)/hipchat/message
		go test $(PKG)/hipchat/loader
		go test -benchmem -benchtime=3s -bench . $(PKG)/hipchat/message
		go test -benchmem -benchtime=3s -bench . $(PKG)/hipchat/loader

.PHONY: root clean generate test build
