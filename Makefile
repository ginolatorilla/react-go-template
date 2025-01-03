# Copyright © 2024 Gino Latorilla
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in
# all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
# THE SOFTWARE.
APP           = react-go-template
VERSION       = 0.1.0
GITHUB_OWNER  = ginolatorilla
GITHUB_DOMAIN = github.com

COMMIT_HASH = $(shell git rev-parse HEAD)
PACKAGE     = $(GITHUB_DOMAIN)/$(GITHUB_OWNER)/$(APP)

BUILD_FLAGS = -v -buildvcs
LD_FLAGS    = -ldflags="-X '$(PACKAGE)/server.AppName=$(APP)' -X '$(PACKAGE)/server.Version=$(VERSION)' -X '$(PACKAGE)/server.CommitHash=$(COMMIT_HASH)'"
TEST_REGEX  = ".*"
TEST_PACKAGE = "./..."

.PHONY: all
all: test build/ui build

.PHONY: test
test: tidy
	@echo "🌡  Running tests..."
	go test -race $(BUILD_FLAGS) $(LD_FLAGS) -run $(TEST_REGEX) $(TEST_PACKAGE)

.PHONY: test/cover
test/cover: tidy
	@echo "🌡️  Running tests..."
	@go test -coverprofile=/tmp/coverage.out -race $(BUILD_FLAGS) $(LD_FLAGS) -run $(TEST_REGEX) $(TEST_PACKAGE)
	@go tool cover -html=/tmp/coverage.out

.PHONY: tidy
tidy:
	@echo "🧹 Tidying up package dependencies..."
	go mod tidy

.PHONY: build
build:
	@echo "🏗️  Building the server..."
	go build $(BUILD_FLAGS) $(LD_FLAGS) -o bin/$(APP) $(PACKAGE) 

.PHONY: build/ui
build/ui:
	@echo "🏗️  Building the UI..."
	npm -C ui run build

.PHONY: clean
clean:
	go clean
	rm -rf bin/* ui/dist/*

.PHONY: doc
doc:
	@go install golang.org/x/pkgsite/cmd/pkgsite@latest
	@pkgsite -open

.PHONY: install
install: all
	go install $(BUILD_FLAGS) $(LD_FLAGS) $(PACKAGE)

.PHONY: help
help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  help       - Show this help message"
	@echo "  all        - Run test, tidy, and build (default)"
	@echo "  install    - Install the application"
	@echo "  test       - Run tests"
	@echo "  test/cover - Run tests with coverage"
	@echo "  tidy       - Sort out package dependencies"
	@echo "  build      - Build the server"
	@echo "  build/ui   - Build the UI"
	@echo "  clean      - Clean up the build artifacts"
	@echo "  doc        - Open the documentation in the browser"
