#!/bin/sh

init() {
	set -e

	exec 3>&1 # make stdout available as fd 3 for the result
	exec 1>&2 # redirect all output to stderr for logging

	# for jq
	PATH=/usr/local/bin:$PATH

	payload=$TMPDIR/resource-definition

	cat > "$payload" <&0

	export HOST
	HOST=$(jq -r '.source.host // ""' < "$payload")

	export USER
	USER=$(jq -r '.source.user // ""' < "$payload")

	export PASSWORD
	PASSWORD=$(jq -r '.source.password // ""' < "$payload")

	export PIPELINE_WHITELIST
	PIPELINE_WHITELIST=$(jq -r '.source.pipeline_whitelist // ""' < "$payload")
}

fetch-status() {
	apt-get -qq install golang
	cd concourse-status-hue
	export GOPATH
	GOPATH=$(pwd)
	go build -o runme ./concourse.go
	./runme -host="$HOST" -user="$USER" -password="$PASSWORD" "$PIPELINE_WHITELIST"
}
