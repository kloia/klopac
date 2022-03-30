
# KLOPAC ORCHESTRATOR RESULTS

<hr/>

## IMAGE LAYER

* Provider: AWS
* Image ID: 123456
* Healthcheck: Healthy


<hr/>

## INSTANCE LAYER


* OS Name   : ubuntu
* OS Version:   20.04.4
* Disk Utilization:  %20
* Memory Utilization : %30
* CPU Utilization : %40




<hr/>

## ENGINE LAYER

* Kubernetes Engine Type  : rke2
* Kubernetes Engine Version  : v1.21.6





<hr/>

## GITOPS LAYER

* Type: argocd
* Status: True
* Applications:
    * sock-shop-gitops
        * Healthy
        * localhost:8082
    * test-gitops
        * Unhealthy
        * localhost:8083



<hr/>

## APPLICATIONS LAYER
* sock-shop
    * Healthy
    * localhost:8080
* test
    * Unhealthy
    * localhost:8081


