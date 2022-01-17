# Klopac Ansible

## Overview

This playbook installs stable version of Rancher with kubectl and HELM CLIs.

Currently it uses privateCA signed certificates for TLS options.

## How to use

### Prereqs

- Docker installed on jump server

### First steps
1. Clone this repository into jumper server
2. Build docker image with `docker build . -t rancher-ansible`
3. Update variables on `vars/general-config.yml` if necessary
4. Copy your clusters kubeconfig to root path of repository. When connecting to clusters, playbooks uses `kubeconfig` file by default. You can change this on `./vars/general-config.yml`. 

### Longhorn Disk Mounts
1. Copy your id_rsa file to repo's root directory. 
2. Add instance ip addresses to inventory.
3. Run playbook with `docker run -it -v $(pwd):/app rancher-ansible "./provision_longhorn_disks.sh`

### Rancher Installation
1. Copy your CA certficate, TLS key and certificate files to the root of the repository (e.g. /home/my-user/rancher-ansible/tls.crt ca.crt and tls.key)
2. Copy your kubeconfig file to the root of this repository **(You may need to set insecure-skip-tls-verify: true and remove certificate-authority-data if you're getting TLS errors)**
3. Update variables on `vars/rancher-config.yml` if necessary **(Don't forget to update hostname variable)**
4. Run `docker run -it -v $(pwd):/app rancher-ansible "./provision_rancher.sh"`

### Monitoring Installation
1. Update variables on `vars/monitoring-config` according to your installation and infrastructure
2. Run `docker run -it -v $(pwd):/app rancher-ansible "./provision_rancher_monitoring.sh`

### Logging Installation

This playbook installs fluentd on all nodes(including masters). There is a feature enabled for reading
 `/var/log/auth.log` and this feature gets root privilages on node. Please see https://github.com/fluen
t/fluentd-kubernetes-daemonset#run-as-root for details.

1. Update fluentd configuration on `roles/logging/files` if necessary.
2. Run `docker run -it -v $(pwd):/app rancher-ansible "./provision_logging.sh`
## How to contribute

You can check TODO.md for contributions and suggestions.
