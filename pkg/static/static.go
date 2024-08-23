package static

// Static is expansions manually copied for upstream methods

func GetListersExpansions(target string) string {
	return ListersExpansions[target]
}

var ListersExpansions = map[string]string{
	"ReplicationControllerLister": StaticReplicationControllerListerExpansion,
}
