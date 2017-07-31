package features

type FeaturesStore interface {
    Get(string, map[string]interface{}) (map[string]interface{}, error)
    Insert(string, string, []string, map[string]interface{}) error
    CreateFeatureSet(string, []string) error
    DeleteFeatureSet(string) error
}
