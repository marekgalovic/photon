package deployer

import (
    "fmt";
    "time";
    "io";
    // "io/ioutil";

    "github.com/marekgalovic/photon/go/core/storage";
    "github.com/marekgalovic/photon/go/core/storage/files";
    "github.com/marekgalovic/photon/go/core/repositories";

    log "github.com/Sirupsen/logrus"
)

type Deployer struct {
    config *Config
    deployerRepository repositories.DeployerRepository
    modelsStore files.FilesStore
    modelsDir files.FilesStore
    logger *log.Entry
    versions map[int64]string
    stopper chan struct{}
}

func NewDeployer(config *Config, zookeeper *storage.Zookeeper, modelsStore files.FilesStore) (*Deployer, error) {
    modelsDir, err := files.NewLocalStorage(files.LocalStorageConfig{Dir: config.ModelsDir})
    if err != nil {
        return nil, err
    }

    return &Deployer{
        config: config,
        deployerRepository: repositories.NewDeployerRepository(zookeeper),
        modelsStore: modelsStore,
        modelsDir: modelsDir,
        logger: log.WithFields(log.Fields{"model_id": config.ModelId}),
        versions: make(map[int64]string),
        stopper: make(chan struct{}),
    }, nil
}

func (d *Deployer) Close() {
    close(d.stopper)
}

func (d *Deployer) Run() error {
    if err := d.waitForModel(); err != nil {
        return err
    }
    d.logger.Info("Watching model versions.")

    for {
        versions, event, err := d.deployerRepository.ListVersionsW(d.config.ModelId)
        if err != nil {
            return err
        }
        d.updateVersions(versions)

        select {
        case <- event:
            continue
        case <- d.stopper:
            return nil
        }
    }
    return nil
}

func (d *Deployer) waitForModel() error {
    exists, event, err := d.deployerRepository.ModelExistsW(d.config.ModelId)
    if err != nil {
        return err
    }
    if exists {
        return nil
    }

    d.logger.Info("Waiting for model.")
    select {
    case <- event:
        return nil
    case <- time.After(30 * time.Second):
        return fmt.Errorf("Timeout while waiting for model to be available.")
    }
}

func (d *Deployer) updateVersions(versions map[int64]string) {
    fetched := make(map[int64]struct{})
    for versionId, sourcePath := range versions {
        fetched[versionId] = struct{}{}
        if _, exists := d.versions[versionId]; !exists {
            d.versions[versionId] = sourcePath
            go d.deployVersion(versionId, sourcePath)
        }
    }

    for versionId, fileName := range d.versions {
        if _, exists := fetched[versionId]; !exists {
            delete(d.versions, versionId)
            go d.undeployVersion(versionId, fileName)
        }
    }
}

func (d *Deployer) deployVersion(id int64, fileName string) {
    reader, err := d.modelsStore.Reader(fileName)
    if err != nil {
        log.Errorf("Failed to deploy model version %d. %v", id, err)
        return
    }
    defer reader.Close()

    writer, err := d.modelsDir.Writer(fileName)
    if err != nil {
        log.Errorf("Failed to deploy model version %d. %v", id, err)
        return
    }
    defer writer.Close()

    _, err = io.Copy(writer, reader); 
    if err != nil {
        log.Errorf("Failed to deploy model version %d. %v", id, err)
        return
    }

    log.Infof("Deployed version: %d, filename: %s", id, fileName)
}

func (d *Deployer) undeployVersion(id int64, fileName string) {
    if err := d.modelsDir.Delete(fileName); err != nil {
        log.Error("Failed to undeploy version %d. %v", id, err)
        return
    }
    log.Infof("Undeployed version: %d", id)
}

