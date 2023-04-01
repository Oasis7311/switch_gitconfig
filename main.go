package main

import (
	"flag"
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// 解析命令行参数
	configPathPtr := flag.String("f", "", "path to gitconfig file")
	flag.Parse()

	// 检查输入参数是否正确
	if *configPathPtr == "" {
		fmt.Println("Missing -f argument (path to gitconfig file)")
		os.Exit(1)
	}
	configPath, err := filepath.Abs(*configPathPtr)
	if err != nil {
		fmt.Printf("Error finding absolute path of gitconfig file: %v\n", err)
		os.Exit(1)
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Printf("Gitconfig file '%s' does not exist\n", configPath)
		os.Exit(1)
	}

	// 检查 Git 配置文件是否正确
	cfg, err := ini.Load(configPath)
	if err != nil {
		fmt.Printf("Error loading gitconfig file '%s': %v\n", configPath, err)
		os.Exit(1)
	}
	if cfg.Section("user") == nil || cfg.Section("user").Key("name").String() == "" || cfg.Section("user").Key("email").String() == "" {
		fmt.Printf("Error: invalid gitconfig file '%s'\n", configPath)
		os.Exit(1)
	}

	// 切换全局 Git 配置
	globalGitconfigPath := filepath.Join(os.Getenv("HOME"), ".gitconfig")
	if _, err := os.Stat(globalGitconfigPath); err == nil {
		if err := os.Remove(globalGitconfigPath); err != nil {
			fmt.Printf("Error removing global gitconfig file: %v\n", err)
			os.Exit(1)
		}
	}
	if err := os.Symlink(configPath, globalGitconfigPath); err != nil {
		fmt.Printf("Error creating symlink for global gitconfig file: %v\n", err)
		os.Exit(1)
	}
	if err := exec.Command("git", "config", "--global", "include.path", "~/.gitconfig.local").Run(); err != nil { // 将用户配置引入全局配置
		fmt.Printf("Error setting global include.path: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Successfully switched global gitconfig to '%s'\n", configPath)
}
