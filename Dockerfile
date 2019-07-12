FROM golang:1.12.6

WORKDIR $GOPATH

# Download all the dependencies
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
RUN go get -d -v github.com/gorilla/mux
RUN go get -d -v github.com/gorilla/handlers
RUN go get -d -v github.com/patrickmn/go-cache
RUN go get -d -v github.com/lithammer/shortuuid

# Install the package
RUN go install -v github.com/gorilla/mux
RUN go install -v github.com/gorilla/handlers
RUN go install -v github.com/patrickmn/go-cache
RUN go get -d -v github.com/lithammer/shortuuid
