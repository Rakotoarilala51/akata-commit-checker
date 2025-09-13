package subcommand

var (
	globalThreshold *int
	globalVerbose   *bool
)

func SetGlobalConfig(threshold *int, verbose *bool) {
	globalThreshold = threshold
	globalVerbose = verbose
}

func GetConfig() (int, bool) {
	if globalThreshold == nil || globalVerbose == nil {
		return 3, false 
	}
	return *globalThreshold, *globalVerbose
}