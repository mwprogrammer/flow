package utilities

func Contains(s string, list []string) bool {
	for _, item := range list {
		if item == s {
			return true
		}
	}
	return false
}

func IsTypeString(i any) bool {

	_, ok := i.(string)

	if ok {
		return true
	} else {
		return false
	}

}
