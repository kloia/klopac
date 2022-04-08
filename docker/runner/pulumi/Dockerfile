FROM ubuntu:20.04
ARG PULUMI_VERSION=3.22.1

RUN apt-get update -y && \
    apt-get install -y curl wget unzip gnupg ssh apt-transport-https ca-certificates gnupg lsb-release software-properties-common sudo && \
    apt-get update && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* 

WORKDIR /app
RUN curl -fsSL https://get.pulumi.com | sudo bash -s -- --version ${PULUMI_VERSION} && \
    sudo mv /root/.pulumi /app/; sudo chown -R root:root /app/.pulumi

ENV PATH="/app/.pulumi/bin:${PATH}"


