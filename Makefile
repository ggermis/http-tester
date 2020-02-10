# PHONY targets are recipes, not actual files
.PHONY: project-name version clean build install uninstall

PROJECT_NAME  := http-tester
VERSION       := 1.0.0

PLATFORMS     := darwin linux
ARCHITECTURES := 386 amd64

ifdef BUILD_NR
VERSION       := "$(VERSION).$(BUILD_NR)"
endif

LDFLAGS       := -ldflags "-s -w -X github.com/ggermis/http-tester/pkg/http_tester.version=${VERSION}"

all: clean build


project-name:
	@echo $(PROJECT_NAME)

version:
	@echo $(VERSION)


clean:
	@rm -rf dist


build:
	$(foreach GOOS, $(PLATFORMS),\
	  $(foreach GOARCH, $(ARCHITECTURES), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); go build ${LDFLAGS} -o dist/$(PROJECT_NAME)-$(GOOS)-$(GOARCH)-$(VERSION))))

install:
	@go install ${LDFLAGS}

uninstall:
	@go clean -i


