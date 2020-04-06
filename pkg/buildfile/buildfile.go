package buildfile

type Target interface {
	GetName() string
}

type BuildFile interface {
	GetTargets() []Target
	GetTarget(target string) Target
	RunTarget(target Target, args []string) (int, error)
}
