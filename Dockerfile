FROM golang:alpine
COPY test .
RUN ls -la
