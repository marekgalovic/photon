package runner

import (
    "fmt";
    "path/filepath";

    "github.com/marekgalovic/photon/go/core/storage/repositories";

    "github.com/fsnotify/fsnotify";
    log "github.com/Sirupsen/logrus"
)

type ModelManager struct {
    config *Config
    instancesRepository *repositories.InstancesRepository
    watcher *fsnotify.Watcher
    zookeeperUids map[string]string
    stopper chan struct{}
}

func NewModelManager(config *Config, instancesRepository *repositories.InstancesRepository) (*ModelManager, error) {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        return nil, err
    }

    return &ModelManager{
        config: config,
        instancesRepository: instancesRepository,
        watcher: watcher,
        zookeeperUids: make(map[string]string),
        stopper: make(chan struct{}),
    }, nil
}

func (m *ModelManager) Close() {
    m.watcher.Close()
    close(m.stopper)
}

func (m *ModelManager) Get(versionUid string) {

}

func (m *ModelManager) Watch() error {
    if err := m.watcher.Add(m.config.ModelsDir); err != nil {
        return err
    }

    for {
        select {
        case <- m.stopper:
            return nil
        case event := <- m.watcher.Events:
            fileName := filepath.Base(event.Name)

            switch {
            case event.Op & fsnotify.Create == fsnotify.Create:
                if err := m.create(fileName); err != nil {
                    log.Errorf("Failed to create model. %v", err)
                }
                log.Infof("Created model: %s", fileName)
            case event.Op & fsnotify.Remove == fsnotify.Remove:
                if err := m.remove(fileName); err != nil {
                    log.Errorf("Failed to remove model. %v", err)
                }
                log.Infof("Removed model: %s", fileName)
            }
        case err := <- m.watcher.Errors:
            log.Error(err)
        }
    }
}

func (m *ModelManager) create(fileName string) error {
    uid, err := m.instancesRepository.Register(fileName, m.config.Address, m.config.Port)
    if err != nil {
        return err
    }

    m.zookeeperUids[fileName] = uid
    return nil
}

func (m *ModelManager) remove(fileName string) error {
    uid, exists := m.zookeeperUids[fileName]
    if !exists {
        return fmt.Errorf("Unknown model uid: %s", fileName)
    }

    return m.instancesRepository.Unregister(fileName, uid)
}
