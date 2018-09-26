FROM golang:1.11 AS builder
ENV GOPATH=/gocode
RUN mkdir -p /go-study-group/agenda-api
WORKDIR /go-study-group/agenda-api
COPY . .
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOPROXY=https://microsoftgoproxy.azurewebsites.net
RUN go build -o /bin/agenda-api .

FROM alpine

COPY --from=builder /bin/agenda-api /bin/agenda-api

RUN apk update && apk add ca-certificates

ENV AGENDA_API_PORT=3000

EXPOSE 3000

CMD ["/bin/agenda-api"]
