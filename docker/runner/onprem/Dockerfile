FROM ubuntu:20.04

RUN useradd -s /bin/bash -d /app -m -u 1000 platform && \
apt-get update -y && \
apt-get install -y curl wget unzip gnupg tar && \
apt-get clean && \
rm -rf /var/lib/apt/lists/* && \
curl -L -o - "https://github.com/vmware/govmomi/releases/download/v0.27.4/govc_Linux_x86_64.tar.gz" | tar -C /usr/local/bin -xvzf - govc

USER 1000

WORKDIR /app
