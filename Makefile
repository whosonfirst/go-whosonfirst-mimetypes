CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep rmdeps
	if test -d src; then rm -rf src; fi
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-mimetypes
	cp *.go src/github.com/whosonfirst/go-whosonfirst-mimetypes/
	cp -r lookup src/github.com/whosonfirst/go-whosonfirst-mimetypes/
	if test -d vendor; then cp -r vendor/* src/; fi

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	fmt bin

docker-build:
	docker build -t wof-readwrited .

deps:
	@echo "no dependencies yet"

vendor-deps: rmdeps deps
	if test ! -d vendor; then mkdir vendor; fi
	if test -d vendor; then rm -rf vendor; fi
	cp -r src vendor
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src

fmt:
	go fmt cmd/*.go
	go fmt lookup/*.go
	go fmt *.go

bin: 	self
	@GOPATH=$(GOPATH) go build -o bin/wof-mimetype-lookup cmd/wof-mimetype-lookup.go

lookup-tables:	self
	@GOPATH=$(GOPATH) go build -o bin/build-lookup-tables cmd/build-lookup-tables.go
	if test -d lookup; then rm -rf lookup; fi
	mkdir lookup
	bin/build-lookup-tables -lookup extension > lookup/extension.go
	bin/build-lookup-tables -lookup mimetype > lookup/mimetype.go
	go fmt lookup/*.go
	rm bin/build-lookup-tables
