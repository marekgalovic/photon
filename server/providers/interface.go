package providers

type Instance struct {
    Address string
    Port int
}

type InstanceProvider interface {
    Get(string) (*Instance, error)
}
