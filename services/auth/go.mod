module github.com/vladdoroniuk/rose/services/auth

go 1.21.5

require (
	github.com/caarlos0/env/v10 v10.0.0
	github.com/vladdoroniuk/rose/proto_gen v0.0.0-20231230150022-4174934940d4
	google.golang.org/grpc v1.60.1
)

replace github.com/vladdoroniuk/rose/proto_gen/auth => ../../proto_gen

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.16.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231002182017-d307bd883b97 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
)
