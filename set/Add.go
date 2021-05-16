package set

func WriteToSet(hep map[string]int, nodes map[string]int, k string, slice []string) map[string]int {
	var exist bool

	for _, item := range slice {
		if k != item {
			_, exist = hep[k+" --> "+item]

			if exist {
				hep[k+" --> "+item] += 1
			} else {
				hep[k+" --> "+item] = 1
			}

			nodes["node "+k] = 1
			nodes["node "+item] = 1
		}
	}

	return hep
}
