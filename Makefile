.PHONY: grpc-go
grpc-go:
	@buf format -w ./proto
	@buf lint ./proto
	@buf generate ./proto
