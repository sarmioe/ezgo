package main

import (
	"archive/zip"
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	version = "ezgoit_V0.0.3"
	hden    = `If you want see it , Please view this page :https://github.com/Sarmioe/ezgoit/blob/main/README.md`
)

var targets = []struct {
	os   string
	arch string
}{
	{"windows", "amd64"},
	{"windows", "386"},
	{"windows", "arm"},
	{"windows", "arm64"},
	{"linux", "amd64"},
	{"linux", "386"},
	{"linux", "arm"},
	{"linux", "arm64"},
	{"linux", "ppc64"},
	{"linux", "ppc64le"},
	{"linux", "mips"},
	{"linux", "mipsle"},
	{"linux", "mips64"},
	{"linux", "mips64le"},
	{"linux", "riscv64"},
	{"darwin", "amd64"},
	{"darwin", "arm64"},
	{"freebsd", "amd64"},
	{"freebsd", "386"},
	{"freebsd", "arm"},
	{"freebsd", "arm64"},
	{"openbsd", "amd64"},
	{"openbsd", "386"},
	{"openbsd", "arm"},
	{"openbsd", "arm64"},
	{"netbsd", "amd64"},
	{"netbsd", "386"},
	{"netbsd", "arm"},
	{"netbsd", "arm64"},
	{"dragonfly", "amd64"},
	{"solaris", "amd64"},
	{"plan9", "amd64"},
	{"plan9", "386"},
	{"plan9", "arm"},
	{"aix", "ppc64"},
	{"illumos", "amd64"},
	{"hurd", "amd64"},
}

