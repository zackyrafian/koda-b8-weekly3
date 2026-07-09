package utils

import (
    "golang.org/x/text/language"
    "golang.org/x/text/message"
)

var printer = message.NewPrinter(language.Indonesian)

func FormatIDR(amount int) string {
    return printer.Sprintf("Rp%d", amount)
}