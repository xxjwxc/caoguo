NAME := caoguo
RPCFILE := rpc 
all: # 构建
	make clear
build:
	make clear
	./tools/gormt -o internal/model
	go build -o $(NAME) main.go 
	./$(NAME) debug
run:
	# make clear
	# ./tools/gormt -o internal/model
	go build -o $(NAME) main.go 
	./$(NAME) debug
windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w"  -o $(NAME).exe main.go 
mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w"  -o $(NAME) main.go 
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w"  -o $(NAME) main.go 
clear: 
	test ! -d internal/model/ || rm -rf  internal/model/*
	test ! -d err/ || rm -rf err/
	test ! -f $(NAME) || rm $(NAME)
	test ! -f $(NAME).exe || rm $(NAME).exe
gen:
	#test -h rpc || ln -s ../rpc ./
	go generate
	./tools/gormt.exe -o internal/model
master:
	go get -u github.com/xxjwxc/public@master
	go get -u github.com/xxjwxc/ginrpc@master
	go get -u github.com/gmsec/micro@master
	go get -u github.com/gmsec/goplugins@master