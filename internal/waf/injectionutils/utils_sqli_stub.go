//go:build !gcc

package injectionutils

// DetectSQLInjectionCache detect sql injection in string with cache
func DetectSQLInjectionCache(input string, isStrict bool, cacheLife int) bool {
	// stub
	return false
}

// DetectSQLInjection detect sql injection in string
func DetectSQLInjection(input string, isStrict bool) bool {
	// stub
	return false
}

func detectSQLInjectionOne(input string) bool {
	// stub
	return false
}