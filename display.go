package display

package disp

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/hokaccha/go-prettyjson"
)

var LINEWIDTH int = 80
var Out io.Writer = os.Stdout

var (
	lineThink = "═"
	lineThin  = "━"
)

func BarThick() {
	var line = strings.Repeat(lineThink, LINEWIDTH)
	fmt.Fprintln(Out, color.CyanString(line))
}

func BarThickMessage(msg string) {
	llen := LINEWIDTH - (len(msg) + 5)

	var line = ""
	if llen >= 0 {
		line = strings.Repeat(lineThink, llen)
	}

	fmt.Fprintf(Out, "%s%s%s%s%s\n",
		color.CyanString(line),
		color.CyanString("╡ "),
		msg,
		color.CyanString(" ╞"),
		color.CyanString(lineThink),
	)
}

func BarThin() {
	var line = strings.Repeat(lineThin, LINEWIDTH)
	fmt.Fprintln(Out, color.CyanString(line))
}

func BarThinMessage(msg string) {
	llen := LINEWIDTH - (len(msg) + 5)
	var line = ""
	if llen >= 0 {
		line = strings.Repeat(lineThin, llen)
	}

	fmt.Fprintf(Out, "%s%s%s%s%s\n",
		color.CyanString(lineThin),
		color.CyanString(" ┝"),
		msg,
		color.CyanString("┥ "),
		color.CyanString(line),
	)
}

func KeyValue(key string, val interface{}, size int) {
	key = fmt.Sprintf("%-*s", size, key)
	val = fmt.Sprint(val)
	fmt.Fprintf(Out, "%s: %s\n", color.MagentaString(key), val)
}

func KeyValueFromMap(data map[string]interface{}) {
	maxLen := 0
	for k, _ := range data {
		if len(k) > maxLen {
			maxLen = len(k)
		}
	}

	for k, v := range data {
		KeyValue(k, v, maxLen)
	}
}

func JsonObj(name string, obj interface{}) {
	BarThinMessage(name)
	s, _ := prettyjson.Marshal(obj)
	fmt.Fprintln(Out, string(s))
	BarThin()
}

func GetAnswerString(question string, defaultVal string) (out string) {
	fmt.Print(color.MagentaString(fmt.Sprintf("%s: ", question)))
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	out = scanner.Text()
	return
}

func GetAnswerIntFromList(msg string, options []string) int {

	msg = fmt.Sprintf("%s: ", msg)
	fmt.Println(color.MagentaString(msg))
	for i, o := range options {
		fmt.Printf("[%d] %s\n", i, o)
	}
	fmt.Print(color.MagentaString("Enter value: "))

	txtValue := "0"
	fmt.Scanln(&txtValue)

	intValue, err := strconv.Atoi(txtValue)
	if err != nil {
		fmt.Printf("%s Failed to convert value (%s) to int.\n", color.RedString("[ERROR]"), txtValue)
	}
	return intValue
}

func GetAnswerBool(msg string) bool {
	ans := GetAnswerString(msg, "n")
	return ans == "y"
}

func PrintError(msg string) {
	fmt.Printf("%s %s\n", color.RedString("[ERROR]"), msg)
}