FROM ubuntu:20.04

RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y python3-pip ssh apt-transport-https ca-certificates curl

RUN pip3 install ansible

# Install kubectl
RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
RUN curl -LO "https://dl.k8s.io/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl.sha256"
RUN install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
RUN chmod +x /usr/local/bin/kubectl

# Install HELM
RUN curl -LO "https://get.helm.sh/helm-v3.7.2-linux-amd64.tar.gz"
RUN tar -zxvf helm-v3.7.2-linux-amd64.tar.gz
RUN mv linux-amd64/helm /usr/local/bin/helm
RUN chmod +x /usr/local/bin/helm

WORKDIR /app

ENTRYPOINT ["/bin/bash"]
