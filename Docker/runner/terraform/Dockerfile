FROM ubuntu:20.04 as base
ARG TF_VERSION
ARG TG_VERSION
RUN apt-get update -y && \
    apt-get install -y curl wget unzip gnupg ssh apt-transport-https ca-certificates gnupg lsb-release software-properties-common sudo && \
    curl -fsSL https://apt.releases.hashicorp.com/gpg | sudo apt-key add - && \
    add-apt-repository "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main" && \
    apt-get update && \
    apt-get install -y terraform=$TF_VERSION && \
    wget -q -O /usr/local/bin/terragrunt "https://github.com/gruntwork-io/terragrunt/releases/download/${TG_VERSION}/terragrunt_linux_amd64" && \
    chmod +x /usr/local/bin/terragrunt && \
    apt-get clean

WORKDIR /app
CMD [ "sh", "-c","terraform init && terraform plan -out 'tfplan' && terraform apply tfplan"]