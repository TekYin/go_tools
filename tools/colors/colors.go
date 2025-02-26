package colors

var (
	// Foreground colors
	reset   = "\033[0m"
	black   = "\033[30m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	white   = "\033[37m"

	// Bright foreground colors
	brightBlack   = "\033[90m"
	brightRed     = "\033[91m"
	brightGreen   = "\033[92m"
	brightYellow  = "\033[93m"
	brightBlue    = "\033[94m"
	brightMagenta = "\033[95m"
	brightCyan    = "\033[96m"
	brightWhite   = "\033[97m"

	// Background colors
	bgBlack   = "\033[40m"
	bgRed     = "\033[41m"
	bgGreen   = "\033[42m"
	bgYellow  = "\033[43m"
	bgBlue    = "\033[44m"
	bgMagenta = "\033[45m"
	bgCyan    = "\033[46m"
	bgWhite   = "\033[47m"

	// Bright background colors
	bgBrightBlack   = "\033[100m"
	bgBrightRed     = "\033[101m"
	bgBrightGreen   = "\033[102m"
	bgBrightYellow  = "\033[103m"
	bgBrightBlue    = "\033[104m"
	bgBrightMagenta = "\033[105m"
	bgBrightCyan    = "\033[106m"
	bgBrightWhite   = "\033[107m"

	// Text formatting
	bold          = "\033[1m"
	dim           = "\033[2m"
	italic        = "\033[3m"
	underline     = "\033[4m"
	blink         = "\033[5m"
	reverse       = "\033[7m"
	hidden        = "\033[8m"
	strikethrough = "\033[9m"
)

type Colorized string

func Red(s string) string {
	return string(Colorized(s).Red())
}

func Green(s string) string {
	return string(Colorized(s).Green())
}

func Yellow(s string) string {
	return string(Colorized(s).Yellow())
}

func Blue(s string) string {
	return string(Colorized(s).Blue())
}

func Magenta(s string) string {
	return string(Colorized(s).Magenta())
}

func Cyan(s string) string {
	return string(Colorized(s).Cyan())
}

func White(s string) string {
	return string(Colorized(s).White())
}

func BrightBlack(s string) string {
	return string(Colorized(s).BrightBlack())
}

func BrightRed(s string) string {
	return string(Colorized(s).BrightRed())
}

func BrightGreen(s string) string {
	return string(Colorized(s).BrightGreen())
}

func BrightYellow(s string) string {
	return string(Colorized(s).BrightYellow())
}

func BrightBlue(s string) string {
	return string(Colorized(s).BrightBlue())
}

func BrightMagenta(s string) string {
	return string(Colorized(s).BrightMagenta())
}

func BrightCyan(s string) string {
	return string(Colorized(s).BrightCyan())
}

func BrightWhite(s string) string {
	return string(Colorized(s).BrightWhite())
}

func (c Colorized) Red() Colorized {
	return Colorized(red + string(c) + reset)
}

func (c Colorized) Green() Colorized {
	return Colorized(green + string(c) + reset)
}

func (c Colorized) Yellow() Colorized {
	return Colorized(yellow + string(c) + reset)
}

func (c Colorized) Blue() Colorized {
	return Colorized(blue + string(c) + reset)
}

func (c Colorized) Magenta() Colorized {
	return Colorized(magenta + string(c) + reset)
}

func (c Colorized) Cyan() Colorized {
	return Colorized(cyan + string(c) + reset)
}

func (c Colorized) White() Colorized {
	return Colorized(white + string(c) + reset)
}

func (c Colorized) BrightBlack() Colorized {
	return Colorized(brightBlack + string(c) + reset)
}

func (c Colorized) BrightRed() Colorized {
	return Colorized(brightRed + string(c) + reset)
}

func (c Colorized) BrightGreen() Colorized {
	return Colorized(brightGreen + string(c) + reset)
}

func (c Colorized) BrightYellow() Colorized {
	return Colorized(brightYellow + string(c) + reset)
}

func (c Colorized) BrightBlue() Colorized {
	return Colorized(brightBlue + string(c) + reset)
}

func (c Colorized) BrightMagenta() Colorized {
	return Colorized(brightMagenta + string(c) + reset)
}

func (c Colorized) BrightCyan() Colorized {
	return Colorized(brightCyan + string(c) + reset)
}

func (c Colorized) BrightWhite() Colorized {
	return Colorized(brightWhite + string(c) + reset)
}

func (c Colorized) BgBlack() Colorized {
	return Colorized(bgBlack + string(c) + reset)
}

func (c Colorized) BgRed() Colorized {
	return Colorized(bgRed + string(c) + reset)
}

func (c Colorized) BgGreen() Colorized {
	return Colorized(bgGreen + string(c) + reset)
}

func (c Colorized) BgYellow() Colorized {
	return Colorized(bgYellow + string(c) + reset)
}

func (c Colorized) BgBlue() Colorized {
	return Colorized(bgBlue + string(c) + reset)
}

func (c Colorized) BgMagenta() Colorized {
	return Colorized(bgMagenta + string(c) + reset)
}

func (c Colorized) BgCyan() Colorized {
	return Colorized(bgCyan + string(c) + reset)
}

func (c Colorized) BgWhite() Colorized {
	return Colorized(bgWhite + string(c) + reset)
}

func (c Colorized) BgBrightBlack() Colorized {
	return Colorized(bgBrightBlack + string(c) + reset)
}

func (c Colorized) BgBrightRed() Colorized {
	return Colorized(bgBrightRed + string(c) + reset)
}

func (c Colorized) BgBrightGreen() Colorized {
	return Colorized(bgBrightGreen + string(c) + reset)
}

func (c Colorized) BgBrightYellow() Colorized {
	return Colorized(bgBrightYellow + string(c) + reset)
}

func (c Colorized) BgBrightBlue() Colorized {
	return Colorized(bgBrightBlue + string(c) + reset)
}

func (c Colorized) BgBrightMagenta() Colorized {
	return Colorized(bgBrightMagenta + string(c) + reset)
}

func (c Colorized) BgBrightCyan() Colorized {
	return Colorized(bgBrightCyan + string(c) + reset)
}

func (c Colorized) BgBrightWhite() Colorized {
	return Colorized(bgBrightWhite + string(c) + reset)
}

func (c Colorized) Bold() Colorized {
	return Colorized(bold + string(c) + reset)
}

func (c Colorized) Dim() Colorized {
	return Colorized(dim + string(c) + reset)
}

func (c Colorized) Italic() Colorized {
	return Colorized(italic + string(c) + reset)
}

func (c Colorized) Underline() Colorized {
	return Colorized(underline + string(c) + reset)
}

func (c Colorized) Blink() Colorized {
	return Colorized(blink + string(c) + reset)
}

func (c Colorized) Reverse() Colorized {
	return Colorized(reverse + string(c) + reset)
}

func (c Colorized) Hidden() Colorized {
	return Colorized(hidden + string(c) + reset)
}

func (c Colorized) Strikethrough() Colorized {
	return Colorized(strikethrough + string(c) + reset)
}
