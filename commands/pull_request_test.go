package commands

import (
	"github.com/bmizerany/assert"
	"github.com/jingweno/gh/github"
	"testing"
)

func TestParsePullRequestProject(t *testing.T) {
	c := &github.Project{Host: "github.com", Owner: "jingweno", Name: "gh"}

	s := "develop"
	p, ref := parsePullRequestProject(c, s)
	assert.Equal(t, "develop", ref)
	assert.Equal(t, "github.com", p.Host)
	assert.Equal(t, "jingweno", p.Owner)
	assert.Equal(t, "gh", p.Name)

	s = "mojombo:develop"
	p, ref = parsePullRequestProject(c, s)
	assert.Equal(t, "develop", ref)
	assert.Equal(t, "github.com", p.Host)
	assert.Equal(t, "mojombo", p.Owner)
	assert.Equal(t, "gh", p.Name)

	s = "mojombo/jekyll:develop"
	p, ref = parsePullRequestProject(c, s)
	assert.Equal(t, "develop", ref)
	assert.Equal(t, "github.com", p.Host)
	assert.Equal(t, "mojombo", p.Owner)
	assert.Equal(t, "jekyll", p.Name)
}
