package storage

type Scannable interface {
    Scan(...interface{}) error
}

type Countable interface {
    Count(string) (int, error)
}
