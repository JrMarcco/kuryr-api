.PHONY: grpc
grpc:
	@buf format -w .
	@buf lint .
	@buf generate .
