FROM golang:1.17-alpine as build
WORKDIR /app
COPY ./entrypoint/ .
RUN CGO_ENABLED=0 go build ./cmd/main.go


FROM ubuntu:20.04

RUN groupadd docker -g 998 && useradd -s /bin/bash -d /app -m -u 1000 -G docker platform && \
apt-get update -y && \
apt-get install -y curl wget unzip gnupg ssh apt-transport-https ca-certificates gnupg lsb-release software-properties-common sudo && \
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add - && \
curl -fsSL https://apt.releases.hashicorp.com/gpg | sudo apt-key add - && \
add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" && \
add-apt-repository "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main" && \
apt-get update && \
apt-get install -y python3-pip=20.0.2-5ubuntu1.6 docker-ce=5:20.10.12~3-0~ubuntu-focal terraform=1.1.4 packer=1.7.9 && \
wget -q -O /usr/local/bin/terragrunt "https://github.com/gruntwork-io/terragrunt/releases/download/v0.36.0/terragrunt_linux_amd64" && \
chmod +x /usr/local/bin/terragrunt && \
curl -LO "https://dl.k8s.io/release/v1.23.3/bin/linux/amd64/kubectl" && \
curl -LO "https://dl.k8s.io/v1.23.3/bin/linux/amd64/kubectl.sha256" && \
install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl && \
chmod +x /usr/local/bin/kubectl && \
curl -LO "https://get.helm.sh/helm-v3.8.0-linux-amd64.tar.gz" && \
tar -zxvf helm-v3.8.0-linux-amd64.tar.gz && \
mv linux-amd64/helm /usr/local/bin/helm && \
chmod +x /usr/local/bin/helm && \
apt-get clean && \
echo '%docker ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers && \
pip3 install ansible==5.2.0

USER 1000

RUN curl -fsSL https://get.pulumi.com | sudo bash -s -- --version 3.22.1 --silent && \
sudo mv /root/.pulumi /app/; sudo chown -R platform:platform /app/.pulumi; echo "export PATH=$PATH:~/.pulumi/bin" >> /app/.bashrc

WORKDIR /app

COPY --from=build /app/main /usr/local/bin/entrypoint
RUN chmod +x /usr/local/bin/entrypoint

COPY . .
ENTRYPOINT ["entrypoint"]


