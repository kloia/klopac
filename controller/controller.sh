#!/bin/bash
ansible-playbook -c local playbooks/img.yml
ansible-playbook -c local playbooks/ins.yml
ansible-playbook -c local playbooks/engine.yml
ansible-playbook -c local playbooks/int.yml
ansible-playbook -c local playbooks/gitops.yml
#ansible-playbook -c local playbooks/app.yml