LINUX_OS = linux
WIN_OS = windows
DARWIN_OS = darwin
ARCH = amd64

DIST_DIR = _dist
LINUX_DIR = $(DIST_DIR)/$(LINUX_OS)
WIN_DIR = $(DIST_DIR)/$(WIN_OS)
DARWIN_DIR = $(DIST_DIR)/$(DARWIN_OS)

BIN_FILE = litepub
OTHER_FILES = LICENSE README.md

VERSION = $(shell git describe --tags --abbrev=0)

build: clean
	@go build

install: clean
	@go install

dist: clean
	@echo "Building Linux distribution"
	@mkdir -p $(LINUX_DIR)
	@GOOS=$(LINUX_OS) GOARCH=$(ARCH) go build -o $(LINUX_DIR)/$(BIN_FILE)
	@zip -qj9 $(DIST_DIR)/$(BIN_FILE)-$(VERSION)-$(LINUX_OS)-$(ARCH).zip $(LINUX_DIR)/$(BIN_FILE) $(OTHER_FILES)
	@rm -rf $(LINUX_DIR)

	@echo "Building Windows distribution"
	@mkdir -p $(WIN_DIR)
	@GOOS=$(WIN_OS) GOARCH=$(ARCH) go build -o $(WIN_DIR)/$(BIN_FILE).exe
	@zip -qj9 $(DIST_DIR)/$(BIN_FILE)-$(VERSION)-$(WIN_OS)-$(ARCH).zip $(WIN_DIR)/$(BIN_FILE).exe $(OTHER_FILES)
	@rm -rf $(WIN_DIR)

	@echo "Building Darwin distribution"
	@mkdir -p $(DARWIN_DIR)
	@GOOS=$(DARWIN_OS) GOARCH=$(ARCH) go build -o $(DARWIN_DIR)/$(BIN_FILE)
	@zip -qj9 $(DIST_DIR)/$(BIN_FILE)-$(VERSION)-$(DARWIN_OS)-$(ARCH).zip $(DARWIN_DIR)/$(BIN_FILE) $(OTHER_FILES)
	@rm -rf $(DARWIN_DIR)

clean:
	@go clean
	@rm -rf $(DIST_DIR)

.PHONY: build install dist clean
