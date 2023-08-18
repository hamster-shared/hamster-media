package parameter

type NftAirdropParam struct {
	FkActivityId    int    `json:"fkActivityId"`
	WalletAddress   string `json:"walletAddress"`
	DeployNetwork   string `json:"deployNetwork"`
	ContractName    string `json:"contractName"`
	ContractAddress string `json:"contractAddress"`
	ContractAbi     string `json:"contractAbi"`
}
