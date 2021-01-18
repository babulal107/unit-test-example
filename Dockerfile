ARG APP_DIR=/go/src/perennial/unit_test
FROM golang:latest AS build-stage
ARG APP_DIR
RUN mkdir -p ${APP_DIR}
COPY . ${APP_DIR}
WORKDIR ${APP_DIR}
RUN export GOPATH=/go/src/ && export CGO_ENABLED=0 && export GOOS=linux && export GOMOD111=on &&\
 go mod tidy &&\
 go build -tags netgo -a -v -o unit_test_app .

#Final Build
FROM alpine:3.12
ARG APP_DIR
WORKDIR  /go/bin
COPY --from=build-stage ${APP_DIR}/unit_test_app ./
COPY --from=build-stage ${APP_DIR}/configs/* ./configs/
RUN cat ./configs/config.yml
EXPOSE 8080/tcp
ENTRYPOINT ["/go/bin/unit_test_app"]