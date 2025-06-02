#
# Simple build and lint environment
#

#######################################################################################################################

PROJECT := goidgen

#######################################################################################################################

.SUFFIXES	:
.SUFFIXES	:	.go

#
# print colored output
RESET_COLOR    := \033[0m
make_std_color := \033[3$1m      # defined for 1 through 7
make_color     := \033[38;5;$1m  # defined for 1 through 255
OK_COLOR       := $(strip $(call make_std_color,2))
WRN_COLOR      := $(strip $(call make_std_color,3))
ERR_COLOR      := $(strip $(call make_std_color,1))
STD_COLOR      := $(strip $(call make_color,8))

COLOR_OUTPUT = 2>&1 |                                        \
    while IFS='' read -r line; do                            \
        if  [[ $$line == --- FAIL:* ]]; then                 \
            echo -e "$(ERR_COLOR)$${line}$(RESETCOLOR)";     \
        elif [[ $$line == ?* ]]; then                				 \
            echo -e "$(WARN_COLOR)$${line}$(RESET_COLOR)";   \
        elif [[ $$line == ok* ]]; then                			 \
            echo -e "$(OK_COLOR)$${line}$(RESET_COLOR)";     \
        else                                                 \
            echo -e "$(STD_COLOR)$${line}$(RESET_COLOR)";    \
        fi;                                                  \
    done; exit $${PIPESTATUS[0]};

.DEFAULT: $(help)

#BUILD_NUM     := $(shell expr `cat .buildnum 2>/dev/null` + 1 >.buildnum && cat .buildnum)
BUILD_NUM     := $(shell cat .buildnum)
BUILD_HASH    := $(shell git rev-parse --short HEAD)
BUILD_DATE    := $(shell date +'%Y-%m-%d.%H:%M:%S')
BUILD_HOST    := $(shell hostname)
BUILD_TAG     := $(BUILD_HASH).$(BUILD_NUM)
BUILD_INFO    := $(BUILD_HASH).$(BUILD_NUM)-($(BUILD_HOST))-($(BUILD_DATE))

SHELL           := /bin/bash
#GO              := /usr/local/go/bin/go
GO              := go
#GO_PREFIX       := GOCACHE=off
GRC				      := grc
GO_PREFIX				:=
GO_POSTFIX			:= |egrep 'ok|fail'
GO_FLAGS        := -v
BIN_FILE        := $(PROJECT)
MAIN_FILE       := $(BIN_FILE).go
TEST_FILES      := $(wildcard *_test.go)
TEST_FILES      += $(wildcard test/*.go)
SRCS            := $(filter-out $(wildcard *_test.go), $(wildcard *.go))
SRCS_TEST       := ./...
BUILD_DIR       := ./build
BUILD_DIR_DEV   := $(BUILD_DIR)/dev
RUNTIME_DEV     := ./runtime/develop

#GO_FLAGS			   := -v gcflags=\"-m\"
## eval when target build is called with actual BUILD_INFO setting
GO_LDFLAGS      = -ldflags="-X github.com/raibru/$(PROJECT)/cmd.buildInfo=$(BUILD_INFO)"

CLEAN_FILES 	  :=                    \
									tags                \
									$(wildcard ./tmp/*) \
									./$(BUILD_DIR_DEV)/*

help:
	-@echo "Makefile with following options (make <option>):"
	-@echo "	clean"
	-@echo "	clean-all"
	-@echo "	tdd"
	-@echo "	test"
	-@echo "	test-cover"
	-@echo "	test-cache"
	-@echo "	test-verbose"
	-@echo "	test-coverage"
	-@echo "	ctags"
	-@echo "	build (curent os)"
	-@echo "	build-windows"
	-@echo "	build-linux"
	-@echo "	build-deploy"
	-@echo "	run"
	-@echo "    (*) not implemented"
	-@echo ""

print:
	-@echo "SRCS       ==> [$(SRCS)]"
	-@echo "SRCS_TEST  ==> [$(SRCS_TEST)]"
	-@echo "TEST_FILES ==> [$(TEST_FILES)]"

.PHONY: all deploy-all clean tdd test test-cover test-cache test-verbose test-trace-view test-coverage
all: test build
deploy-all: clean test build build-windows build-linux deploy-dev

clean:
	$(GO) clean
	rm -f $(CLEAN_FILES)
	rm -r $(BUILD_DIR_DEV)

clean-all:
	$(GO) clean
	rm -r $(BUILD_DIR)

tdd:
	@$(GRC) $(GO_PREFIX) $(GO) test ./... 

test: $(TEST_FILES) $(SRCS)
	@$(GO_PREFIX) $(GO) test ./...

test-cover: $(TEST_FILES) $(SRCS)
	@$(GO_PREFIX) $(GO) test -cover ./...

test-cache: $(TEST_FILES) $(SRCS)
	@$(GO_PREFIX) $(GO) test $(SRCS_TEST)

test-verbose: $(TEST_FILES) $(SRCS)
	@$(GO_PREFIX) $(GO) test -v $(SRCS_TEST)

test-trace-view: $(TEST_FILES) $(SRCS)
	@$(GO_PREFIX) $(GO) test -trace ./tmp/trace.out
	@$(GO) tool trace ./tmp/trace.out

test-coverage: $(TEST_FILES) $(SRCS)
	@$(GO_PREFIX) $(GO) test -coverprofile=./tmp/coverage.out ./...

run:
	$(GO) run $(GO_FLAGS) $(MAIN_FILE)

build:
	$(eval BUILD_NUM := $(shell expr `cat .buildnum 2>/dev/null` + 1 >.buildnum && cat .buildnum))
	$(eval BUILD_INFO := $(BUILD_HASH).$(BUILD_NUM)-($(BUILD_HOST))-($(BUILD_DATE)))
	$(GO) build $(GO_FLAGS) $(GO_LDFLAGS) -o $(BUILD_DIR_DEV)/$(BIN_FILE) $(MAIN_FILE)

build-windows:
	GOOS=windows GOARCH=amd64 $(GO) build $(GO_FLAGS) $(GO_LDFLAGS) -o $(BUILD_DIR)/windows/amd64/$(BIN_FILE).exe $(MAIN_FILE)

build-linux:
	GOOS=linux GOARCH=amd64 $(GO) build $(GO_FLAGS) $(GO_LDFLAGS) -o $(BUILD_DIR)/linux/amd64/$(BIN_FILE) $(MAIN_FILE)

build-deploy: build build-windows build-linux
	-@echo "Build deployment versions..."

deploy-dev:
	cp $(BUILD_DIR_DEV)/$(BIN_FILE) $(RUNTIME_DEV)/$(BIN_FILE)

ctags:
	ctags -RV .

# EOF