func downloadZip(url, dest string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download ZIP file: %w", err)
	}
	defer resp.Body.Close()

	out, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save ZIP file: %w", err)
	}

	return nil
}
func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return fmt.Errorf("failed to open zip file: %w", err)
	}
	defer r.Close()

	if err := os.MkdirAll(dest, 0755); err != nil {
		return fmt.Errorf("failed to create extraction directory: %w", err)
	}

	for _, file := range r.File {
		filePath := filepath.Join(dest, file.Name)

		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(filePath, file.Mode()); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
			continue
		}

		outFile, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
		defer outFile.Close()

		rc, err := file.Open()
		if err != nil {
			return fmt.Errorf("failed to open file in zip: %w", err)
		}
		defer rc.Close()

		_, err = io.Copy(outFile, rc)
		if err != nil {
			return fmt.Errorf("failed to extract file: %w", err)
		}
	}

	return nil
}
func buildSourceCode(srcDir string) error {
	cmd := exec.Command("go", "build")
	cmd.Dir = srcDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("build failed: %w", err)
	}

	return nil
}
func getVersion(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
func enver() {
	fmt.Println("Checking environment...")
	if _, err := exec.LookPath("git"); err != nil {
		fmt.Println("Git not found.")
		os.Exit(0)
	}
	if _, err := exec.LookPath("go"); err != nil {
		fmt.Println("Go not found.")
		os.Exit(0)
	}
	gitVersion, err := getVersion("git", "--version")
	if err != nil {
		fmt.Println("Error getting Git version:", err)
		os.Exit(0)
	} else {
		fmt.Println("Git version:", gitVersion)
	}
	goVersion, err := getVersion("go", "version")
	if err != nil {
		fmt.Println("Error getting Go version:", err)
		os.Exit(0)
	} else {
		fmt.Println("Go version:", goVersion)
	}
	fmt.Println("All the environment is ready.")
}
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
func atfs() {
	versione := flag.Bool("v", false, "Display Version")
	help := flag.Bool("h", false, "Display Help")
	ezgoupdate := flag.Bool("u", false, "Update ezgoit")
	checkenv := flag.Bool("env", false, "Check environment")
	build := flag.Bool("b", false, "Build the go project.")
	buildall := flag.Bool("ba", false, "Build all the go project.")
	gitpush := flag.Bool("cm", false, "Commit and push the project to remote repository.")
	gitpushall := flag.Bool("cmt", false, "Commit and push the project to remote repository in every seconds.")
	flag.Parse()
	if *versione {
		fmt.Println("Version is:" + version)
		fmt.Println("Will download the latest version and update the program?")
		fmt.Println("Please use 'ezgo -update' or view : https://github.com/Sarmioe/ezgoit/releases to update the program.")
		os.Exit(0)
	}
	if *help {
		fmt.Println(hden)
		os.Exit(0)
	}
	if *ezgoupdate {
		fmt.Println("ezgoit Start run build to update , download zip from https://github.com/Sarmioe/ezgoit/archive/refs/heads/main.zip")
		zipURL := "https://github.com/Sarmioe/ezgoit/archive/refs/heads/main.zip"
		zipDest := "source.zip"
		extractDir := "source"

		fmt.Println("Downloading ZIP file...")
		if err := downloadZip(zipURL, zipDest); err != nil {
			fmt.Println("Error downloading ZIP file:", err)
			return
		}

		fmt.Println("Extracting ZIP file...")
		if err := unzip(zipDest, extractDir); err != nil {
			fmt.Println("Error extracting ZIP file:", err)
			return
		}
		enver()
		fmt.Println("Building source code...")
		if err := buildSourceCode(extractDir); err != nil {
			fmt.Println("Error building source code:", err)
			os.Exit(1)
		}

		fmt.Println("Build complete! The program is ready.")
		fmt.Println("After 5 seconds , the program will be auto exit , you need restart it.")
		os.Exit(5)
	}
	if *checkenv {
		enver()
		os.Exit(0)
	}
	if *build {
		fmt.Println("Start build your go project.")
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter the absolute path to the Go project: ")
		projectPath, _ := reader.ReadString('\n')
		projectPath = strings.TrimSpace(projectPath)

		if _, err := os.Stat(projectPath); os.IsNotExist(err) {
			fmt.Println("The specified path does not exist. Please try again.")
			os.Exit(1)
		}

		fmt.Printf("Enter target OS (default: %s): ", runtime.GOOS)
		targetOS, _ := reader.ReadString('\n')
		targetOS = strings.TrimSpace(targetOS)
		if targetOS == "" {
			targetOS = runtime.GOOS
		}

		fmt.Printf("Enter target architecture (default: %s): ", runtime.GOARCH)
		targetArch, _ := reader.ReadString('\n')
		targetArch = strings.TrimSpace(targetArch)
		if targetArch == "" {
			targetArch = runtime.GOARCH
		}
		fmt.Println("INFO:If your opriting system is windows , please add ", "'.exe'", ".")
		fmt.Print("Enter output binary name (default: Go-project.exe): ")
		outputName, _ := reader.ReadString('\n')
		outputName = strings.TrimSpace(outputName)
		if outputName == "" {
			outputName = "Go-project.exe"
		}

		outputPath := filepath.Join(projectPath, outputName)

		fmt.Println("Setting up build environment...")
		env := os.Environ()
		env = append(env, "GOOS="+targetOS)
		env = append(env, "GOARCH="+targetArch)

		cmd := exec.Command("go", "build", "-o", outputPath, projectPath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Env = env

		fmt.Println("Starting the build process...")
		if err := cmd.Run(); err != nil {
			fmt.Printf("Build failed: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Build succeeded! Output file: %s\n", outputPath)
		os.Exit(0)
	}
	if *buildall {
		fmt.Println("Start build your all the go project.")
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the path to the Go project: ")
		projectPath, _ := reader.ReadString('\n')
		projectPath = strings.TrimSpace(projectPath)

		if _, err := os.Stat(projectPath); os.IsNotExist(err) {
			fmt.Println("The specified path does not exist. Please try again.")
			os.Exit(1)
		}

		fmt.Print("Enter abstule output binary base name (default: Go-project): ")
		outputBaseName, _ := reader.ReadString('\n')
		outputBaseName = strings.TrimSpace(outputBaseName)
		if outputBaseName == "" {
			outputBaseName = "Go-project"
		}

		outputDir := filepath.Join(projectPath, "build")
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			fmt.Printf("Failed to create output directory: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Starting the build process for all platforms...")
		fmt.Println("The ended output filename just like : Go-project-windows-amd64.exe")
		for _, target := range targets {
			targetOS := target.os
			targetArch := target.arch

			outputFile := fmt.Sprintf("%s-%s-%s", outputBaseName, targetOS, targetArch)
			if targetOS == "windows" {
				outputFile += ".exe"
			}
			outputPath := filepath.Join(outputDir, outputFile)
			env := os.Environ()
			env = append(env, "GOOS="+targetOS, "GOARCH="+targetArch)

			cmd := exec.Command("go", "build", "-o", outputPath, projectPath)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Env = env

			fmt.Printf("Building for %s/%s...\n", targetOS, targetArch)
			if err := cmd.Run(); err != nil {
				fmt.Printf("Build failed for %s/%s: %v\n", targetOS, targetArch, err)
				continue
			}
			fmt.Println("At the floader:", outputPath)
		}
		fmt.Println("Build succeeded: %s\n")
		fmt.Println("Created all the 12 Systems , 11 archs , and 36 output files.")
		os.Exit(0)
		if *gitpush {
			fmt.Println("Start push your project to remote repository.")

		}
		if *gitpushall {
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
				cmd := exec.Command("git", "-C", repoPath, "add", ".")
				err := cmd.Run()
				if err != nil {
					fmt.Printf("Error adding changes: %v\n", err)
					time.Sleep(syncInterval)
					continue
				}

				cmd = exec.Command("git", "-C", repoPath, "status", "--porcelain")
				out, err := cmd.Output()
				if err != nil {
					fmt.Printf("Error checking repository status: %v\n", err)
					time.Sleep(syncInterval)
					continue
				}

				if len(out) > 0 {
					fmt.Println("Changes detected in the repository.")
					fmt.Print("Do you want to sync these changes? (y/n): ")
					scanner.Scan()
					response := scanner.Text()

					if response == "y" {
						commitMessage := generateRandomCommitMessage()
						fmt.Printf("Committing changes with message: %s\n", commitMessage)

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
					} else if response == "n" {
						fmt.Println("Skipping sync. Waiting for the next check...")
					} else {
						fmt.Println("Invalid input. Please enter 'y' or 'n'.")
					}
				} else {
					fmt.Println("No changes detected , If you not need sync , please enter 'ctrl+c' to exit.")
				}

				time.Sleep(syncInterval)
			}
		}
	}
}
func main() {
	fmt.Println("Special edition for the Year of the Snake.")
	fmt.Println("Welcome to ezgoit!")
	fmt.Println("Powered by Sarmioe and Golang V1.23.4")
	atfs()
	fmt.Println("To get help document , view this page :https://github.com/Sarmioe/ezgoit/blob/main/README.md")
	fmt.Println("Now , you no add any bool value , the program will be exit...")
}
