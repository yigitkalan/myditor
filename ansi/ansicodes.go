package ansi

const (
	// Text colors
	reset     = "\033[0m"
	black     = "\033[30m"
	red       = "\033[31m"
	green     = "\033[32m"
	yellow    = "\033[33m"
	blue      = "\033[34m"
	magenta   = "\033[35m"
	cyan      = "\033[36m"
	white     = "\033[37m"
	
	// Background colors
	bgBlack   = "\033[40m"
	bgRed     = "\033[41m"
	bgGreen   = "\033[42m"
	bgYellow  = "\033[43m"
	bgBlue    = "\033[44m"
	bgMagenta = "\033[45m"
	bgCyan    = "\033[46m"
	bgWhite   = "\033[47m"
	
	// Text styles
	bold      = "\033[1m"
	dim       = "\033[2m"
	italic    = "\033[3m"
	underline = "\033[4m"
	blink     = "\033[5m"
	reverse   = "\033[7m"
	hidden    = "\033[8m"
	strike    = "\033[9m"
)
