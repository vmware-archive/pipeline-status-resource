FROM cfmobile/toronto-ci

ENV LANG C

ADD http://stedolan.github.io/jq/download/linux64/jq /usr/local/bin/jq
RUN chmod +x /usr/local/bin/jq

ADD assets/ /opt/resource/
RUN chmod +rx /opt/resource
RUN chmod +rx /opt/resource/*

ADD concourse.go /opt/go/concourse.go
ADD src/ /opt/go/src/
RUN chmod +x /opt/go
RUN	apt-get -qq install golang
RUN	GOPATH=/opt/go go build -o /opt/go/runme /opt/go/concourse.go
