module movieexample.com

go 1.16

require (
	github.com/confluentinc/confluent-kafka-go v1.9.1
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang/mock v1.6.0
	github.com/google/go-cmp v0.6.0
	github.com/hashicorp/consul/api v1.12.0
	github.com/m3db/prometheus_client_golang v1.12.8 // indirect
	github.com/stretchr/testify v1.8.4
	github.com/twmb/murmur3 v1.1.6 // indirect
	github.com/uber-go/tally v3.5.0+incompatible
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.46.0
	go.opentelemetry.io/otel v1.20.0
	go.opentelemetry.io/otel/exporters/jaeger v1.9.0
	go.opentelemetry.io/otel/sdk v1.9.0
	go.uber.org/zap v1.23.0
	golang.org/x/time v0.3.0
	google.golang.org/grpc v1.59.0
	google.golang.org/protobuf v1.34.0
	gopkg.in/yaml.v3 v3.0.1
)
