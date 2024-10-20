FROM golang:latest
RUN apt update
RUN apt install tree

WORKDIR ./app
COPY app/ ./app