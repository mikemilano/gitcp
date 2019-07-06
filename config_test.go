package main

import (
	"testing"
)

func TestNewConfigSrcDstValues(t *testing.T) {
	// expect no error for default paths
	ci := ConfigInput{
		target: "https://github.com/mikemilano/seeder.git",
		cdir: "/tmp",
		proto: "auto",
	}
	_, err := NewConfig(ci)
	if err != nil {
		t.Error("Expected no errors with empty src/dst, got error")
	}

	// expect no error for default paths
	ci = ConfigInput{
		target: "https://github.com/mikemilano/seeder.git",
		src:   "README.md",
		dst:   ".",
		cdir: "/tmp",
		proto: "auto",
	}
	_, err = NewConfig(ci)
	if err != nil {
		t.Error("Expected no errors with default src/dst, got error")
	}
}

func TestTarget(t *testing.T) {
	// expect error if target is empty
	ci := ConfigInput{
		src: "./",
		dst: "./",
		cdir: "/tmp",
		proto: "auto",
	}
	_, err := NewConfig(ci)
	if err == nil {
		t.Error("Expected error from empty target, got none")
	}

	// convert short form with https
	ci = ConfigInput{
		target: "mikemilano/seeder",
		src: "./",
		dst: "./",
		cdir: "/tmp",
		proto: "https",
	}
	config, err := NewConfig(ci)

	if config.url.String() != "https://@github.com/" + ci.target + ".git" {
		t.Error("Expected short form target with https to get https://@github.com/" + ci.target + ".git, got", config.url.String())
	}

	// convert short form with ssh
	ci = ConfigInput{
		target: "mikemilano/seeder",
		src: "./",
		dst: "./",
		cdir: "/tmp",
		proto: "ssh",
	}
	config, err = NewConfig(ci)

	if config.url.String() != "ssh://git@github.com/" + ci.target {
		t.Error("Expected short form target with ssh url to be ssh://git@github.com/" + ci.target + ", got", config.url.String())
	}

	// convert short form with auto
	ci = ConfigInput{
		target: "mikemilano/seeder",
		src: "./",
		dst: "./",
		cdir: "/tmp",
		proto: "auto",
	}
	config, err = NewConfig(ci)

	if config.url.String() != "ssh://git@github.com/" + ci.target {
		t.Error("Expected short form target with auto url to be ssh://git@github.com/" + ci.target + ", got", config.url.String())
	}
}

func TestCloneDir(t *testing.T) {
	// expect error if clone dir is empty
	ci := ConfigInput{
		target: "https://github.com/mikemilano/seeder.git",
		src: "./",
		dst: "./",
		proto: "auto",
	}
	_, err := NewConfig(ci)
	if err == nil {
		t.Error("Expected error from empty clone dir, got none")
	}

	// expect error if clone dir does not exist
	ci = ConfigInput{
		target: "https://github.com/mikemilano/seeder.git",
		src: "./",
		dst: "./",
		cdir: "/invalid-directory",
		proto: "auto",
	}
	_, err = NewConfig(ci)
	if err == nil {
		t.Error("Expected error from invalid clone dir, got none")
	}

	// expect no error if clone dir exists
	ci = ConfigInput{
		target: "https://github.com/mikemilano/seeder.git",
		src: "./",
		dst: "./",
		cdir: "/tmp",
		proto: "auto",
	}
	_, err = NewConfig(ci)
	if err != nil {
		t.Error("Expected no error from valid clone dir, got error")
	}
}

func TestNewConfigProtocol(t *testing.T) {
	// expect error with empty proto
	ci := ConfigInput{
		target: "mikemilano/seeder",
		src:   "./",
		dst:   "./",
		cdir: "/tmp",
		proto: "",
	}
	_, err := NewConfig(ci)
	if err == nil {
		t.Error("Expected error from empty proto, got none")
	}

	// expect error invalid proto
	ci = ConfigInput{
		target: "mikemilano/seeder",
		src:   "./",
		dst:   "./",
		cdir: "/tmp",
		proto: "foobar",
	}
	_, err = NewConfig(ci)
	if err == nil {
		t.Error("Expected error from invalid proto, got none")
	}

	// expect no error with auto proto
	ci = ConfigInput{
		target: "mikemilano/seeder",
		src:   "./",
		dst:   "./",
		cdir: "/tmp",
		proto: "auto",
	}
	_, err = NewConfig(ci)
	if err != nil {
		t.Error("Expected no error with auto proto, got error")
	}

	// expect no error with https proto
	ci = ConfigInput{
		target: "mikemilano/seeder",
		src:   "./",
		dst:   "./",
		cdir: "/tmp",
		proto: "https",
	}
	_, err = NewConfig(ci)
	if err != nil {
		t.Error("Expected no error with https proto, got error")
	}

	// expect no error with ssh proto
	ci = ConfigInput{
		target: "mikemilano/seeder",
		src:   "./",
		dst:   "./",
		cdir: "/tmp",
		proto: "ssh",
	}
	_, err = NewConfig(ci)
	if err != nil {
		t.Error("Expected no error with ssh proto, got error")
	}
}
