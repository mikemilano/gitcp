package main

import (
	"fmt"
	"github.com/otiai10/copy"
	"golang.org/x/crypto/ssh"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	go_git_ssh "gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
	//"io"
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
	// TODO: support branches
	// TODO: pull & branch (no clone) if the cached repo is the same

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
		opts.Auth = GetSSHKeyAuth(s.config.key)
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

func (s Seed) process() error {
	// absolute (Abs) paths to source and dst direct
	srcDirAbs := filepath.Join(s.config.cdir, "seeder")
	dstDirAbs, _ := os.Getwd()

	for i, p := range s.config.src {

		srcGlob := filepath.Join(srcDirAbs, p)
		matches, _ := filepath.Glob(srcGlob)

		// TODO: Error that source did not match any files?
		if len(matches) == 0 {
			continue
		}

		fmt.Printf("%+v\n", matches)

		for _, matchSrcAbs := range matches {
			// stat info of the match
			matchSrcInfo, _ := os.Stat(matchSrcAbs)
			// relative path of the match from the source dir (srcDirAbs)
			matchSrcRel, _ := filepath.Rel(srcDirAbs, matchSrcAbs) //strings.ReplaceAll(matchSrcAbs, srcDirAbs, "")
			// last element (file or dir) of the match
			matchSrcBase := filepath.Base(matchSrcAbs)
			// default destination uses the path as it is in the source project
			matchDstAbs := filepath.Join(dstDirAbs, matchSrcRel)
			// custom destination if specified
			if s.config.dst[i] != "" {
				matchDstAbs = filepath.Join(dstDirAbs, s.config.dst[i], matchSrcBase)
			}
			// determine what the destination directory is
			matchDstDir := matchDstAbs
			if !matchSrcInfo.IsDir() {
				matchDstDir = filepath.Dir(matchDstAbs)
			}
			// create the destination directory if it doesn't exist
			if _, err := os.Stat(matchDstDir); os.IsNotExist(err) {
				_ = os.MkdirAll(matchDstDir, os.ModePerm)
			}
			// print cli output
			if !s.config.quiet {
				fmt.Println("copying:", matchDstAbs)
			}
			// copy
			err := copy.Copy(matchSrcAbs, matchDstAbs)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func GetSSHKeyAuth(privateSshKeyFile string) transport.AuthMethod {
	var auth transport.AuthMethod
	sshKey, _ := ioutil.ReadFile(privateSshKeyFile)
	signer, _ := ssh.ParsePrivateKey([]byte(sshKey))
	auth = &go_git_ssh.PublicKeys{User: "git", Signer: signer}
	return auth
}
