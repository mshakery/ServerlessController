# ServerlessController

God bless us. 

# deploy and test 
1. create a docker image with provided dcoker file
2. update the docker image path in k8s yaml configs in k8s_resources files
3. cd init_etcd && go run .
4. cd .. && kubectl apply -f ./k8s_resources
5. Install Kreya
6. setup a connection in kreya to the cluster gatway with host header of authentication ingress
7. gRpc reflection is on so kreya loads all protos automatically
8. change the request in a way you like and start sending requests

## Info

Group related functions into a package (folder), each package should ideally be maintained by one person only. 

Communication is done via gRPC. Sample code is provided (todo)

## Compilation

go build
(use latest GO version 1.23.whatever)

Building in a subdir (package)? go build [filename].

Also you can `go run [filename]`

Don't push binaries to git, please <3 

## Run

./ServerlessController

## Responsibilities 

Status: Doing, Testing (written but haven't tested), Backlog, Done

| Function                       | Description                     | By     | Status |
|--------------------------------|---------------------------------|--------|--------|
| **Authentication**             | User authentication             | Arshia | Done   | 
| **Authorization**              | Check user has perms            | Arshia | Done   | 
| **WriteToEtcd**                | Self explanatory (lib)          | Arshia | Done   | 
| **AcquireLock**                | require lockid, over etcd (lib) | Arshia | Done   | 
| **ReleaseLock**                | require lockid, over etcd (lib) | Arshia | Done   | 
| **WatchLock**                  | require lockid, over etcd (lib) | Arshia | Done   | 
| **NotifyWorkerNode**           | ?                               | Masoud | Done   | 
| **NodeCheckerAwakerCron**      | ?                               | Masoud | Done   | 
| **NodeChecker**                | ?                               | Masoud | Done   | 
| **DeploymentScalerAwakerCron** | ?                               | Masoud | Done   | 
| **DeploymentScaler**           | ?                               | Masoud | Done   | 
| **Kubelet**                    |                                 | Roy    | Done   |

## Time Tracking

As per requirement of the project. Update this once in a while :) 

XP-Time: Experiment Time


| **Time Type**     | **Arshia (hours)** | **Masoud (hours)** | **Roy (hours)** |
|-------------------|--------------------|--------------------|-----------------|
| **Total-Time**    |                    |                    |                 |
| **Think-Time**    | 28                 | 33                 | 12              |
| **Dev-Time**      | 30                 | 57                 | 9               |
| **XP-Time**       | 15                 | 15                 | 3               |
| **Analysis-Time** | 8                  | 8                  | 4               |
| **Write-Time**    | 8                  | 13                 | 16              |
| **Wasted-Time**   | 11                 | 8                  | 5               |
