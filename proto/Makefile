.PHONY: protolint
protolint:
	buf format -w metalstack/io/accounting/api/v1
	buf lint -v

.PHONY: protoc
protoc: protolint
	buf generate -v