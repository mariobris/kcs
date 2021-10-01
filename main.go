package main

import (
	"fmt"
	"github.com/mariobris/pecolify"
	"os"
	"path"
	"path/filepath"
)

var files []string

func main() {
	// Instanciate pecolify
	pf := pecolify.New()

	// Pass data to pecolify
	kubeDir, exists := os.LookupEnv("KCS_KUBEDIR")
	if exists {
		searchFiles(kubeDir)
	}
	//err := filepath.Walk(kubeDir, func(path string, info os.FileInfo, err error) error {
	//	// We do not want to see directories
	//	if info.IsDir() {
	//		return nil
	//	}
	//	files = append(files, path)
	//	return nil
	//})
	//if err != nil {
	//	panic(err)
	//}

	// Append default k8s config file in home dir if exists
	homeDir, _ := os.UserHomeDir()
	defaultKubeconfig := path.Join(homeDir, ".kube/config")

	if _, err := os.Stat(defaultKubeconfig); err == nil {
		files = append(files, defaultKubeconfig)
	}

	// pecolify!
	selected, err := pf.Transform(files)
	if err != nil {
		fmt.Printf("Error was occured: %v\n", err)
		return
	}

	fmt.Printf("export KUBECONFIG=%s\n", selected)
}

func searchFiles(kubeDir string) {
	err := filepath.Walk(kubeDir, func(path string, info os.FileInfo, err error) error {
		// We do not want to see directories
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
}
