#!/bin/bash

# ANSIBLE_PYTHON_INTERPRETER=auto_silent
# ANSIBLE_DEPRECATION_WARNINGS=false
# ANSIBLE_COMMAND_WARNINGS=no
# ANSIBLE_DISPLAY_SKIPPED_HOSTS=false

# get enabled state

img=`yq -r '.img.enabled' /data/vars/image.yaml`
ins=`yq -r '.ins.enabled' /data/vars/instance.yaml`
engine=`yq -r '.engine.enabled' /data/vars/engine.yaml`
int=`yq -r '.int.enabled' /data/vars/integrations.yaml`
gitops=`yq -r '.gitops.enabled' /data/vars/gitops.yaml`
app=`yq -r '.app.enabled' /data/vars/applications.yaml`

# get operational type

if [ $img = "true" ]; then
 imgop=`yq -r '.img.operation.type' /data/vars/image.yaml`
fi

if [ $ins = "true" ]; then
 insop=`yq -r '.ins.operation.type' /data/vars/instance.yaml`
fi

if [ $engine = "true" ]; then
 engineop=`yq -r '.engine.operation.type' /data/vars/engine.yaml`
fi

if [ $int = "true" ]; then
 intop=`yq -r '.int.operation.type' /data/vars/integrations.yaml`
fi

if [ $gitops = "true" ]; then
 gitopsop=`yq -r '.gitops.operation.type' /data/vars/gitops.yaml`
fi

if [ $app = "true" ]; then
 appop=`yq -r '.app.operation.type' /data/vars/applications.yaml`
fi

#create order

if [[ -n "$imgop" ]] && [[ $imgop = "create" ]]; then
 ansible-playbook -c local playbooks/img.yml
fi

if [[ -n "$insop" ]] && [[ $insop = "create" ]]; then
 ansible-playbook -c local playbooks/ins.yml
fi

if [[ -n "$engineop" ]] && [[ $engineop = "create" ]]; then
 ansible-playbook -c local playbooks/engine.yml
fi

if [[ -n "$intop" ]] && [[ $intop = "create" ]]; then
 ansible-playbook -c local playbooks/int.yml
fi

if [[ -n "$gitopsop" ]] && [[ $gitopsop = "create" ]]; then
 ansible-playbook -c local playbooks/gitops.yml
fi

if [[ -n "$appop" ]] && [[ $appop = "create" ]]; then
 ansible-playbook -c local playbooks/app.yml
fi

#update order

if [[ -n "$appop" ]] && [[ $appop = "update" ]]; then
 ansible-playbook -c local playbooks/app.yml
fi

if [[ -n "$gitopsop" ]] && [[ $gitopsop = "update" ]]; then
 ansible-playbook -c local playbooks/gitops.yml
fi

if [[ -n "$intop" ]] && [[ $intop = "update" ]]; then
 ansible-playbook -c local playbooks/int.yml
fi

if [[ -n "$engineop" ]] && [[ $engineop = "update" ]]; then
 ansible-playbook -c local playbooks/engine.yml
fi

if [[ -n "$insop" ]] && [[ $insop = "update" ]]; then
 ansible-playbook -c local playbooks/ins.yml
fi

if [[ -n "$imgop" ]] && [[ $imgop = "update" ]]; then
 ansible-playbook -c local playbooks/img.yml
fi

#delete or clean order

if [[ -n "$appop" ]] && [[ $appop = "delete" ]]; then
 ansible-playbook -c local playbooks/app.yml
fi

if [[ -n "$gitopsop" ]] && [[ $gitopsop = "delete" ]]; then
 ansible-playbook -c local playbooks/gitops.yml
fi

if [[ -n "$intop" ]] && [[ $intop = "delete" ]]; then
 ansible-playbook -c local playbooks/int.yml
fi

if [[ -n "$engineop" ]] && [[ $engineop = "clean" ]]; then
 ansible-playbook -c local playbooks/engine.yml
fi

if [[ -n "$engineop" ]] && [[ $engineop = "delete" ]]; then
 ansible-playbook -c local playbooks/engine.yml
fi

if [[ -n "$insop" ]] && [[ $insop = "delete" ]]; then
 ansible-playbook -c local playbooks/ins.yml
fi

if [[ -n "$imgop" ]] && [[ $imgop = "delete" ]]; then
 ansible-playbook -c local playbooks/img.yml
fi