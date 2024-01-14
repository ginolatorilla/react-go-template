$( shell mkdir -p bin )

PROJ      = react-go-template
ORG_PATH  = github.com/ginolatorilla
REPO_PATH = $(ORG_PATH)/$(PROJ)
VERSION  ?= 0.0.1

export GOBIN=$(PWD)/bin
LD_FLAGS="-w -X main.version=$(VERSION)"

build: bin/$(PROJ)

.PHONY: release
release: LD_FLAGS = "-w -X main.version=$(VERSION) -X main.app=$(PROJ) -extldflags \"-static\""
release: bin/$(PROJ)

bin/$(PROJ): clean
	@mkdir -p bin/
	@go install -v -ldflags $(LD_FLAGS) $(REPO_PATH)

clean:
	@rm -rf bin/