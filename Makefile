GOCMD = go
GOBUILD = $(GOCMD) build
GOMOD = $(GOCMD) mod
GOTEST = $(GOCMD) test
BINARY_NAME = cobweb
CLI = adm

all: serve

serve:
	$(GOCMD) run .

build:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./build/$(BINARY_NAME) -v ./

generate:
	$(CLI) generate -l cn -c adm.ini

test: black-box-test user-acceptance-test

black-box-test: ready-for-data
	$(GOTEST) -v -test.run=TestMainBlackBox
	make clean

user-acceptance-test: ready-for-data
	$(GOTEST) -v -test.run=TestMainUserAcceptance
	make clean

ready-for-data:
	cp admin.db admin_test.db

clean:
	rm admin_test.db

.PHONY: all serve build generate test black-box-test user-acceptance-test ready-for-data clean