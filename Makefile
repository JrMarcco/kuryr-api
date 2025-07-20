.PHONY: grpc
grpc:
	@buf format -w ./proto
	@buf lint ./proto
	@buf generate ./proto
