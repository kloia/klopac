FROM golang:1.18-alpine as build
WORKDIR /app
COPY entrypoint ./entrypoint
RUN cd entrypoint && \
    CGO_ENABLED=0 go build ./cmd/main.go


FROM ubuntu:20.04

RUN groupadd docker -g 998 && useradd -s /bin/bash -d /app -m -u 1000 -G docker platform && \
apt-get update -y && \
apt-get install -y --no-install-recommends curl wget unzip ca-certificates gnupg lsb-release sudo git && \
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg && \
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null && \
apt-get update && \
apt-get install -y --no-install-recommends python3-pip=20.0.2-5ubuntu1.6 docker-ce=5:20.10.12~3-0~ubuntu-focal jq && \
apt-get update && \
apt-get install -y apt-transport-https ca-certificates curl && \
curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg && \
echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list && \
apt-get update && \
apt-get install -y kubectl && \
apt-get clean && \
rm -rf /var/lib/apt/lists/* && \
echo '%docker ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers && \
mkdir -p /data/repo /data/outputs /data/sensitive && \
chown -R 1000:1000 /data && \
pip3 --no-cache-dir install ansible==5.2.0 yq boto3 botocore boto docker && \
apt-get update && \
apt-get install -y sshpass

USER 1000

WORKDIR /app

COPY --from=build /app/entrypoint/main /usr/local/bin/entrypoint
RUN sudo chmod +x /usr/local/bin/entrypoint

COPY --chown=1000:1000 vars /data/vars
COPY --chown=1000:1000 manifests /data/manifests
COPY --chown=1000:1000 entrypoint ./entrypoint
COPY --chown=1000:1000 provisioner ./provisioner
COPY --chown=1000:1000 validator ./validator
COPY --chown=1000:1000 controller ./controller
COPY --chown=1000:1000 finalizer ./finalizer
#ENTRYPOINT ["entrypoint"]