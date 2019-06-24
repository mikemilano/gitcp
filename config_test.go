package main

import "testing"

func TestNewConfigSrcDstValues(t *testing.T) {
	// expect error for missing src
	ci := ConfigInput{
		dst:   "./",
		cdir: "/tmp",
		proto: "auto",
	}
	_, err := NewConfig(ci)
	if err == nil {
		t.Error("Expected empty src error, got none")
	}

	// expect error for missing dst
	ci = ConfigInput{
		src:   "./",
		cdir: "/tmp",
		proto: "auto",
	}
	_, err = NewConfig(ci)
	if err == nil {
		t.Error("Expected empty dst error, got none")
	}

	// expect no error for default paths
	ci = ConfigInput{
		src:   "./",
		dst:   "./",
		cdir: "/tmp",
		proto: "auto",
	}
	_, err = NewConfig(ci)
	if err != nil {
		t.Error("Expected no errors with default src/dst, got error")
	}
}

func TestNewConfigSrcDstLengths(t *testing.T) {
	// expect error for differing path counts where dst is not 1
	ci := ConfigInput{
		src:   "path/1,path/2",
		dst:   "path/1,path/2,path/3",
		cdir: "/tmp",
		proto: "auto",
	}
	_, err := NewConfig(ci)
	if err == nil {
		t.Error("Expected src/dst length error, got none")
	}

	// expect no error when both paths have an equal count
	ci = ConfigInput{
		src:   "./,./",
		dst:   "./,./",
		cdir: "/tmp",
		proto: "auto",
	}
	_, err = NewConfig(ci)
	if err != nil {
		t.Error("Expected src/dst length to pass with equal lengths, got error")
	}

	// expect error when destination does not exist
	ci = ConfigInput{
		src:   "./",
		dst:   "/invalid-dst",
		cdir: "/tmp",
		proto: "auto",
	}
	_, err = NewConfig(ci)
	if err == nil {
		t.Error("Expected error for invalid dst, got none")
	}

	// expect no error when src has multiple paths, and dst has 1
	ci = ConfigInput{
		src:   "./,./",
		dst:   "./",
		cdir: "/tmp",
		proto: "auto",
	}
	_, err = NewConfig(ci)
	if err != nil {
		t.Error("Expected multi src single dst to pass, got error")
	}
}

func TestCloneDir(t *testing.T) {
	// expect error if clone dir is empty
	ci := ConfigInput{
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
