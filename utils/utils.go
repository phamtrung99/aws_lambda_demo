package utils

func NewHeaders() map[string]string {
	return map[string]string{
		"X-Requested-With":             "*",
		"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,x-requested-with",
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "*",
	}
}
