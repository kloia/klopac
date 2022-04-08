FROM ubuntu:20.04
ARG PACKER_VERSION=1.8.0
RUN apt-get update -y && apt-get install -y \
    software-properties-common apt-transport-https ca-certificates gnupg lsb-release curl && \
    curl -fsSL https://apt.releases.hashicorp.com/gpg | apt-key add - && \
    apt-add-repository "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main" && \
    apt-get update && apt-get install -y packer=$PACKER_VERSION && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* 
