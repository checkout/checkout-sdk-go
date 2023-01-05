package nas

type DefaultSdkBuilder interface {
	Build() (*Api, error)
}
