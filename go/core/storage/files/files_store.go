package files

import (
    "io";
)

type FilesStore interface {
    Close() error
    Reader(string) (io.ReadCloser, error)
    Writer(string) (io.WriteCloser, error)
    Delete(string) error
    ReadBytes(string) ([]byte, error)
    WriteBytes(string, []byte) error
}
