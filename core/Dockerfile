FROM golang:1.22.5-bookworm as build-stage
COPY ./core/ /usr/src/core/
WORKDIR /usr/src/core/
RUN go mod tidy
RUN go mod download
RUN go mod tidy
RUN go build -o "core"

FROM debian:bookworm as production-stage 
RUN apt update && apt upgrade -y && apt install -y ca-certificates
COPY --from=build-stage /usr/src/core/core /root/

CMD [ "/root/core" ]
