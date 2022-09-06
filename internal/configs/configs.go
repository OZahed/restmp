package configs

var (
	appname string
	build   string
	commit  string
	branch  string
	tag     string
)

func AppName() string {
	return appname
}

func Build() string {
	return build
}

func Commit() string {
	return commit
}

func Branch() string {
	return branch
}

func Tag() string {
	return tag
}
