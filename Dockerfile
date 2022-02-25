FROM golang:1.18rc1-alpine3.15
RUN apk add git

COPY . /home/src
WORKDIR /home/src
RUN go build -o /bin/action ./
RUN go install github.com/chyroc/dropbox-cli@9147e3e && mv $(which dropbox-cli) /bin/dropbox-cli

ENTRYPOINT [ "/bin/action" ]
