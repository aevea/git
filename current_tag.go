package git

import (
	"errors"

	"github.com/apex/log"
)

var (
	// ErrCommitNotOnTag is returned by CurrentTag if the commit is not on a tag
	ErrCommitNotOnTag = errors.New("current commit is not on a tag")
)

// CurrentTag returns a Tag if the current HEAD is on a tag
func (g *Git) CurrentTag() (*Tag, error) {
	head, err := g.repo.Head()

	if err != nil {
		return nil, err
	}

	tags, err := g.getTags()

	if err != nil {
		return nil, err
	}

	log.Debugf("head hash: %s", head.Hash())

	for _, tag := range tags {
		log.Debugf("tag: %v, hash: %v", tag.Name, tag.Hash)

		if tag.Hash == head.Hash() {
			return tag, nil
		}
	}

	return nil, ErrCommitNotOnTag
}
