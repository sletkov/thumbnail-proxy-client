.PHONY: build

build:
	 go build -o thumbnail-proxy-client

install: build
	go build -o thumbnail-proxy-client-tmp
	mv ./thumbnail-proxy-client-tmp /usr/local/bin/thumbnail-proxy-client
	chmod +x /usr/local/bin/thumbnail-proxy-client
	cp thumbnail-proxy-client.yml ${HOME}/thumbnail-proxy-client.yml
