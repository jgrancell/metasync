## Copyright 2020-2022 Josh Grancell. All rights reserved.
## Use of this source code is governed by an MIT License
## that can be found in the LICENSE file.
TEST?=$$(go list ./... | grep -v vendor)
WORKDIR=$$(pwd)
BINARY=$$(pwd | xargs basename)
VERSION=$$(grep version main.go | head -n1 | cut -d\" -f2)
GOBIN=${GOPATH}/bin

default: build

build:
	go build -o ${BINARY}
	chmod +x ${BINARY}

install: build
	mkdir -p ${GOBIN}
	mv ${BINARY} ${GOPATH}/bin/${BINARY}

## Testing tasks
test:
	rm -f coverage.txt profile.out
	rm -f gosec-report.json
	/bin/sh go.test.sh

test-sonarqube: test
	gosec --no-fail -fmt=sonarqube -out gosec-report.json ./...
	/opt/sonar-scanner/bin/sonar-scanner

test-view: test
	go tool cover -html=coverage.txt