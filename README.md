# go-proto-grpc-demo

A very simple go grpc demo

To generate go code from the proto

```
protoc -I foo/ foo/*.proto --go_out=plugins=grpc:foo
```

#### Prerequisites

- protoc https://developers.google.com/protocol-buffers/docs/reference/go-generated
