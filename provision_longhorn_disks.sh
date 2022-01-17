#!/bin/bash -e

PRIVATE_KEY_PATH="./id_rsa"

ansible-playbook -i inventory --private-key $PRIVATE_KEY_PATH \
    playbooks/longhorn-disk-mount-playbook.yml
