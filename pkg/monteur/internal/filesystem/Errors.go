package filesystem

//nolint:stylecheck,revive,lll
const (
	ERROR_BAD_REPO             = "current directory is not a git repo with monteur"
	ERROR_CURRENT_DIRECTORY    = "failed to obtain current directory path"
	ERROR_MISSING_DIR          = "missing directory"
	ERROR_FAILED_CONFIG        = "unable to open workspace.toml file"
	ERROR_FAILED_CONFIG_DECODE = "failed to decode workspace.toml file"
)
