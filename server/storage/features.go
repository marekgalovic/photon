package storage

type FeaturesRepository struct {
    db *Mysql
}

func NewFeaturesRepository(db *Mysql) *FeaturesRepository {
    return &FeaturesRepository{db: db}
}


