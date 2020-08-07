# SPDX-FileCopyrightText: 2020 inovex GmbH <https://www.inovex.de>
#
# SPDX-License-Identifier: MIT
LINKERFLAGS = -X main.Version=`git describe --tags --always --long --dirty` -X main.BuildTimestamp=`date -u '+%Y-%m-%d_%I:%M:%S_UTC'`

test:
	go test ./... -cover -coverprofile=coverage.txt

coverage:
	make test
	go tool cover -html=coverage.txt

build:
	packr2
	go build -o main -ldflags "$(LINKERFLAGS)"
	packr2 clean

lint:
	golint ./...

zip:
	zip terraform/function.zip main

clean:
	rm -rf main terraform/function.zip smyle-integration-tests/ coverage.txt
