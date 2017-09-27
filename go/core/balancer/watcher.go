package balancer

import (
    "github.com/marekgalovic/photon/go/core/repositories";

    "google.golang.org/grpc/naming";
)

type Watcher struct {
    instancesRepository repositories.InstancesRepository
    modelVersionUid string
    instances map[string]*repositories.Instance
    updates chan map[string]*repositories.Instance
    errors chan error
    stopper chan struct{}
}

func NewWatcher(instancesRepository repositories.InstancesRepository, modelVersionUid string) *Watcher {
    watcher := &Watcher {
        instancesRepository: instancesRepository,
        modelVersionUid: modelVersionUid,
        instances: make(map[string]*repositories.Instance),
        updates: make(chan map[string]*repositories.Instance),
        errors: make(chan error),
        stopper: make(chan struct{}),
    }

    go watcher.watchUpdates()
    return watcher
}

func (w *Watcher) Close() {
    close(w.stopper)
}

func (w *Watcher) Next() ([]*naming.Update, error) {
    select {
    case <- w.stopper:
        return nil, nil
    case err := <- w.errors:
        return nil, err
    case instances := <- w.updates:
        return w.generateUpdates(instances), nil
    }
}

func (w *Watcher) watchUpdates() {
    for {
        instances, err := w.instancesRepository.List(w.modelVersionUid)
        if err != nil {
            w.errors <- err
            return
        }
        w.updates <- instances

        updateEvent, err := w.instancesRepository.Watch(w.modelVersionUid)
        if err != nil {
            w.errors <- err
            return
        }

        select {
        case <- w.stopper:
            return
        case <- updateEvent:
            continue
        }
    }
}

func (w *Watcher) generateUpdates(updatedInstances map[string]*repositories.Instance) []*naming.Update {
    updates := make([]*naming.Update, 0)

    for uid, instance := range w.instances {
        if _, exists := updatedInstances[uid]; !exists {
            updates = append(updates, &naming.Update{Op: naming.Delete, Addr: instance.FullAddress()})
            delete(w.instances, uid)
        }
    }

    for uid, instance := range updatedInstances {
        if _, exists := w.instances[uid]; !exists {
            updates = append(updates, &naming.Update{Op: naming.Add, Addr: instance.FullAddress()})
            w.instances[uid] = instance
        }
    }

    return updates
}
