package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var in string
var out string

func init() {
	flag.StringVar(&in, "in", "", "-in path")
	flag.StringVar(&out, "out", "", "-out path")
	flag.Parse()
}

func main() {
	files, err := ioutil.ReadDir(in)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	merged := make([]byte, 0)
	imports := make(map[string]bool, 0)
	for _, file := range files {
		n := file.Name()
		if n == "vendor" {
			fmt.Println("Warning: merging the vendor folder is not supported")
			continue
		}

		if strings.HasSuffix(n, ".go") && !strings.Contains(n, "_test") {
			f, err := os.Open(fmt.Sprintf("%s/%s", in, n))
			if err != nil {
				fmt.Println(err)
				continue
			}

			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				t := scanner.Text()
				if strings.HasPrefix(t, "package") {
					continue
				}

				if strings.HasPrefix(t, "import (") {
					for scanner.Scan() {
						line := scanner.Text()
						if line == ")" {
							break
						} else if len(strings.TrimSpace(line)) > 0 {
							// remove tab & quotes
							unquoted := strings.Replace(line[1:], `"`, ``, 2)
							imports[unquoted] = true
						}
					}
					continue
				}

				if strings.HasPrefix(t, `import "`) {
					unquoted := strings.Replace(t[len(`import `):], `"`, ``, 2)
					imports[unquoted] = true
					continue
				}

				merged = append(merged, fmt.Sprintln(t)...)
			}
			f.Close()
		}
	}

	header := make([]byte, 0)
	header = append(header, fmt.Sprint("package main\n\n")...)
	if len(imports) > 0 {
		header = append(header, fmt.Sprint("import (\n")...)
		for k := range imports {
			s := fmt.Sprintf("\t\"%s\"\n", k)
			header = append(header, s...)
		}
		header = append(header, fmt.Sprint(")\n")...)
	}
	err = ioutil.WriteFile(out, append(header, merged...), 0777)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
