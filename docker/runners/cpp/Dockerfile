FROM alpine:3.19

RUN apk add --no-cache g++ make

WORKDIR /sandbox

# Script to compile and run
COPY runner.sh /runner.sh
RUN chmod +x /runner.sh

ENTRYPOINT ["/runner.sh"]
