package promptdata

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func PromptData(prompt string) string {
	fmt.Print(prompt + ": ")
	data, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	data = strings.TrimSpace(data)
	return data
}

func PrintError(value any) {
	switch t := value.(type) {
	case string:
		color.Red(t)
	case int:
		color.Red("Код ошибки: %d", t)
	case error:
		color.Red(t.Error())
	default:
		fmt.Println("Неизвестная ошибка!")
	}
}
