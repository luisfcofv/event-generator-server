FROM golang:1.8-onbuild
EXPOSE 3000

# FROM golang
# RUN mkdir -p /go/src/github.com/luisfcofv/
# WORKDIR /go/src/github.com/luisfcofv/indexter
# COPY . /go/src/github.com/luisfcofv/indexter
# RUN go get && go build
# EXPOSE 3000
# CMD ["./indexter"]
