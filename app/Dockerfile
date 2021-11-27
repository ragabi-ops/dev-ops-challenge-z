#Using go image to build the code
FROM golang:1.17.3-alpine3.13 AS GO_BUILD
COPY app/ /server
WORKDIR /server
RUN go build -o /go/bin/go-app

FROM alpine:3.14.0
WORKDIR /app
#Providing Dummy values to work with aws-go-sdk
ENV AWS_ACCESS_KEY_ID="thisisaccessdummy"
ENV AWS_SECRET_ACCESS_KEY="thisissecretdummy"
#Copy binary file from go builder 
COPY --from=GO_BUILD /go/bin/go-app ./
CMD ./go-app
