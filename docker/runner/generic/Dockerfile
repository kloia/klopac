FROM ubuntu:20.04
ARG PYTHON_VERSION=3.9
ARG GOLANG_VERSION=1.17.8
ARG BASH_VERSION=5.0-6ubuntu1.1
RUN apt-get update -y && apt-get install -y \ 
    curl python$PYTHON_VERSION python3-distutils bash=$BASH_VERSION && \
    ln -s /usr/bin/python$PYTHON_VERSION /usr/bin/python && \
    curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py && python get-pip.py && \
    curl -OL https://go.dev/dl/go$GOLANG_VERSION.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go$GOLANG_VERSION.linux-amd64.tar.gz && \
    rm -rf *.tar.gz && \
    rm -rf /var/lib/apt/lists/* && 
ENV PATH=/usr/local/go/bin:$PATH