package cloud

type AuthOpts struct {
	Token string `json:"token"`
	Host  string
	User  string
	Key   string
	VM    string
}

type BunnyConfig struct {
	Version string               `json:"version"`
	Clouds  map[string]CloudProp `json:"clouds"`
	Apps    map[string]AppProp   `json:"apps"`
	VMs     map[string]VMProp    `json:"vms"`
}

type VMProp struct {
	Host         string `json:"host"`
	User         string `json:"user"`
	Key          string `json:"key"`
	RegisteredAt string `json:"registered_at"`
}

type CloudProp struct {
	Provider        string `json:"provider"`
	AuthenticatedAt string `json:"authenticated_at"`

	// Hetzner
	Token string `json:"token,omitempty"`

	// AWS
	AccessKeyID     string `json:"access_key_id,omitempty"`
	SecretAccessKey string `json:"secret_access_key,omitempty"`

	// GCP
	ServiceAccountJSON string `json:"service_account_json,omitempty"`
	ProjectID          string `json:"project_id,omitempty"`

	// Azure
	TenantID       string `json:"tenant_id,omitempty"`
	ClientID       string `json:"client_id,omitempty"`
	ClientSecret   string `json:"client_secret,omitempty"`
	SubscriptionID string `json:"subscription_id,omitempty"`

	// Shared
	Region string `json:"region,omitempty"`
}

type AppProp struct {
	VM            string            `json:"vm"`
	Image         string            `json:"image"`
	ContainerID   string            `json:"container_id,omitempty"`
	Port          string            `json:"port"`
	Cloud         string            `json:"cloud,omitempty"`
	InstanceType  string            `json:"instance_type,omitempty"`
	ProvisionedBy bool              `json:"provisioned_by_bunny"`
	InstanceID    string            `json:"instance_id,omitempty"`
	DeployedAt    string            `json:"deployed_at,omitempty"`
	Env           map[string]string `json:"env,omitempty"`
}
