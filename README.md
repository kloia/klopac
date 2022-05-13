![Klopac Media Banner](img/klopac-banner.jpg)

Klopac is a tool that let's you easily provision your infrastructure and applications. Cloud or on-premise it can schedule its runners anywhere so its extendable and scalable. Making job easier with default configurations but also lets users fine tune with YAML formatted configuration files.

## Releases

Current and upcoming releases listed below.

| Release | Release Date |
|:---:|:---:|
| 0.1-alpha | 16.05.2022 |

## Getting Started

- Clone klopac repository
- Set your working dir as repository folder
- Create two folders named 'bundle' and 'inputs'
- Create a file named 'input.yaml' under 'inputs' folder (You can use example input.yaml content below for AWS)

```
---

# inputs.yaml example

engine:
  enabled: true
gitops:
  enabled: false
img:
  enabled: true
ins:
  enabled: true
int:
  enabled: false
platform:
  environment: dev
  name: klopac
  provider:
    auth:
      id: ''
      key: ''
      type: id
    aws:
      iam:
        policy: eks-all
      role:
        arn: <eks-admin-role-arn>
      ec2:
        master:
          instanceType: t2.large
        worker:
          instanceType: t2.large
      vpc:
        cidrBlock: 172.31.0.0/16
        id: <vpc-id>
        subnet:
          id1: <subnet1-id>
          id2: <subnet2-id>
    name: aws
    region1: <region1-id>
    region2: <region2-id>
    type: ec2
    sourceimg: <ami-id>
```
- Execute command below. Whole platform creation should take 20-23 minutes after that.

```
docker run --rm -v $PWD/inputs:/data/inputs -v /var/run/docker.sock:/var/run/docker.sock -v /tmp:/data/repo -v $PWD/bundle:/data/bundle kloiadocker/klopac-runner:latest entrypoint --valuesFile=/data/inputs/input.yaml
```

- If command doesn't encounter any errors files below should get created under 'bundle' folder.

```
-rw-r--r--  1 user  staff    27K date bundle.tar.gz
-rw-r--r--@ 1 user  staff   2.9K date klopac-dev.kubeconfig
-rw-r-----@ 1 user  staff   584B date output.md
-rwxr-x---@ 1 user  staff   674B date output.yaml
-rw-r--r--@ 1 user  staff    70K date terraform.tfstate
```

## Get Involved

Klopac repository licensed under MIT. Such contributions are always welcome. To fill a bug, suggest an improvement or request a new feature open an issue against Klopac repository. 

Refer to our contribution guideline for more information about how you can get involved.

## License

This repository (Klopac) licensed under MIT.

