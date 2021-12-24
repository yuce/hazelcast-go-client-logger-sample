.PHONY: build clean

build:
	go build .

clean:
	rm -f hazelcast-go-client-logger-sample