package ai

const (
	ALI_BAILIAN = "ali_bai_lian"
)

type Config struct {
	Dialect    string
	AliBaiLian AliBaiLianConfig
}

type AliBaiLianConfig struct {
	AccessKey       string
	AccessSecretKey string
	AgentKey        string
	AppId           string
}
