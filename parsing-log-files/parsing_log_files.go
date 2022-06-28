package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[TRC\]|^\[DBG\]|^\[INF\]|^\[ERR\]|^\[WRN\]|^\[FTL\]`)

	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<(~|\*|=|-|)*>`)

	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)"(?:...)*password(?:...)*"`)
	count := 0

	for _, l := range lines {
		if re.MatchString(l) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line(\d)*`)

	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`(?:...)*(User)(\s)+(\w*)`)

	for i, line := range lines {
		matches := re.FindStringSubmatch(line)
		for matchGroup, match := range matches {
			if matchGroup == 3 {
				lines[i] = "[USR] " + match + " " + lines[i]
			}
		}
	}

	return lines
}
