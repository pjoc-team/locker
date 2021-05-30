module github.com/pjoc-team/locker

go 1.14

require (
	github.com/coreos/bbolt v1.3.5 // indirect
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/btree v1.0.1 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/pjoc-team/tracing v0.0.0-20210526041458-a4e0b42cb369
	github.com/prometheus/client_golang v1.10.0 // indirect
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20201229170055-e5319fda7802 // indirect
	go.etcd.io/etcd v3.3.25+incompatible
	go.uber.org/zap v1.15.0
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	google.golang.org/grpc v1.37.0 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5
