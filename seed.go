package main

import (
	"fmt"
	"github.com/whilp/git-urls"
	"golang.org/x/crypto/ssh"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	go_git_ssh "gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	"io/ioutil"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
)

type Seed struct {
	config Config
	url    *url.URL
}

func NewSeed(ci ConfigInput, target string) (Seed, error) {

	// TODO: Handle auto
	if ci.proto == "auto" {
		ci.proto = "ssh"
	}

	config, err := NewConfig(ci)
	if err != nil {
		return Seed{}, err
	}

	gitUrl, err := giturls.Parse(target)
	if err != nil {
		return Seed{}, err
	}

	filePath := regexp.MustCompile(`^file://@[a-z0-9-]{0,38}/`)

	if filePath.MatchString(gitUrl.String()) {
		// TODO: Refactor the way url is set
		url := "https://github.com/" + target + ".git"
		if config.proto == "ssh" {
			url = "git@github.com:" + target
		}
		gitUrl, _ = giturls.Parse(url)

		fmt.Println(gitUrl)
		return Seed{config: config, url: gitUrl}, nil
	}

	return Seed{config: config, url: gitUrl}, nil
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
		URL: s.url.String(),
	}

	// TODO: Refactor to use argument for private key
	if s.config.proto == "ssh" {
		usr, _ := user.Current()
		opts.Auth = get_ssh_key_auth(usr.HomeDir + "/.ssh/id_rsa")
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
