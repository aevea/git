package git

import (
	"strconv"
	"time"
)

// commitDate gets the commit at hash and returns the time of the commit
func (g *Git) commitDate(commit Hash) (time.Time, error) {
	hashStr := commit.String()

	dateStr, err := g.runGitCommand("log", "-1", "--format=%ct", hashStr)
	if err != nil {
		return time.Now(), err
	}

	timestamp, err := strconv.ParseInt(dateStr, 10, 64)
	if err != nil {
		return time.Now(), err
	}

	return time.Unix(timestamp, 0), nil
}
