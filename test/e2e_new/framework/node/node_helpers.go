// +build e2e_new

package node

import (
	"context"
	"time"

	"github.com/Azure/aad-pod-identity/test/e2e_new/framework"

	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
)

const (
	listTimeout = 10 * time.Second
	listPolling = 1 * time.Second
)

// ListInput is the input for List.
type ListInput struct {
	Lister framework.Lister
}

// List lists all nodes in the cluster
func List(input ListInput) *corev1.NodeList {
	Expect(input.Lister).NotTo(BeNil(), "input.Lister is required for Node.List")

	nodes := &corev1.NodeList{}
	Eventually(func() error {
		return input.Lister.List(context.TODO(), nodes)
	}, listTimeout, listPolling).Should(Succeed())

	return nodes
}
