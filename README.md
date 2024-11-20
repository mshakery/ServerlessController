# ServerlessController

God bless us. 

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

| Function                         | Description                     | By     | Status  |
|----------------------------------|---------------------------------|--------|---------|
| **Authentication**               | User authentication             | Arshia | Backlog | 
| **Authorization**                | Check user has perms            | Arshia | Backlog | 
| **WriteToEtcd**                  | Self explanatory (lib)          | Arshia | Doing   | 
| **ReadFromEtcd**                 | Self explanatory (lib)          | Arshia | Doing   | 
| **AcquireLock**                  | require lockid, over etcd (lib) | Arshia | Backlog | 
| **ReleaseLock**                  | require lockid, over etcd (lib) | Arshia | Backlog | 
| **NotifyWorkerNode**             | ?                               | ?      | Backlog | 
| **NodeCheckerAwakerCron**        | ?                               | ?      | Backlog | 
| **NodeChecker**                  | ?                               | ?      | Backlog | 
| **DeploymentScalerAwakerCron**   | ?                               | ?      | Backlog | 
| **DeploymentScaler**             | ?                               | ?      | Backlog | 
| **Kubelet**                      | TODO: Split to subfuncs         | Roy    | Doing   |

## Time Tracking

As per requirement of the project. Update this once in a while :) 

XP-Time: Experiment Time


| **Time Type**     | **Arshia (hours)** | **Masoud (hours)** | **Roy (hours)** |
|-------------------|--------------------|--------------------|-----------------|
| **Total-Time**    |                    |                    |                 |
| **Think-Time**    | 6                  | 6                  | 6               |
| **Dev-Time**      | 0                  | 1                  | 0               |
| **XP-Time**       | 0                  | 0                  | 0               |
| **Analysis-Time** | 0                  | 0                  | 0               |
| **Write-Time**    | 0                  | 0                  | 0               |
| **Wasted-Time**   | 5                  | 3                  | 3               |
