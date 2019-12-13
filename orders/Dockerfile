# Code base get from https://medium.com/travis-on-docker/how-to-dockerize-your-go-golang-app-542af15c27a2

# We specify the base image we need for our
# go application
FROM golang:1.9.2
# We create an /app directory within our
# image that will hold our application source
# files
RUN mkdir /app
# We copy everything in the root directory
# into our /app directory
ADD . /app
# We specify that we now wish to execute 
# any further commands inside our /app
# directory
WORKDIR /app
# we run go build to compile the binary
# executable of our Go program
RUN go get -d -v github.com/gorilla/mux
RUN go build -o main .
# Our start command which kicks off
# our newly created binary executable
CMD ["/app/main"]

# Set your SOURCE_LOCATION with next arg in docker build
# --build-arg SOURCE_LOCATION=/Users/israel-martinez/code/src/your_go_source_files
