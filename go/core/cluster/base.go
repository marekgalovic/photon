package cluster

type DeploymentType struct {
    RunnerImage string
    DeployerImage string
}

type ModelDeployment struct {
    Type string
    ModelId int64
    Replicas int32
    AvailableReplicas int32
    ReadyReplicas int32
}
