package names

// some names

// OperatorConfig is the name of the CRD that defines the complete
// operator configuration
const OPERATOR_CONFIG = "default"

// CLUSTER_CONFIG is the name of the higher-level cluster configuration
// and status object.
const CLUSTER_CONFIG = "cluster"

// APPLIED_PREFIX is the prefix applied to the config maps
// where we store previously applied configuration
const APPLIED_PREFIX = "applied-"

// APPLIED_NAMESPACE is the namespace where applied configuration
// configmaps are stored.
// Should match 00_namespace.yaml
const APPLIED_NAMESPACE = "openshift-network-operator"
