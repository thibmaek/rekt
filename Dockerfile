# Build rekt-cli binary
FROM golang:1.21 AS builder

ARG REKT_CLI_VERSION=1.0.0-alpha

WORKDIR /build
COPY rekt-cli .
COPY Makefile .
RUN go build -ldflags="-X 'main.Version=${REKT_CLI_VERSION}'" -o rekt .

# Install jadx & hermes toolchain
FROM openjdk:8-jre-slim

LABEL description="Rekt is app rekking tool for Android APKs"
LABEL repository="https://github.com/thibmaek/rekt"
LABEL maintainer="thibmaek"

ARG JADX_VERSION=1.4.7

RUN apt update -y && \
    apt install -y curl git file python3 python3-pip unzip && \
    rm -rf /var/lib/apt/lists/*
RUN ln -s /usr/bin/python3 /usr/bin/python
RUN python -m pip install --upgrade pip

RUN pip install --upgrade git+https://github.com/P1sec/hermes-dec

RUN mkdir -p /opt/jadx && \
    curl -L https://github.com/skylot/jadx/releases/download/v${JADX_VERSION}/jadx-${JADX_VERSION}.zip -o /opt/jadx/jadx-${JADX_VERSION}.zip && \
    unzip /opt/jadx/jadx-${JADX_VERSION}.zip -d /opt/jadx && \
    rm /opt/jadx/jadx-${JADX_VERSION}.zip

ENV PATH="/opt/jadx/bin:${PATH}"

# Add Docker image assets
COPY --from=builder /build/rekt /usr/bin/rekt

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
