all: mydyn

mydyn: deps
	go build -o mydyn ./src

deps:
	go get github.com/divan/gorilla-xmlrpc/xml
	touch deps

clean: go-clean
	rm -f deps

go-clean:
	rm -f mydyn
