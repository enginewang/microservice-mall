module github.com/enginewang/microservice-mall/cart

go 1.19

require (
	github.com/enginewang/microservice-mall/common v0.0.0-20201210084148-1ab1eced812e
	github.com/golang/protobuf v1.4.0
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/opentracing/opentracing-go v1.2.0
)
