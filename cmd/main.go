package main

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	k8sauth "github.com/argoproj/argo-cd/v2/cmd/argocd-k8s-auth/commands"
	apiserver "github.com/argoproj/argo-cd/v2/cmd/argocd-server/commands"
	cli "github.com/argoproj/argo-cd/v2/cmd/argocd/commands"
)

const (
	binaryNameEnv = "ARGOCD_BINARY_NAME"
)

func main() {
	var command *cobra.Command

	binaryName := filepath.Base(os.Args[0])
	if val := os.Getenv(binaryNameEnv); val != "" {
		binaryName = val
	}
	switch binaryName {
	case "argocd-server":
		command = apiserver.NewCommand()
	case "argocd-k8s-auth":
		command = k8sauth.NewCommand()
	default:
		command = cli.NewCommand()
	}

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
