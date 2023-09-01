package repository

type Bucket string

const (
	AccessTokens  Bucket = "access_tokens"
	RequestTokens Bucket = "request_tokens"
)

type tokenRepository interface {
	Save(chatID int64, token string) error
	Get(chatID int64, bucket Bucket) (string, error)
}
