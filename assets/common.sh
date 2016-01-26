#!/bin/bash

init() {
	set -ex

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

	export LAST_CHECK_DATE
	LAST_CHECK_DATE=$(jq -r '.version.date // ""' < "$payload")
}

fetch_status() {
	pushd /opt/go
	./runme -host="$HOST" -user="$USER" -password="$PASSWORD" $PIPELINE_WHITELIST
	popd
}
