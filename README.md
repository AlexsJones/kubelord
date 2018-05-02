```                                            
    )           )       (             (     
 ( /(    (   ( /(    (  )\      (     )\ )  
 )\())  ))\  )\())  ))\((_) (   )(   (()/(  
((_)\  /((_)((_)\  /((_)_   )\ (()\   ((_))
| |(_)(_))( | |(_)(_)) | | ((_) ((_)  _| |  
| / / | || || '_ \/ -_)| |/ _ \| '_|/ _` |  
|_\_\  \_,_||_.__/\___||_|\___/|_|  \__,_|  

```

Kubernetes CLI dashboard

- A simple dashboard for people who really work with kubernetes every single day across many clusters.

## usage

`Glide` is used for package management 

```
go get github.com/AlexsJones/kubelord
```

`kubelord`

`q` to quit the dashboard

Enjoy a simple and useful dashboard for your current cluster context
`Updates periodically`

```
┌─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│  Namespace   | Deployments              | Type       | Replicas | Status                                                                        │
│─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────│
│  kube-system | event-exporter-v0.1.7    | Deployment | 1/1      | Deployment has minimum availability.                                          │
│─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────│
│  kube-system | heapster-v1.5.0          | Deployment | 1/1      | Deployment has minimum availability.                                          │
│─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────│
│  kube-system | kube-dns                 | Deployment | 1/1      | Deployment has minimum availability.                                          │
│─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────│
│  kube-system | kube-dns-autoscaler      | Deployment | 1/1      | Deployment has minimum availability.                                          │
│─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────│
│  kube-system | kube-state-metrics       | Deployment | 1/1      | Deployment has minimum availability.                                          │
│─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────│
│  kube-system | kubernetes-dashboard     | Deployment | 1/1      | ReplicaSet "kubernetes-dashboard-59b7589566" has successfully progressed.     │
│─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────│
│  kube-system | l7-default-backend       | Deployment | 1/1      | Deployment has minimum availability.                                          │
│─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────│
│  kube-system | metrics-server-v0.2.1    | Deployment | 1/1      | Deployment has minimum availability.                                          │
│─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────│
│  preview     | api-core                 | Deployment | 2/2      | ReplicaSet "api-core-5565b84f77" has successfully progressed.                 │
│─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────│
│  preview     | api-mail                 | Deployment | 3/3      | Deployment has minimum availability.                                          │
│─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────│
│  preview     | api-notifications        | Deployment | 1/1      | ReplicaSet "api-notifications-6cf6b6d598" has successfully progressed.        │
│─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────│
│  preview     | api-portal               | Deployment | 2/2      | ReplicaSet "api-portal-7cbfd97f84" has successfully progressed.               │
│─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────│
│  preview     | app-admin                | Deployment | 1/1      | ReplicaSet "app-admin-555f959c69" has successfully progressed.                │
│─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────│
│  preview     | app-bath                 | Deployment | 1/1      | ReplicaSet "app-bath-b7cc8bfb6" has successfully progressed.                  │
│─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────│
│  preview     | app-beamery              | Deployment | 2/2      | Deployment has minimum availability.                                          │
│─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────│
│  preview     | app-pages                | Deployment | 1/1      | ReplicaSet "app-pages-568c67d6c8" has successfully progressed.                │
│─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────│
│  preview     | app-portal               | Deployment | 1/1      | ReplicaSet "app-portal-985cb5676" has successfully progressed.                │
│─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────│
│  preview     | app-site                 | Deployment | 1/1      | ReplicaSet "app-site-54444975b9" has successfully progressed.                 │
└─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
```
