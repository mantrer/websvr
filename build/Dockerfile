# STEP 1 build executable binary
FROM golang:alpine as builder
# Create appuser
RUN adduser -D -g '' appuser

FROM scratch

USER appuser
ENV PORT 58080
EXPOSE $PORT

COPY --from=builder /etc/passwd /etc/passwd
# Copy our static executable
COPY websvr /websvr

CMD ["/websvr"]


