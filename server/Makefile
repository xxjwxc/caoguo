.PHONY: all
all:
	make clear
	gen
	run

.PHONY: build
build: $(NAME)
$(NAME):
	go mod tidy
	CGO_ENABLED=$(CGO) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags $(LDFLAGS) \
	-o $(NAME)$(BINARY_EXT);

.PHONY: run
run:
	go mod tidy
	CGO_ENABLED=$(CGO) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags $(LDFLAGS) \
    -o $(NAME)$(BINARY_EXT);
	./$(NAME)$(BINARY_EXT) debug

.PHONY: clear
clear:
	test ! -d internal/model/ || rm -rf  internal/model/*
	test ! -d err/ || rm -rf err/
	test ! -f $(NAME) || rm $(NAME)
	test ! -f $(NAME).exe || rm $(NAME).exe

.PHONY: gen
gen:
	#test -h rpc || ln -s ../rpc ./
	go generate
	make gormt
	gormt$(BINARY_EXT) -o internal/model

.PHONY: master
master:
	go get -u github.com/xxjwxc/public@master
	go get -u github.com/xxjwxc/ginrpc@master
	go get -u github.com/gmsec/micro@master
	go get -u github.com/gmsec/goplugins@master

.PHONY: install_gormt
gormt:
	go get -u -v github.com/xxjwxc/gormt@master


################################################################################
# Variables																       #
################################################################################
LDFLAGS := "-s -w"
CGO	    ?= 0
NAME 	:= caoguo
RPCFILE := rpc

LOCAL_ARCH := $(shell uname -m)
ifeq ($(LOCAL_ARCH),x86_64)
	TARGET_ARCH_LOCAL = amd64
else ifeq ($(shell echo $(LOCAL_ARCH) | head -c 5),armv8)
	TARGET_ARCH_LOCAL = arm64
else ifeq ($(shell echo $(LOCAL_ARCH) | head -c 5),aarch64)
	TARGET_ARCH_LOCAL = arm64
else ifeq ($(shell echo $(LOCAL_ARCH) | head -c 4),armv)
	TARGET_ARCH_LOCAL = arm
else
	TARGET_ARCH_LOCAL = amd64
endif
export GOARCH ?= $(TARGET_ARCH_LOCAL)

LOCAL_OS := $(shell uname)
ifeq ($(LOCAL_OS),Linux)
	TARGET_OS_LOCAL = linux
else ifeq ($(LOCAL_OS),Darwin)
	TARGET_OS_LOCAL = darwin
else
	TARGET_OS_LOCAL ?= windows
	BINARY_EXT_LOCAL = .exe
endif
export GOOS ?= $(TARGET_OS_LOCAL)
export BINARY_EXT ?= $(BINARY_EXT_LOCAL)