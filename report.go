package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

func run_repo(wg *sync.WaitGroup, repo_url string) {
	const author = "Jammes"
	cmd_str_list := []string{
		fmt.Sprintf("git clone --bare %s tmp", repo_url),
		fmt.Sprintf("cd tmp && git log --pretty=format:\"%s %s\" --author=\"%s\" >> ../stats.txt", repo_url, ": %ad - %an : %s", author)
	}
	cmd := exec.Command("sh", "-c", cmd_str)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	fmt.Printf("command is: %s\n", cmd_str)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	wg.Done()
}

func main() {
	const repo_url = [...]string{"https://github.com/lsst/qserv", "https://github.com/lsst/qserv_testdata"}
	const tmp_dir = "tmp"
	wg := new(sync.WaitGroup)

	for _, repo_url := range repo_urls {
		wg.Add(repo_url)
		go run_repo(wg, repo_url)
		wg.Wait()
	}
}
