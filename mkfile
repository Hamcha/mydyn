mydyn: deps
	go build -o mydyn ./src

deps:
	go get github.com/kolo/xmlrpc
	touch deps

clean:
	rm -f deps mydyn
