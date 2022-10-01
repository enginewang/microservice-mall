module github.com/enginewang/microservice-mall/paymentApi

go 1.15

require (
	github.com/enginewang/microservice-mall/common v0.0.0-20201220124206-86137b1ec049
	github.com/enginewang/microservice-mall/payment v0.0.0-20201221121831-df2fbfbc61ed
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5 // indirect
	github.com/golang/protobuf v1.4.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/select/roundrobin/v2 v2.9.1 // indirect
	github.com/plutov/paypal/v3 v3.1.0
)
