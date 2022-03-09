FROM golang:1.17-alpine as build
WORKDIR /app
COPY . .
RUN cd entrypoint && \
    CGO_ENABLED=0 go build ./cmd/main.go


FROM ubuntu:20.04

RUN groupadd docker -g 998 && useradd -s /bin/bash -d /app -m -u 1000 -G docker platform && \
apt-get update -y && \
apt-get install -y curl wget unzip gnupg ssh apt-transport-https ca-certificates gnupg lsb-release software-properties-common sudo && \
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add - && \
add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" && \
apt-get update && \
apt-get install -y python3-pip=20.0.2-5ubuntu1.6 docker-ce=5:20.10.12~3-0~ubuntu-focal && \
curl -LO "https://dl.k8s.io/release/v1.23.3/bin/linux/amd64/kubectl" && \
curl -LO "https://dl.k8s.io/v1.23.3/bin/linux/amd64/kubectl.sha256" && \
install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl && \
chmod +x /usr/local/bin/kubectl && \
apt-get clean && \
echo '%docker ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers && \
pip3 install ansible==5.2.0

USER 1000

WORKDIR /app

COPY --from=build /app/entrypoint/main /usr/local/bin/entrypoint
RUN sudo chmod +x /usr/local/bin/entrypoint

COPY . .
ENTRYPOINT ["entrypoint"]
