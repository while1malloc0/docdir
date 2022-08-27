.PHONY: build
build:
	go build -o dist/docdir ./cmd

.PHONY: docs
docs:
	go run github.com/campoy/embedmd ./docsdata/README.tpl.md > README.tmp.md
	sed -i 's:.*embedmd.*::g' README.tmp.md
	mv README.tmp.md README.md
