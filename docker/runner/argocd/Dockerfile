FROM ubuntu:20.04

ARG VERSION=

RUN useradd -s /bin/bash -d /app -m -u 1000 platform && \
apt-get update && \
apt-get -y install curl wget && \
wget -O /usr/local/bin/argocd https://github.com/argoproj/argo-cd/releases/download/$VERSION/argocd-linux-amd64 && \
chmod +x /usr/local/bin/argocd && \
apt-get clean && \
rm -rf /var/lib/apt/lists/* && \
echo '%platform ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers

USER 1000

WORKDIR /app