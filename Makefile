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

binaries: build
	rm -rf packaging/binaries
	mkdir -p packaging/binaries
	bash packaging/generate-binaries.sh ${BINARY} ${WORKDIR}

package: binaries
	bash packaging/generate-containers.sh ${WORKDIR}/packaging

test:
	rm -f coverage.txt profile.out
	rm -f gosec-report.json
	/bin/sh go.test.sh

test-sonarqube: test
	gosec --no-fail -fmt=sonarqube -out gosec-report.json ./...
	/opt/sonar-scanner/bin/sonar-scanner

test-view: test
	go tool cover -html=coverage.txt

clean:
	rm -rf packaging/binaries
	rm -rf packaging/workdir