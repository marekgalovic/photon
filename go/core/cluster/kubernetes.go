package cluster

import (
    "fmt";
    "time";

    "github.com/marekgalovic/photon/go/core/metrics";

    rest "k8s.io/client-go/rest";
    kubernetes "k8s.io/client-go/kubernetes";
    k8sAppsV1beta1 "k8s.io/client-go/kubernetes/typed/apps/v1beta1";
    appsv1beta1 "k8s.io/api/apps/v1beta1";
    apiv1 "k8s.io/api/core/v1";
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1";
    kubeError "k8s.io/apimachinery/pkg/api/errors";
)

type KubernetesConfig struct {
    Host string
    Namespace string
    Insecure bool
    CertificateAuthority string
    CertificateFile string
    KeyFile string
}

type Kubernetes interface {
    HasType(string) bool
    GetType(string) *DeploymentType
    ListModels() ([]*ModelDeployment, error)
    DeployModel(*ModelDeployment) error
    UndeployModel(int64) error
    ScaleModel(int64, int32) error
}

type kubernetesCluster struct {
    client *kubernetes.Clientset
    deploymentsClient k8sAppsV1beta1.DeploymentInterface
    deploymentTypes map[string]*DeploymentType
}

func NewKubernetes(config KubernetesConfig, deploymentTypes map[string]*DeploymentType) (*kubernetesCluster, error) {
    kubeConfig := &rest.Config {
        Host: config.Host,
        TLSClientConfig: rest.TLSClientConfig{
            Insecure: config.Insecure,
            CAFile: config.CertificateAuthority,
            CertFile: config.CertificateFile,
            KeyFile: config.KeyFile,
        },
    }

    client, err := kubernetes.NewForConfig(kubeConfig)
    if err != nil {
        return nil, err
    }

    return &kubernetesCluster {
        client: client,
        deploymentsClient: client.AppsV1beta1().Deployments(config.Namespace),
        deploymentTypes: deploymentTypes,
    }, nil
}

func (k *kubernetesCluster) HasType(typeName string) bool {
    _, exists := k.deploymentTypes[typeName]
    return exists
}

func (k *kubernetesCluster) GetType(typeName string) *DeploymentType {
    return k.deploymentTypes[typeName]
}

func (k *kubernetesCluster) ListModels() ([]*ModelDeployment, error) {
    defer metrics.Runtime("kubernetes.runtime", []string{"method:list_models"})()

    list, err := k.deploymentsClient.List(metav1.ListOptions{})
    if err != nil {
        return nil, err
    }

    modelDeployments := make([]*ModelDeployment, len(list.Items))
    for i, item := range list.Items {
        modelDeployments[i] = k.convertDeployment(item)
    }
    
    return modelDeployments, nil
}

func (k *kubernetesCluster) DeployModel(deployemntSpec *ModelDeployment) error {
    defer metrics.Runtime("kubernetes.runtime", []string{"method:deploy_model"})()

    _, err := k.deploymentsClient.Create(k.modelDeploymentTemplate(deployemntSpec))
    return err
}

func (k *kubernetesCluster) UndeployModel(modelId int64) error {
    defer metrics.Runtime("kubernetes.runtime", []string{"method:undeploy_model"})()

    deletePolicy := metav1.DeletePropagationForeground
    return k.deploymentsClient.Delete(k.modelDeploymentName(modelId), &metav1.DeleteOptions{
        PropagationPolicy: &deletePolicy,
    })
}

func (k *kubernetesCluster) ScaleModel(modelId int64, replicas int32) error {
    defer metrics.Runtime("kubernetes.runtime", []string{"method:scale_model"})()

    deployment, err := k.deploymentsClient.Get(k.modelDeploymentName(modelId), metav1.GetOptions{})
    if err != nil {
        return err
    }
    deployment.Spec.Replicas = &replicas

    retries := 0
    for {
        _, err = k.deploymentsClient.Update(deployment)
        if kubeError.IsConflict(err) && retries < 5 {
            retries += 1
            time.Sleep(time.Duration(250 * retries) * time.Millisecond)
            continue
        }
        if err != nil {
            return err
        }
        return nil
    }
}

func (k *kubernetesCluster) modelDeploymentName(id int64) string {
    return fmt.Sprintf("photon-model-%d", id)
}

func (k *kubernetesCluster) modelDeploymentTemplate(deploymentSpec *ModelDeployment) *appsv1beta1.Deployment {
    return &appsv1beta1.Deployment{
        ObjectMeta: metav1.ObjectMeta{
            Name: k.modelDeploymentName(deploymentSpec.ModelId),
        },
        Spec: appsv1beta1.DeploymentSpec{
            Replicas: &deploymentSpec.Replicas,
            Template: apiv1.PodTemplateSpec{
                ObjectMeta: metav1.ObjectMeta{
                    Labels: map[string]string{"model_id": fmt.Sprintf("%d", deploymentSpec.ModelId), "runner_type": deploymentSpec.Type},
                },
                Spec: apiv1.PodSpec{
                    RestartPolicy: apiv1.RestartPolicyAlways,
                    Containers: []apiv1.Container{
                        {
                            Name: "runner",
                            Image: k.GetType(deploymentSpec.Type).RunnerImage,
                            Env: []apiv1.EnvVar{
                                {Name: "PHOTON_ENV", Value: "production"},
                                {Name: "PHOTON_MODELS_DIR", Value: "/models"},
                            },
                            Ports: []apiv1.ContainerPort{
                                {Name: "http", Protocol: apiv1.ProtocolTCP, ContainerPort: 80},
                            },
                            VolumeMounts: []apiv1.VolumeMount{
                                {Name: "models-dir", MountPath: "/models"},
                            },
                        },
                        {
                            Name: "deployer",
                            Image: k.GetType(deploymentSpec.Type).DeployerImage,
                            Env: []apiv1.EnvVar{
                                {Name: "PHOTON_ENV", Value: "production"},
                                {Name: "PHOTON_MODELS_DIR", Value: "/models"},
                                {Name: "PHOTON_MODEL_ID", Value: fmt.Sprintf("%d", deploymentSpec.ModelId)},
                            },
                            VolumeMounts: []apiv1.VolumeMount{
                                {Name: "models-dir", MountPath: "/models"},
                            },
                        },
                    },
                    Volumes: []apiv1.Volume{
                        {Name: "models-dir", VolumeSource: apiv1.VolumeSource{EmptyDir: &apiv1.EmptyDirVolumeSource{}}},
                    },
                },
            },
        },
    }  
}

func (k *kubernetesCluster) convertDeployment(deployment appsv1beta1.Deployment) *ModelDeployment {
    return &ModelDeployment{
        Type: deployment.Spec.Template.ObjectMeta.Labels["runner_type"],
        Replicas: *(deployment.Spec.Replicas),
        AvailableReplicas: deployment.Status.AvailableReplicas,
        ReadyReplicas: deployment.Status.ReadyReplicas,
    }
}
