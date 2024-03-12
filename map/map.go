package mapp

// mergeMaps is a generic function that merges two maps. The maps can have any type for keys and values.
// It clones the first map and then adds/overwrites its entries with those from the second map.
func Merge[K comparable, V any](originalMap, newProperties map[K]V) map[K]V {
	// Clone the original map
	clonedMap := make(map[K]V)
	for key, value := range originalMap {
		clonedMap[key] = value
	}
	// Merge in new properties
	for key, value := range newProperties {
		clonedMap[key] = value
	}
	return clonedMap
}
