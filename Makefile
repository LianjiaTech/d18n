default: build

include test/Makefile.mysql
include test/Makefile.csvq
include test/Makefile.sqlite
include test/Makefile.postgres
include test/Makefile.oracle
include test/Makefile.mssql
include test/Makefile.clickhouse
include test/Makefile.presto
include test/Makefile.hive
include test/Makefile.h2
include test/Makefile.cassandra

# colors compatible setting
CRED:=$(shell tput setaf 1 2>/dev/null)
CGREEN:=$(shell tput setaf 2 2>/dev/null)
CYELLOW:=$(shell tput setaf 3 2>/dev/null)
CEND:=$(shell tput sgr0 2>/dev/null)

.PHONY: build
build: fmt
	@echo "$(CGREEN)Building ...$(CEND)"
	@mkdir -p bin
	@ret=0 && for d in $$(go list -f '{{if (eq .Name "main")}}{{.ImportPath}}{{end}}' ./...); do \
		b=$$(basename $${d}) ; \
		go build -trimpath -o bin/$${b} $$d || ret=$$? ; \
	done ; exit $$ret
	@echo "build Success!"

GO_VERSION_MIN=1.16
.PHONY: go_version_check
# Parse out the x.y or x.y.z version and output a single value x*10000+y*100+z (e.g., 1.9 is 10900)
# that allows the three components to be checked in a single comparison.
VER_TO_INT:=awk '{split(substr($$0, match ($$0, /[0-9\.]+/)), a, "."); print a[1]*10000+a[2]*100+a[3]}'
go_version_check:
	@echo "$(CGREEN)Go version check ...$(CEND)"
	@go version
	@if test $(shell go version | $(VER_TO_INT) ) -lt \
	$(shell echo "$(GO_VERSION_MIN)" | $(VER_TO_INT)); \
	then printf "go version $(GO_VERSION_MIN)+ required, found: "; go version; exit 1; \
	fi

.PHONY: release
release: build
	@echo "$(CGREEN)Cross platform building for release ...$(CEND)"
	@mkdir -p release
	@for GOOS in darwin; do \
		for GOARCH in amd64 arm64; do \
			for d in $$(go list -f '{{if (eq .Name "main")}}{{.ImportPath}}{{end}}' ./...); do \
				b=$$(basename $${d}) ; \
				echo "Building $${b}.$${GOOS}-$${GOARCH} ..."; \
				GOOS=$${GOOS} GOARCH=$${GOARCH} go build -trimpath -v -o release/$${b}.$${GOOS}-$${GOARCH} $$d 2>/dev/null; \
			done ; \
		done ;\
	done
	@for GOOS in linux windows; do \
		for GOARCH in amd64; do \
			for d in $$(go list -f '{{if (eq .Name "main")}}{{.ImportPath}}{{end}}' ./...); do \
				b=$$(basename $${d}) ; \
				echo "Building $${b}.$${GOOS}-$${GOARCH} ..."; \
				GOOS=$${GOOS} GOARCH=$${GOARCH} go build -trimpath -v -o release/$${b}.$${GOOS}-$${GOARCH} $$d ; \
			done ; \
		done ;\
	done


# Code format
.PHONY: fmt
fmt: go_version_check
	@echo "$(CGREEN)Run gofmt on all source files ...$(CEND)"
	@echo "gofmt -l -s -w ..."
	@ret=0 && for d in $$(go list -f '{{.Dir}}' ./... | grep -v /vendor/); do \
		gofmt -l -s -w $$d/*.go || ret=$$? ; \
	done ; exit $$ret

.PHONY: tidy
tidy:
	@echo "$(CGREEN)Go mod check...$(CEND)"
	go mod tidy

.PHONY: check-diff
check-diff:
	@./test/check-diff.sh

# Run golang test cases
.PHONY: test
test: mask-typo fmt
	@echo "$(CGREEN)Run all test cases ...$(CEND)"
	@#go test -timeout 10m -race ./...
	@# modernc.org/sqlite can't pass -race
	@# ==956455==ERROR: ThreadSanitizer failed to allocate 0x40000 (262144) bytes at address 600008100000 (errno: 12)
	go test -timeout 10m ./...
	@if test $(shell git diff --name-only test/ | wc -l | awk '{print $1}') -gt 0; \
	then \
		echo "Golden test file checked diff!"; \
		exit 2; \
	fi
	@echo "test Success!"

# Code Coverage
# colorful coverage numerical >=90% GREEN, <80% RED, Other YELLOW
.PHONY: cover
cover: test
	@echo "$(CGREEN)Run test cover check ...$(CEND)"
	@go test $(shell go list ./... | grep -v parser) -coverprofile=test/coverage.data | column -t
	@go tool cover -html=test/coverage.data -o test/coverage.html
	@go tool cover -func=test/coverage.data -o test/coverage.txt
	@tail -n 1 test/coverage.txt | awk '{sub(/%/, "", $$NF); \
		if($$NF < 80) \
			{print "$(CRED)"$$0"%$(CEND)"} \
		else if ($$NF >= 90) \
			{print "$(CGREEN)"$$0"%$(CEND)"} \
		else \
			{print "$(CYELLOW)"$$0"%$(CEND)"}}'

.PHONY: ci
ci: cover test-mysql test-sqlite check-diff

.PHONY: mask-typo
mask-typo:
	@echo "$(CGREEN)Auto generate mask/typo.go source file ...$(CEND)"
	@cat COPYRIGHT > mask/typo.tmp
	@echo "package mask" >> mask/typo.tmp
	@echo "" >> mask/typo.tmp
	@echo "type MaskFunc func(args ...interface{}) (ret string, err error)" >> mask/typo.tmp
	@echo "" >> mask/typo.tmp
	@echo "// maskFuncs support functions list, case insensitive" >> mask/typo.tmp
	@echo "var maskFuncs = map[string]MaskFunc{" >> mask/typo.tmp
	@go doc --short d18n/mask | grep "^func" | grep '(args' | awk -F '(' '{print $$1}'  | \
	awk '{print "\""tolower($$2)"\":", $$2","}' >> mask/typo.tmp
	@echo "}" >> mask/typo.tmp
	@mv mask/typo.tmp mask/typo.go

# check docker installed
DOCKER_CMD:=$(shell which podman 2>/dev/null || which docker 2>/dev/null)
DOCKER_RM := $(or ${DOCKER_RM}, ${DOCKER_RM}, --rm)

# read line wrap
RLWRAP:=$(shell which rlwrap 2>/dev/null || echo "")

.PHONY: docker-exist
docker-exist:
ifndef DOCKER_CMD
	@echo "please install docker/podman first"
	@exit 1
endif
	@echo "using $(DOCKER_CMD)"

.PHONY: docker-stop
docker-stop: docker-exist
	@${DOCKER_CMD} stop d18n-mysql 2>/dev/null || true
	@${DOCKER_CMD} stop d18n-postgres 2>/dev/null || true
	@${DOCKER_CMD} stop d18n-oracle 2>/dev/null || true
	@${DOCKER_CMD} stop d18n-mssql 2>/dev/null || true
	@${DOCKER_CMD} stop d18n-clickhouse 2>/dev/null || true
	@${DOCKER_CMD} stop d18n-presto 2>/dev/null || true
	@${DOCKER_CMD} volume prune --force

.PHONY: clean
clean:
	git clean -x -f
	${DOCKER_CMD} volume prune
	rm -rf bin/
	rm -rf release/
