package cluster

func ExpandComponent(strComponents []string) []ComponentList {
	var components []ComponentList
	for _, v := range strComponents {
		components = append(components, ComponentList{
			ComponentName: v,
		})
	}
	return components
}
