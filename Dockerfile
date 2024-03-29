FROM golang:1.17.1-buster as builder

# Create appuser.
ENV USER=appuser
ENV UID=10001
# See https://stackoverflow.com/a/55757473/12429735RUN
RUN adduser \
    --disabled-password \
    --gecos "" \
    --no-create-home \
    --shell "/sbin/nologin" \
    --uid "${UID}" \
    "${USER}"

WORKDIR /app
COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o server -ldflags="-s -w" ./cmd/server

FROM bitnami/minideb:buster

COPY --from=builder /app/data/sqlite/migrations/ /migrations
COPY --from=builder /app/server /bin/
COPY --from=builder /app/data /app/data
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

RUN chown -R appuser:appuser /app/data && \
	mkdir /paydb && chown -R appuser:appuser /paydb
VOLUME /paydb

USER appuser:appuser

EXPOSE 8446

CMD ["server"]
