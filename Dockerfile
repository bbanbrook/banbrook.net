# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /Users/bb250226/Google Drive/Workspace/go_work/gopl.io.git/trunk/ch1/server2 

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install /Users/bb250226/Google Drive/Workspace/go_work/gopl.io.git/trunk/ch1/server2 

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/cserver

# Document that the service listens on port 8000.
EXPOSE 8000
