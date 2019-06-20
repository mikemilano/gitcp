package main

import (
	"testing"
)

func TestRepoFromString(t *testing.T) {
	githubTestOwner := "mikemilano"
	githubTestProject := "seeder"
	ownerSlashProject := githubTestOwner + "/" + githubTestProject

	// Test with GithubPath
	inputUrl := githubTestOwner + "/" + githubTestProject
	inputPath := "foo"
	expectedUrl := "ssh://git@github.com/" + ownerSlashProject

	repo, err := NewSeed(inputUrl)
	if err != nil {
		t.Error(err)
	}
	if repo == (Seed{}) {
		t.Error("Expected repo, got empty struct")
	}
	if repo.path != "" {
		t.Error("Expected empty path, got ", repo.path)
	}
	if repo.url.String() != expectedUrl {
		t.Error("Expected url to resolve to ssh://git@github.com/"+inputUrl+", got ", repo.url.String())
	}

	// Test with Path and clone URL format
	inputUrl = "git@github.com:" + githubTestOwner + "/" + githubTestProject + ".git"
	inputPath = "foo"
	expectedUrl = "ssh://git@github.com/" + ownerSlashProject + ".git"

	repo, err = NewSeed(inputUrl, inputPath)
	if err != nil {
		t.Error(err)
	}
	if repo.path != inputPath {
		t.Error("Expected path: "+inputPath+", got ", repo.path)
	}
	if repo.url.String() != expectedUrl {
		t.Error("Expected url to resolve to ssh://git@github.com/"+ownerSlashProject+".git, got ", repo.url.String())
	}
}
