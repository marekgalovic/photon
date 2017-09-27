package features

type FeaturesStore interface {
    Get(int64, []string, map[string]interface{}) (map[string]interface{}, error)
    Insert(int64, []string, map[string]interface{}) error
    CreateFeatureSet(int64, []string) error
    DeleteFeatureSet(int64) error
}
