package api

import (
	"fmt"
	"github.com/go-git/go-git/v5"
)

// TODO: Create authenticator github with username and password

func DownloadRepo(url string, dir string) error {
	// Create a new repository object
	repo, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL: url,
	})

	if err != nil {
		return fmt.Errorf("failed to clone repository: %v", err)
	}

	// Print the latest commit hash
	head, err := repo.Head()
	if err != nil {
		return fmt.Errorf("failed to get head: %v", err)
	}

	fmt.Printf("Successfully cloned repository %s to %s (latest commit: %s)\n", url, dir, head.Hash())

	return nil
}
