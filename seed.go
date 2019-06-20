package main

import (
	"errors"
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
	"regexp"
)

type Seed struct {
	url  *url.URL
	path string
}

func NewSeed(parts ...string) (Seed, error) {
	if len(parts) == 0 {
		return Seed{}, errors.New("seed: seed created with no input")
	}
	str := parts[0]

	path := ""
	if len(parts) > 1 {
		path = parts[1]
	}

	gitUrl, err := giturls.Parse(str)
	if err != nil {
		return Seed{}, err
	}

	filePath := regexp.MustCompile(`^file://@[a-z0-9-]{0,38}/`)

	if filePath.MatchString(gitUrl.String()) {
		// TODO: Determine when to set http or ssh format
		gitUrl, _ = giturls.Parse("git@github.com:" + str)
		return Seed{url: gitUrl, path: path}, nil
	}

	return Seed{url: gitUrl, path: path}, nil
}

func (s *Seed) clone() {
	//Info("git clone https://github.com/src-d/go-git")

	usr, err := user.Current()

	// TODO: Optionally add auth if endpoint is ssh

	_, err = git.PlainClone("/tmp/foo", false, &git.CloneOptions{
		URL:      s.url.String(),
		Auth: get_ssh_key_auth(usr.HomeDir + "/.ssh/id_rsa"),
		Progress: os.Stdout,
	})

	if err != nil {
		fmt.Println(err)
	}

	//CheckIfError(err)
}

func get_ssh_key_auth(privateSshKeyFile string) transport.AuthMethod {
	var auth transport.AuthMethod
	sshKey, _ := ioutil.ReadFile(privateSshKeyFile)
	signer, _ := ssh.ParsePrivateKey([]byte(sshKey))
	auth = &go_git_ssh.PublicKeys{User: "git", Signer: signer}
	return auth
}
