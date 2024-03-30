package dateutils

// SplitYmd 分隔Ymd格式的日期
// Ymd => Y-m-d
func SplitYmd(day string) string {
	if len(day) != 8 {
		return day
	}
	return day[:4] + "-" + day[4:6] + "-" + day[6:]
}
