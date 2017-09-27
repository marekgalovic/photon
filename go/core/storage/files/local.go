package files

import (
    "os";
    "io";
    "io/ioutil";
    "fmt";
    "path/filepath";
)

type LocalStorageConfig struct {
    Dir string
}

type LocalStorage struct {
    Dir string
}

func NewLocalStorage(config LocalStorageConfig) (*LocalStorage, error) {
    stat, err := os.Stat(config.Dir)
    if err != nil {
        return nil, err
    }
    if !stat.IsDir() {
        return nil, fmt.Errorf("Path %s is not a directory.", config.Dir)
    }

    return &LocalStorage{Dir: config.Dir}, nil
}

func (ls *LocalStorage) Close() error {
    return nil
}

func (ls *LocalStorage) Reader(path string) (io.ReadCloser, error) {
    return os.Open(filepath.Join(ls.Dir, path))
}

func (ls *LocalStorage) Writer(path string) (io.WriteCloser, error) {
    return os.Create(filepath.Join(ls.Dir, path))
}

func (ls *LocalStorage) Delete(path string) error {
    return os.Remove(filepath.Join(ls.Dir, path))
}

func (ls *LocalStorage) ReadBytes(path string) ([]byte, error) {
    reader, err := ls.Reader(path)
    if err != nil {
        return nil, err
    }
    defer reader.Close()

    return ioutil.ReadAll(reader)
}

func (ls *LocalStorage) WriteBytes(path string, data []byte) error {
    writer, err := ls.Writer(path)
    if err != nil {
        return err
    }
    defer writer.Close()

    _, err = writer.Write(data)
    return err
}

