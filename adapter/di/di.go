package di

var Container dject.Container

func CreateContainer() {
	container := dject.NewContainer()
	Container = container
}
