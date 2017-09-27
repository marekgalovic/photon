package runner

import (
    "os";
    "fmt";
    "io/ioutil";
    "path/filepath";

    "github.com/marekgalovic/photon/go/core/repositories";

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

    manager := &ModelManager{
        config: config,
        instancesRepository: instancesRepository,
        watcher: watcher,
        zookeeperUids: make(map[string]string),
        stopper: make(chan struct{}),
    }

    if err := manager.load(); err != nil {
        return nil, err
    }

    if err := manager.watch(); err != nil {
        return nil, err
    }

    return manager, nil
}

func (m *ModelManager) Close() {
    m.watcher.Close()
    close(m.stopper)
}

func (m *ModelManager) Get(versionUid string) {

}

func (m *ModelManager) load() error {
    info, err := os.Stat(m.config.ModelsDir); 
    if err != nil {
        return err
    }

    if !info.IsDir() {
        return fmt.Errorf("%s is not a directory.", m.config.ModelsDir)
    }

    files, err := ioutil.ReadDir(m.config.ModelsDir)
    if err != nil {
        return err
    }

    for _, file := range files {
        if err = m.create(file.Name()); err != nil {
            return err
        }
        log.Infof("Loaded model: %s", file.Name())
    }

    return nil
}

func (m *ModelManager) watch() error {
    if err := m.watcher.Add(m.config.ModelsDir); err != nil {
        return err
    }

    go func() {
        for {
            select {
            case <- m.stopper:
                return
            case event := <- m.watcher.Events:
                switch {
                case event.Op & fsnotify.Create == fsnotify.Create:
                    if err := m.create(event.Name); err != nil {
                        log.Errorf("Failed to create model. %v", err)
                    }
                    log.Infof("Created model: %s", event.Name)
                case event.Op & fsnotify.Remove == fsnotify.Remove:
                    if err := m.remove(event.Name); err != nil {
                        log.Errorf("Failed to remove model. %v", err)
                    }
                    log.Infof("Removed model: %s", event.Name)
                }
            case err := <- m.watcher.Errors:
                log.Error(err)
            }
        }
    }()

    return nil
}

func (m *ModelManager) create(path string) error {
    fileName := filepath.Base(path)
    uid, err := m.instancesRepository.Register(fileName, &repositories.Instance{Address: m.config.Address, Port: m.config.Port})
    if err != nil {
        return err
    }

    m.zookeeperUids[fileName] = uid
    return nil
}

func (m *ModelManager) remove(path string) error {
    fileName := filepath.Base(path)
    uid, exists := m.zookeeperUids[fileName]
    if !exists {
        return fmt.Errorf("Unknown model uid: %s", fileName)
    }

    return m.instancesRepository.Unregister(fileName, uid)
}
