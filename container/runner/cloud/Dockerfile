FROM ubuntu:20.04
RUN apt-get update -y && apt-get install curl unzip python3 -y
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" && \
    unzip awscliv2.zip && \ 
    ./aws/install && \
    curl "https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-380.0.0-linux-x86_64.tar.gz" -o "google-cloud.tar.gz" && \
    tar -xf google-cloud.tar.gz && \
    ./google-cloud-sdk/install.sh -q && \
    curl -sL https://aka.ms/InstallAzureCLIDeb | bash
ENV PATH=/google-cloud-sdk/bin:$PATH
ENTRYPOINT ["/bin/bash", "-c"]
