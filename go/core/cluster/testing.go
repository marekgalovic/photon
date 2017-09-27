package cluster

import (
    log "github.com/Sirupsen/logrus"
)

func NewTestKubernetes() *kubernetesCluster {
    kubernetes, err := NewKubernetes(KubernetesConfig{Host: "127.0.0.1", Namespace: "photon_test", Insecure: true}, map[string]*DeploymentType{"test_type": {RunnerImage: "alpine", DeployerImage: "alpine"}})
    if err != nil {
        log.Fatal(err)
    }
    return kubernetes
}
