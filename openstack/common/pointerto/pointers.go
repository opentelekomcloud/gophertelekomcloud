package pointerto

// Int returns pointer to given int value.
func Int(src int) *int {
	return &src
}

// String returns pointer to given string value.
func String(src string) *string {
	return &src
}

// Bool returns pointer to given bool value.
func Bool(src bool) *bool {
	return &src
}
