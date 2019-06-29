package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	go_git_ssh "gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Seed struct {
	config Config
}

func NewSeed(ci ConfigInput) (Seed, error) {

	config, err := NewConfig(ci)
	if err != nil {
		return Seed{}, err
	}

	return Seed{config: config}, nil
}

func (s *Seed) clone() error {
	// remove directory if exists
	path := filepath.Join(s.config.cdir, "seeder")
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	// create directory
	err = os.Mkdir(path, os.ModePerm)
	if err != nil {
		return err
	}

	opts := &git.CloneOptions{
		URL: s.config.url.String(),
	}

	if s.config.proto == "ssh" && s.config.key != "" {
		opts.Auth = get_ssh_key_auth(s.config.key)
	}
	// TODO: Fix so it works
	if s.config.quiet == false {
		opts.Progress = os.Stdout
	}

	_, err = git.PlainClone(path, false, opts)

	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func get_ssh_key_auth(privateSshKeyFile string) transport.AuthMethod {
	var auth transport.AuthMethod
	sshKey, _ := ioutil.ReadFile(privateSshKeyFile)
	signer, _ := ssh.ParsePrivateKey([]byte(sshKey))
	auth = &go_git_ssh.PublicKeys{User: "git", Signer: signer}
	return auth
}
