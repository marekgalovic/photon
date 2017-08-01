package storage

type Scannable interface {
    Scan(...interface{}) error
}

type Countable interface {
    Count(string) (int, error)
}

type FeaturesStore interface {
    Get(string, []string, map[string]interface{}) (map[string]interface{}, error)
    Insert(string, string, []string, map[string]interface{}) error
    CreateFeatureSet(string, []string) error
    DeleteFeatureSet(string) error
}
