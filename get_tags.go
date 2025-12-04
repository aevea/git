package git

import (
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/apex/log"
)

func (g *Git) getTags() ([]*Tag, error) {
	var tags []*Tag

	// Get all tags with for-each-ref using a unique delimiter
	// Use <TAG_SEP> as delimiter between tags to avoid conflicts with field delimiters
	refsOutput, err := g.runGitCommand("for-each-ref", "--format=%(refname)|%(objectname)|%(objecttype)|%(taggerdate:unix)<TAG_SEP>", "refs/tags/")
	if err != nil {
		return nil, err
	}

	// Split by tag delimiter
	tagBlocks := strings.Split(refsOutput, "<TAG_SEP>")

	for _, block := range tagBlocks {
		block = strings.TrimSpace(block)
		if block == "" {
			continue
		}

		fields := strings.Split(block, "|")
		if len(fields) < 3 {
			continue
		}

		refName := strings.TrimSpace(fields[0])
		if !strings.HasPrefix(refName, "refs/tags/") {
			continue
		}

		objectHash := strings.TrimSpace(fields[1])
		if objectHash == "" {
			continue
		}

		objectType := strings.TrimSpace(fields[2])
		if objectType == "" {
			continue
		}

		dateStr := ""
		if len(fields) > 3 {
			dateStr = strings.TrimSpace(fields[3])
		}

		var hash Hash
		var commitDate time.Time
		var tagName string

		if objectType == "tag" {
			// Annotated tag - resolve to commit
			targetHashStr, err := g.runGitCommand("rev-parse", objectHash+"^{commit}")
			if err != nil {
				log.Debugf("Failed to resolve annotated tag target %s: %v", refName, err)
				continue
			}

			hash, err = NewHash(targetHashStr)
			if err != nil {
				log.Debugf("Failed to parse target hash for annotated tag %s: %v", refName, err)
				continue
			}

			// Use tagger date if available, otherwise use commit date
			if dateStr != "" {
				timestamp, err := strconv.ParseInt(dateStr, 10, 64)
				if err == nil {
					commitDate = time.Unix(timestamp, 0)
				} else {
					commitDate, err = g.commitDate(hash)
					if err != nil {
						log.Debugf("tag: %v ignored due to missing date: %v", refName, err)
						continue
					}
				}
			} else {
				commitDate, err = g.commitDate(hash)
				if err != nil {
					log.Debugf("tag: %v ignored due to missing commit date: %v", refName, err)
					continue
				}
			}

			// For annotated tags, use just the tag name without refs/tags/ prefix
			tagName = strings.TrimPrefix(refName, "refs/tags/")
		} else {
			// Lightweight tag - objectHash points directly to commit
			var err error
			hash, err = NewHash(objectHash)
			if err != nil {
				log.Debugf("Failed to parse hash for tag %s: %v", refName, err)
				continue
			}

			commitDate, err = g.commitDate(hash)
			if err != nil {
				log.Debugf("tag: %v ignored due to missing commit date: %v", refName, err)
				continue
			}

			// For lightweight tags, use full ref name
			tagName = refName
		}

		tags = append(tags, &Tag{Name: tagName, Date: commitDate, Hash: hash})
	}

	// Tags are alphabetically ordered. We need to sort them by date.
	sortedTags := sortTags(tags)

	log.Debug("Sorted tag output: ")
	for _, taginstance := range sortedTags {
		log.Debugf("hash: %v time: %v", taginstance.Hash, taginstance.Date.UTC())
	}

	return sortedTags, nil
}

// sortTags sorts the tags according to when their parent commit happened.
func sortTags(tags []*Tag) []*Tag {
	sort.Slice(tags, func(i, j int) bool {
		return tags[i].Date.After(tags[j].Date)
	})

	return tags
}
