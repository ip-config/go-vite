package biz

type Reward struct {
	RewardAddr string  `json:"RewardAddr"`
	Name       string  `json:"Name"`
	SecretPub  *string `json:"SecretPub"`
}
