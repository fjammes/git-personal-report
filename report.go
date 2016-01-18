package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func run_repo(wg *sync.WaitGroup, repo_url string) {
	l := strings.Split(repo_url, "/")
	repo_name := l[len(l)-1]
	fmt.Printf("INSIDE %s\n", repo_url)
	const author = "Jammes"
	s := []string{}
	s = append(s, ("mkdir -p /tmp/gitreport && cd /tmp/gitreport"))
	s = append(s, fmt.Sprintf("rm -rf %[2]s && git clone --bare %[1]s %[2]s", repo_url, repo_name))
	s = append(s, fmt.Sprintf("cd %s", repo_name))
	s = append(s, fmt.Sprintf("git log --pretty=format:\"%[1]s -> %[2]s\" "+
		"--author=\"%[3]s\" > /tmp/gitreport/stats-%[2]s.txt",
		"%ad - %an : %s",
		repo_name,
		author))
	cmd_str := strings.Join(s, " && ")
	fmt.Printf("Running command: %s\n", cmd_str)
	cmd := exec.Command("sh", "-c", cmd_str)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	wg.Done()
}

func main() {
	repo_urls := [...]string{"https://github.com/lsst/qserv",
		"https://github.com/lsst/xrootd",
		"https://github.com/lsst/mariadb",
		"https://github.com/fjammes/qserv-cluster",
		"https://github.com/fjammes/vagrant-openstack-example",
		"https://github.com/fjammes/packer-openstack-example"}
	const tmp_dir = "tmp"
	wg := new(sync.WaitGroup)

	wg.Add(len(repo_urls))
	for _, repo_url := range repo_urls {
		fmt.Printf("Stat for: %s\n", repo_url)
		go run_repo(wg, repo_url)
	}
	wg.Wait()
	cmd_str := "cat /tmp/gitreport/stats-* | sort | less"
	cmd := exec.Command("sh", "-c", cmd_str)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
