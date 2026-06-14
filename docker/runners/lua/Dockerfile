FROM ubuntu:22.04

RUN apt-get update && apt-get install -y \
    lua5.4 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY runner.sh /runner.sh
RUN chmod +x /runner.sh

ENTRYPOINT ["/runner.sh"]
