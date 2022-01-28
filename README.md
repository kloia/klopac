# klopac-runner 

klopac-runner 0.1-alpha version docker image use this tools

- Ansible	(5.2.0)
- Docker	(20.10.12)
- Kubectl	(1.23.3)
- Packer    (1.7.9)
- Terraform	(1.1.4)
- Terragrunt	(0.36.0)
- Pulumi	(3.22.1)

You can use docker image like this

docker run --rm -ti -v /var/run/docker.sock:/var/run/docker.sock -v $PWD:/data  kloiadocker/klopac-runner:0.1-alpha ansible-playbook /data/example.yaml 
