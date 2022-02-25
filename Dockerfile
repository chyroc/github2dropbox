FROM golang:1.18rc1-alpine3.15
RUN apk add git

COPY . /home/src
WORKDIR /home/src
RUN go build -o /bin/action ./
RUN go install github.com/chyroc/dropbox-cli@37e6603 && mv $(which dropbox-cli) /bin/dropbox-cli

ENTRYPOINT [ "/bin/action" ]
