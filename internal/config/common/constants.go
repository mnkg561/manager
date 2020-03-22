package common

// Global constants
const (
	// systemNameSpace is the default k8s namespace i.e, kube-system
	SystemNameSpace = "kube-system"

	ManagerServiceAccountName = "keiko-manager-sa"

	ManagerClusterRole = "keiko-manager-cluster-role"

	ManagerClusterRoleBinding = "keiko-manager-cluster-role-binding"

	RBACApiVersion = "rbac.authorization.k8s.io/v1"

	ServiceAccountKind = "ServiceAccount"

	ClusterRoleKind = "ClusterRole"

	ClusterRoleBindingKind = "ClusterRoleBinding"

	InClusterAPIServerAddr = "https://kubernetes.default.svc"
)
