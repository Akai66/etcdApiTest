module etcdApiTest

go 1.14

require (
	github.com/coreos/etcd v3.3.27+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	go.uber.org/zap v1.19.1 // indirect
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa // indirect
	google.golang.org/grpc v1.40.0
)


replace google.golang.org/grpc => google.golang.org/grpc v1.26.0   //grpc v1.26.0以上版本不支持etcd v3，需要replace到v1.26.0版本