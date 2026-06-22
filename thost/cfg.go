package thost

type ApiCfg struct {
	IsTestVersion bool
	FlowPath      string
	UserID        string
	UserPass      string
	AuthKey       string
	AuthToken     string
}

type ApiOpt func(*ApiCfg) error

func WithTestVersion() ApiOpt {
	return func(ac *ApiCfg) error {
		ac.IsTestVersion = true

		return nil
	}
}
