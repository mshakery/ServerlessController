FROM golang:1.23.4-alpine3.20

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .


RUN for function in "authentication"  "authorization"   "kubelet"  "metricCollector"  "metricCollectorCronjob"   "scheduler"  "writeToEtcd"; do cd $function; go build .; cd ..; done