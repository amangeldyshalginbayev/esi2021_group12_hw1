## We specify the base image we need for our
## go application
FROM golang:1.12.0-alpine3.9
## We create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
ADD . /app
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
WORKDIR /app
## we run go build to compile the binary
## executable of our Go program
RUN go mod download
RUN go build -o main .
## Our start command which kicks off
## our newly created binary executable
CMD ["/app/main"]
# CMD ["go","run","/app/main.go"]


# FROM golang:alpine

# # Set necessary environmet variables needed for our image
# ENV GO111MODULE=on \
#     CGO_ENABLED=0 \
#     GOOS=linux \
#     GOARCH=amd64

# # Move to working directory /build
# WORKDIR /build

# # Copy and download dependency using go mod
# COPY go.mod .
# COPY go.sum .
# RUN go mod download

# # Copy the code into the container
# COPY . .

# # Build the application
# RUN go build -o main .

# # Move to /dist directory as the place for resulting binary folder
# WORKDIR /dist

# # Copy binary from build to main folder
# RUN cp /build/main .

# # Export necessary port
# EXPOSE 3000

# # Command to run when starting the container
# CMD ["/dist/main"]