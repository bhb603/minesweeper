
.PHONY: test
test:
	go test ./...

.PHONY: release-snapshot
build:
	@goreleaser build \
		--snapshot \
		--rm-dist \
		--single-target

.PHONY: release
release:
	@goreleaser release \
		--snapshot \
		--rm-dist \
		--skip-publish
