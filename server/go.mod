module caoguo

go 1.14

require (
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/gin-gonic/gin v1.7.2
	github.com/gmsec/goplugins v0.0.0-20200512180521-4f1b493507af
	github.com/gmsec/micro v0.0.0-20210824093724-29d61bf5ebe3
	github.com/golang/protobuf v1.5.2
	github.com/jinzhu/gorm v1.9.12
	github.com/xxjwxc/ginrpc v0.0.0-20200526160328-764c776e63c0
	github.com/xxjwxc/public v0.0.0-20210929024509-c93144cad67c
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	gorm.io/gorm v1.20.2
)

replace github.com/xxjwxc/ginrpc => ../ginrpc

replace github.com/xxjwxc/public => ../public

replace github.com/gmsec/goplugins => ../goplugins
