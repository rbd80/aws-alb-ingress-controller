package config

import "time"

// Config contains the ALB Ingress Controller configuration
type Config struct {
	ClusterName     string
	AWSDebug        bool
	ALBSyncInterval time.Duration
}

var RestrictScheme bool
var RestrictSchemeNamespace string
