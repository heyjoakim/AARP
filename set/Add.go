package set

func WriteToSet(dependencies map[string]bool, nodes map[string]bool, k string, slice []string) (map[string]bool, map[string]bool) {
	for _, s := range slice {
		if k != s {
			dependencies[k+" --> "+s] = true
		}
		nodes["node "+k] = true
		nodes["node "+s] = true
	}
	return dependencies, nodes
}
