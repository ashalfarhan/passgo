package password

const (
	ALPHA   string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NUMERIC string = "0123456789"
	SYMBOLS string = "~!@#$%^&*"
	ALL     string = NUMERIC + ALPHA + SYMBOLS
)