package cloud

type Cloud interface {
	Auth() (string, error)
}

type CloudManager struct {
	Cloud
}

func New(c Cloud) *CloudManager {
	return &CloudManager{
		Cloud: c,
	}
}

func (cm CloudManager) Auth() (string, error) {
	res, err := cm.Cloud.Auth()
	if err != nil {
		return "", err
	}

	return res, nil
}
