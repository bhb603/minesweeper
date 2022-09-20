
.PHONY: test
test:
	go test ./...

.PHONY: release-snapshot
build:
	@goreleaser build --single-target --snapshot --rm-dist

.PHONY: release
release:
	@goreleaser release --rm-dist --skip-publish
