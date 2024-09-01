package static

// Static is expansions manually copied for upstream methods

func GetListersExpansions(target string) string {
	return ListersExpansions[target]
}

func GetClientSetFakeExpansions(target string) string {
	return ClientSetFakeExpansions[target]
}

var ListersExpansions = map[string]string{
	"ReplicationControllerLister": StaticReplicationControllerListerExpansion,
}

var ClientSetFakeExpansions = map[string]string{
	"StaticFakesClientSetEventNamespacedExpansion": StaticFakesClientSetEventNamespacedExpansion,
	"StaticFakesClientSetEventExpansion":           StaticFakesClientSetEventExpansion,
	"StaticFakesClientSetNamespacesExpansion":      StaticFakesClientSetNamespacesExpansion,
	"StaticFakesClientSetNodesExpansion":           StaticFakesClientSetNodesExpansion,
	"StaticFakesClientSetPodsExpansion":            StaticFakesClientSetPodsExpansion,
	"StaticFakesClientSetServicesExpansion":        StaticFakesClientSetServicesExpansion,
	"StaticFakesClientSetServiceAccountsExpansion": StaticFakesClientSetServiceAccountsExpansion,
}
