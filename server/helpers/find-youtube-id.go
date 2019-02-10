package helpers

import (
	"regexp"
	"strings"
)

// FindVideoID extracts youtube video id from a link
func FindVideoID(url string) string {
	videoID := url
	if strings.Contains(videoID, "youtu") || strings.ContainsAny(videoID, "\"?&/<%=") {
		reList := []*regexp.Regexp{
			regexp.MustCompile(`(?:v|embed|watch\?v)(?:=|/)([^"&?/=%]{11})`),
			regexp.MustCompile(`(?:=|/)([^"&?/=%]{11})`),
			regexp.MustCompile(`([^"&?/=%]{11})`),
		}
		for _, re := range reList {
			if isMatch := re.MatchString(videoID); isMatch {
				subs := re.FindStringSubmatch(videoID)
				videoID = subs[1]
			}
		}
	}
	if strings.ContainsAny(videoID, "?&/<%=") {
		return ""
	}
	if len(videoID) < 10 {
		return ""
	}
	return videoID
}
