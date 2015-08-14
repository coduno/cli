package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"

	"golang.org/x/oauth2"

	"github.com/coduno/captain"
)

// Commands lists the available commands and help topics.
// The order here is the order in which they are printed by 'go help'.
var commands = []*captain.Command{
	cmdAuth,
	// cmdVersion,
	// cmdLogin,
	// cmdChallenge,
	// hlpIssue,
	// cmdRun,
	// cmdPrepare,
	// cmdPush,
	// cmdSSH,
}

var endpoint = oauth2.Endpoint{
	AuthURL:  "https://api.cod.uno/oauth/auth",
	TokenURL: "https://api.cod.uno/oauth/token",
}

func main() {
	/*var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
		tok, err := conf.Exchange(oauth2.NoContext, code)
		if err != nil {
			log.Fatal(err)
		}

		client := conf.Client(oauth2.NoContext, tok)
		client.Get("...")
	*/
	// 	fmt.Fprint(os.Stderr,
	// 		`
	//     @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
	//     @  WARNING: THIS IS EARLY PREVIEW CODE  @
	//     @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
	//     | Do not expect functionality. APIs and |
	//     | underlying architecture are changing  |
	//     | constantly and without notice. If you |
	//     | want to help building Coduno, please  |
	//     | contact root@cod.uno.                 |
	//     \~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~/
	//
	// `)
	flag.Usage = func() {
		printUsage(os.Stdout)
	}
	flag.Parse()
	log.SetFlags(0)

	args := flag.Args()
	if len(args) < 1 {
		printUsage(os.Stderr)
		os.Exit(2)
	}

	if args[0] == "help" {
		help(args[1:])
		return
	}

	captain.Run(commands, flag.Args())

	//fmt.Fprintf(os.Stderr, "coduno: unknown subcommand %q\nRun 'coduno help' for usage.\n", args[0])
	os.Exit(2)
}

var usageTemplate = `This is the command line interface for Coduno.

Usage:

	coduno [arguments] (command [arguments])...

The top level commands are:
{{range .}}{{if .Runnable}}
    {{.Name | printf "%-11s"}} {{.Short}}{{end}}{{end}}

Use "coduno help [command]" for more information about a command.

Additional help topics:
{{range .}}{{if not .Runnable}}
    {{.Name | printf "%-11s"}} {{.Short}}{{end}}{{end}}

Use "coduno help [topic]" for more information about that topic.

`

var helpTemplate = `{{if .Runnable}}usage: coduno {{.UsageLine}}

{{end}}{{.Long | trim}}
`

var documentationTemplate = `
`

// tmpl executes the given template text on data, writing the result to w.
func tmpl(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	t.Funcs(template.FuncMap{"trim": strings.TrimSpace, "capitalize": capitalize})
	template.Must(t.Parse(text))
	if err := t.Execute(w, data); err != nil {
		panic(err)
	}
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToTitle(r)) + s[n:]
}

func printUsage(w io.Writer) {
	tmpl(w, usageTemplate, commands)
}

// help implements the 'help' command.
func help(args []string) {
	// if len(args) == 0 {
	// 	printUsage(os.Stdout)
	// 	// not exit 2: succeeded at 'go help'.
	// 	return
	// }
	// if len(args) != 1 {
	// 	fmt.Fprintf(os.Stderr, "usage: go help command\n\nToo many arguments given.\n")
	// 	os.Exit(2) // failed at 'go help'
	// }
	//
	// arg := args[0]
	//
	// // 'go help documentation' generates doc.go.
	// if arg == "documentation" {
	// 	buf := new(bytes.Buffer)
	// 	printUsage(buf)
	// 	usage := &Command{Long: buf.String()}
	// 	tmpl(os.Stdout, documentationTemplate, append([]*Command{usage}, commands...))
	// 	return
	// }
	//
	// for _, cmd := range commands {
	// 	if cmd.Name() == arg {
	// 		tmpl(os.Stdout, helpTemplate, cmd)
	// 		// not exit 2: succeeded at 'go help cmd'.
	// 		return
	// 	}
	// }
	//
	// fmt.Fprintf(os.Stderr, "Unknown help topic %#q.  Run 'go help'.\n", arg)
	// os.Exit(2) // failed at 'go help cmd'
}

func fatalf(status int, format string, args ...interface{}) {
	errorf(format, args...)
	os.Exit(status)
}

func errorf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
}

func run(cmdargs ...interface{}) {
	cmdline := stringList(cmdargs...)
	fmt.Printf("%s\n", strings.Join(cmdline, " "))

	cmd := exec.Command(cmdline[0], cmdline[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		errorf("%v", err)
	}
}

// stringList's arguments should be a sequence of string or []string values.
// stringList flattens them into a single []string.
func stringList(args ...interface{}) []string {
	var x []string
	for _, arg := range args {
		switch arg := arg.(type) {
		case []string:
			x = append(x, arg...)
		case string:
			x = append(x, arg)
		default:
			panic("stringList: invalid argument")
		}
	}
	return x
}
