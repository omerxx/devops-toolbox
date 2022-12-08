FROM golang as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o testapp .

FROM scratch
COPY --from=builder /app/testapp /
CMD ["/testapp"]
