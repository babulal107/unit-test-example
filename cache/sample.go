package cache

type Sample interface {
	DoSomething(int, string) (string, error)
}
