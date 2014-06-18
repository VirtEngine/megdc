package cmd

import (
	"bytes"
	"launchpad.net/gocheck"
	"os"
	"os/exec"
	"testing"
)

func Test(t *testing.T) { gocheck.TestingT(t) }

type S struct {
	stdin   *os.File
	recover []string
}

var _ = gocheck.Suite(&S{})
var manager *Manager

func (s *S) SetUpSuite(c *gocheck.C) {
	targetFile := os.Getenv("HOME") + "/.megam"
	_, err := os.Stat(targetFile)
	if err == nil {
		old := targetFile + ".old"
		s.recover = []string{"mv", old, targetFile}
		exec.Command("mv", targetFile, old).Run()
	} else {
		s.recover = []string{"rm", targetFile}
	}
	f, err := os.Create(targetFile)
	c.Assert(err, gocheck.IsNil)
	f.Write([]byte("http://localhost"))
	f.Close()
}

func (s *S) TearDownSuite(c *gocheck.C) {
	exec.Command(s.recover[0], s.recover[1:]...).Run()
}

func (s *S) SetUpTest(c *gocheck.C) {
	var stdout, stderr bytes.Buffer
	manager = NewManager("gulpd", "0.1", "", &stdout, &stderr, os.Stdin)
	var exiter recordingExiter
	manager.e = &exiter
}


