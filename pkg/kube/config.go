// Package kube provides utilities for working with Kubernetes configurations.
package kube

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"kcl-lang.io/krm-kcl/pkg/kube"
)

var KubeConfigFlags = genericclioptions.NewConfigFlags(false)

func AddKubeConfigFlags(cmd *cobra.Command) {
	namespace := "default"
	// Try to read the default namespace from the current context.
	if ns, _, err := KubeConfigFlags.ToRawKubeConfigLoader().Namespace(); err == nil {
		namespace = ns
	}
	KubeConfigFlags.Namespace = &namespace

	cmd.PersistentFlags().StringVar(KubeConfigFlags.KubeConfig, "kubeconfig", kube.GetKubeConfigPath(), "Path to the kubeconfig file.")
	cmd.PersistentFlags().StringVar(KubeConfigFlags.Context, "kube-context", "", "The name of the kubeconfig context to use.")
	cmd.PersistentFlags().StringVar(KubeConfigFlags.Impersonate, "kube-as", "", "Username to impersonate for the operation. User could be a regular user or a service account in a namespace.")
	cmd.PersistentFlags().StringArrayVar(KubeConfigFlags.ImpersonateGroup, "kube-as-group", nil, "Group to impersonate for the operation, this flag can be repeated to specify multiple groups.")
	cmd.PersistentFlags().StringVar(KubeConfigFlags.ImpersonateUID, "kube-as-uid", "", "UID to impersonate for the operation.")
	cmd.PersistentFlags().StringVar(KubeConfigFlags.BearerToken, "kube-token", "", "Bearer token for authentication to the API server.")
	cmd.PersistentFlags().StringVar(KubeConfigFlags.APIServer, "kube-server", "", "The address and port of the Kubernetes API server.")
	cmd.PersistentFlags().StringVar(KubeConfigFlags.TLSServerName, "kube-tls-server-name", "", "Server name to use for server certificate validation. If it is not provided, the hostname used to contact the server is used.")
	cmd.PersistentFlags().StringVar(KubeConfigFlags.CertFile, "kube-client-certificate", "", "Path to a client certificate file for TLS.")
	cmd.PersistentFlags().StringVar(KubeConfigFlags.KeyFile, "kube-client-key", "", "Path to a client key file for TLS.")
	cmd.PersistentFlags().StringVar(KubeConfigFlags.CAFile, "kube-certificate-authority", "", "Path to a cert file for the certificate authority.")
	cmd.PersistentFlags().BoolVar(KubeConfigFlags.Insecure, "kube-insecure-skip-tls-verify", false, "if true, the Kubernetes API server's certificate will not be checked for validity. This will make your HTTPS connections insecure.")
	cmd.PersistentFlags().StringVarP(KubeConfigFlags.Namespace, "namespace", "n", *KubeConfigFlags.Namespace, "The the namespace scope for the operation.")
}
