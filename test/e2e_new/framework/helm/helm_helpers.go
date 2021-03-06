// +build e2e_new

package helm

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Azure/aad-pod-identity/test/e2e_new/framework"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	chartName = "aad-pod-identity"
)

// InstallInput is the input for Install.
type InstallInput struct {
	Config         *framework.Config
	ManagedMode    bool
	NamespacedMode bool
}

// Install installs aad-pod-identity via Helm 3.
func Install(input InstallInput) {
	Expect(input.Config).NotTo(BeNil(), "input.Config is required for Helm.Install")

	cwd, err := os.Getwd()
	Expect(err).To(BeNil())

	// Change current working directory to repo root
	// Before installing aad-pod identity through Helm
	os.Chdir("../..")
	defer os.Chdir(cwd)

	args := append([]string{
		"install",
		chartName,
		"charts/aad-pod-identity",
		"--wait",
		fmt.Sprintf("--set=image.repository=%s", input.Config.Registry),
		fmt.Sprintf("--set=mic.tag=%s", input.Config.MICVersion),
		fmt.Sprintf("--set=nmi.tag=%s", input.Config.NMIVersion),
	})

	if input.ManagedMode {
		args = append(args, fmt.Sprintf("--set=operationMode=%s", "managed"))
	}

	helm(args)
}

// Uninstall uninstalls aad-pod-identity via Helm 3.
func Uninstall() {
	helm([]string{
		"uninstall",
		chartName,
	})
}

func Upgrade(config *framework.Config) {
	cwd, err := os.Getwd()
	Expect(err).To(BeNil())

	// Change current working directory to repo root
	// Before installing aad-pod identity through Helm
	os.Chdir("../..")
	defer os.Chdir(cwd)

	args := append([]string{
		"upgrade",
		chartName,
		"charts/aad-pod-identity",
		"--wait",
		fmt.Sprintf("--set=image.repository=%s", config.Registry),
		fmt.Sprintf("--set=mic.tag=%s", config.MICVersion),
		fmt.Sprintf("--set=nmi.tag=%s", config.NMIVersion),
	})

	helm(args)
}

func helm(args []string) {
	By(fmt.Sprintf("helm %s", strings.Join(args, " ")))

	cmd := exec.Command("helm", args...)
	stdoutStderr, err := cmd.CombinedOutput()
	fmt.Printf("%s", stdoutStderr)

	Expect(err).To(BeNil())
}
