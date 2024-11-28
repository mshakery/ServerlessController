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

Status: Doing, Testing (written but haven't tested), Backlog, Done

| Function                       | Description                     | By     | Status  |
|--------------------------------|---------------------------------|--------|---------|
| **Authentication**             | User authentication             | Arshia | Testing | 
| **Authorization**              | Check user has perms            | Arshia | Testing | 
| **WriteToEtcd**                | Self explanatory (lib)          | Arshia | Testing | 
| **ReadFromEtcd**               | Self explanatory (lib)          | Arshia | Testing | 
| **AcquireLock**                | require lockid, over etcd (lib) | Arshia | Testing | 
| **ReleaseLock**                | require lockid, over etcd (lib) | Arshia | Testing | 
| **WatchLock**                  | require lockid, over etcd (lib) | Arshia | Testing | 
| **NotifyWorkerNode**           | ?                               | ?      | Backlog | 
| **NodeCheckerAwakerCron**      | ?                               | ?      | Backlog | 
| **NodeChecker**                | ?                               | ?      | Backlog | 
| **DeploymentScalerAwakerCron** | ?                               | ?      | Backlog | 
| **DeploymentScaler**           | ?                               | ?      | Backlog | 
| **Kubelet**                    | TODO: Split to subfuncs         | Roy    | Doing   |

## Time Tracking

As per requirement of the project. Update this once in a while :) 

XP-Time: Experiment Time


| **Time Type**     | **Arshia (hours)** | **Masoud (hours)** | **Roy (hours)** |
|-------------------|--------------------|--------------------|-----------------|
| **Total-Time**    |                    |                    |                 |
| **Think-Time**    | 10                 | 8                  | 6               |
| **Dev-Time**      | 6                  | 2                  | 1               |
| **XP-Time**       | 0                  | 0                  | 0               |
| **Analysis-Time** | 0                  | 0                  | 0               |
| **Write-Time**    | 0                  | 0                  | 0               |
| **Wasted-Time**   | 5                  | 3                  | 3               |
