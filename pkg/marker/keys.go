package marker

type Key = string

const (
	// Prefix is the common base for Bottlerocket's related annotations.
	Prefix = "bottlerocket.aws"

	// NodeSelectorLabel is used to identify controlled nodes in Kubernetes
	// selectors.
	NodeSelectorLabel = UpdaterInterfaceVersionKey
	// PodSelectorLabel is used to identify Pods participating with the
	// operator.
	PodSelectorLabel = UpdaterInterfaceVersionKey

	// UpdateAvailableKey is used to identify a Node as having an update
	// available. The value itself is not checked at this time but may be used
	// to communicate a version at a later time.
	UpdateAvailableKey Key = Prefix + "/update-available"
	// UpdaterInterfaceVersionKey is where the compatibility version is posted for the
	// given Node.
	UpdaterInterfaceVersionKey Key = Prefix + "/updater-interface-version"
	// OperatorVersionKey is where the compatibility version is posted for the
	// given Node. This version describes the understood "protocol" between
	// Operating Controller and the managed Nodes.
	OperatorVersionKey Key = Prefix + "/operator-version"
	// NodeActionWanted provides the Node with the Controller's wanted action to
	// make update progress.
	NodeActionWanted Key = Prefix + "/action-wanted"
	// NodeActionActiveStatus provides progress information on a
	NodeActionActiveState Key = Prefix + "/action-state"
	// NodeActionActive provides the acknowledged and acted-upon action that was
	// wanted of a Node.
	NodeActionActive Key = Prefix + "/action-active"
)
