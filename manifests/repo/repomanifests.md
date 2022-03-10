# KloPaC Repository Manifests

### **KloPaC Controller** uses managed repos during the platform process



![Layered](../../img/ControllerFlow.png)


>### Some Managed Repos 

* [RKE2-Ansible](rke2-ansible%40main.yaml): Repo install RKE2 Cluster on target machines.
* [RKE-Ansible](rke-ansible%40master.yaml): Repo install RKE Cluster on target machines.
* [Rancher-Ansible](rancher-ansible%40main.yaml): Repo install monitoring and logging tools.

> All managed repositories are defined in **branch or version based yaml files**.

> All managed repositories must support **input and output standards**.