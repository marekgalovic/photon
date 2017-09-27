package files

import (
    "io";
    "io/ioutil";
    "golang.org/x/net/context";

    "cloud.google.com/go/storage"
)

type GoogleCloudStorageConfig struct {
    Bucket string
}

type GoogleCloudStorage struct {
    ctx context.Context
    client *storage.Client
    bucket *storage.BucketHandle
}

func NewGoogleCloudStorage(config GoogleCloudStorageConfig) (*GoogleCloudStorage, error) {
    ctx := context.Background()
    client, err := storage.NewClient(ctx)
    if err != nil {
        return nil, err
    }

    return &GoogleCloudStorage{
        ctx: ctx,
        client: client,
        bucket: client.Bucket(config.Bucket),
    }, nil
}

func (gcs *GoogleCloudStorage) Close() error {
    return gcs.client.Close()
}

func (gcs *GoogleCloudStorage) Reader(path string) (io.ReadCloser, error) {
    return gcs.bucket.Object(path).NewReader(gcs.ctx)
}

func (gcs *GoogleCloudStorage) Writer(path string) (io.WriteCloser, error) {
    return gcs.bucket.Object(path).NewWriter(gcs.ctx), nil
}

func (gcs *GoogleCloudStorage) Delete(path string) error {
    return gcs.bucket.Object(path).Delete(gcs.ctx)
}

func (gcs *GoogleCloudStorage) ReadBytes(path string) ([]byte, error) {
    reader, err := gcs.Reader(path)
    if err != nil {
        return nil, err
    }
    defer reader.Close()

    return ioutil.ReadAll(reader)
}

func (gcs *GoogleCloudStorage) WriteBytes(path string, data []byte) error {
    writer, _ := gcs.Writer(path)
    writer.Write(data)

    return writer.Close()
}
