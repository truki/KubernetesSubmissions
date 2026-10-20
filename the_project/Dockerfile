# Build stage: compile the Go binary using the official Go image.
FROM golang:1.24-alpine AS build
WORKDIR /app
COPY . .
# CGO_ENABLED=0 produces a static binary that can run without any system libraries.
RUN CGO_ENABLED=0 go build -o the_project .

# Final stage: an empty image containing only the binary.
FROM scratch
COPY --from=build /app/the_project /the_project
CMD ["/the_project"]
