package features

type FeaturesStore interface {
    Get(string, []string) (map[string]interface{}, error)
    CreateFeaturesSet(string) error
    DeleteFeaturesSet(string) error
}
