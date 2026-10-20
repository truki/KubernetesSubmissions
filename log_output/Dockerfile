FROM golang:1.24-alpine AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o log_output .

FROM scratch
COPY --from=build /app/log_output /log_output
CMD ["/log_output"]
