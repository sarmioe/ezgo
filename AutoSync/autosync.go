package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func generateRandomCommitMessage() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const digits = "0123456789"
	const length = 10

	b := make([]byte, length)
	for i := range b {
		if rand.Intn(2) == 0 {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		} else {
			b[i] = digits[rand.Intn(len(digits))]
		}
	}
	return string(b)
}

func main() {
	fmt.Print("Enter the Git repository path: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	repoPath := scanner.Text()

	fmt.Print("Enter the sync interval (in seconds): ")
	scanner.Scan()
	var syncIntervalSeconds int
	_, err := fmt.Sscanf(scanner.Text(), "%d", &syncIntervalSeconds)
	if err != nil {
		fmt.Printf("Invalid input: %v\n", err)
		return
	}
	syncInterval := time.Duration(syncIntervalSeconds) * time.Second

	rand.Seed(time.Now().UnixNano())

	for {
		cmd := exec.Command("git", "-C", repoPath, "status", "--porcelain")
		out, err := cmd.Output()
		if err != nil {
			fmt.Printf("Error checking repository status: %v\n", err)
			time.Sleep(syncInterval)
			continue
		}

		if len(out) > 0 {
			fmt.Println("Changes detected in the repository.")
			commitMessage := generateRandomCommitMessage()
			fmt.Printf("Committing changes with message: %s\n", commitMessage)

			cmd = exec.Command("git", "-C", repoPath, "add", ".")
			err = cmd.Run()
			if err != nil {
				fmt.Printf("Error adding changes: %v\n", err)
				continue
			}

			cmd = exec.Command("git", "-C", repoPath, "commit", "-m", commitMessage)
			err = cmd.Run()
			if err != nil {
				fmt.Printf("Error committing changes: %v\n", err)
				continue
			}

			cmd = exec.Command("git", "-C", repoPath, "push")
			err = cmd.Run()
			if err != nil {
				fmt.Printf("Error pushing changes: %v\n", err)
				continue
			}

			fmt.Println("Changes committed and pushed successfully.")
		} else {
			fmt.Println("No changes detected.")
		}

		time.Sleep(syncInterval)
	}
}
