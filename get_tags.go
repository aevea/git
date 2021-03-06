package git

import (
	"sort"

	"github.com/apex/log"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func (g *Git) getTags() ([]*Tag, error) {
	tagrefs, err := g.repo.Tags()

	if err != nil {
		return nil, err
	}

	defer tagrefs.Close()

	var tags []*Tag

	err = tagrefs.ForEach(func(t *plumbing.Reference) error {
		commitDate, err := g.commitDate(t.Hash())

		if err != nil {
			log.Debugf("tag: %v ignored due to missing commit date: %v", t.Name(), err)
			return nil
		}

		tags = append(tags, &Tag{Name: t.Name().String(), Date: commitDate, Hash: t.Hash()})
		return nil
	})

	if err != nil {
		return nil, err
	}

	tagObjects, err := g.repo.TagObjects()

	if err != nil {
		return nil, err
	}

	err = tagObjects.ForEach(func(tag *object.Tag) error {
		tags = append(tags, &Tag{Name: tag.Name, Date: tag.Tagger.When, Hash: tag.Target})

		return nil
	})

	// Tags are alphabetically ordered. We need to sort them by date.
	sortedTags := sortTags(g.repo, tags)

	log.Debug("Sorted tag output: ")
	for _, taginstance := range sortedTags {
		log.Debugf("hash: %v time: %v", taginstance.Hash, taginstance.Date.UTC())
	}

	return sortedTags, nil
}

// sortTags sorts the tags according to when their parent commit happened.
func sortTags(repo *git.Repository, tags []*Tag) []*Tag {
	sort.Slice(tags, func(i, j int) bool {
		return tags[i].Date.After(tags[j].Date)
	})

	return tags
}
