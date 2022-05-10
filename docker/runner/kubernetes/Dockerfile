FROM ubuntu:20.04
RUN apt-get update -y && apt-get install -y \
    software-properties-common apt-transport-https ca-certificates gnupg lsb-release curl sudo && \
    curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" && \
    install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl && \
    curl -s "https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3" | bash && \
    curl -s "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh" | bash && \
    mv /kustomize /usr/local/bin/ && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* 
