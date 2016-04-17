package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

func GetDoc(lang string, p string) string {
	cmd := exec.Command("./fix", "-d")
	cmd.Env = []string{"LANG=" + lang}
	cmd.Dir = p
	bs, err := cmd.Output()
	if err != nil {
		panic(err.Error())
	}
	return string(bs)
}
func RenderScript(w io.Writer, baseUrl string, pathScript string, doc string) {
	id := strings.Replace(pathScript, "/", ".", -1)
	if strings.TrimSpace(doc) == "" {
		fmt.Printf("W: document of %s is empty\n", id)
		return
	}
	fmt.Fprintf(w, "- [%s](%s/%s)\n", id, baseUrl, pathScript)
	fmt.Fprintln(w)

	buf := bufio.NewReader(strings.NewReader(doc))
	var err error
	var line string
	for err == nil {
		line, err = buf.ReadString('\n')
		fmt.Fprintf(w, "    %s", line)
	}
	fmt.Fprintln(w)
}

func ScanScript(base string) []string {
	fs, err := ioutil.ReadDir(base)
	if err != nil {
		fmt.Println("E:", err)
		return nil
	}
	var r []string
	for _, finfo := range fs {
		if finfo.IsDir() {
			r = append(r, ScanScript(path.Join(base, finfo.Name()))...)
		}
		if finfo.Name() == "fix" {
			r = append(r, base)
			return r
		}
	}
	return r
}

func main() {
	BaseURL := "https://github.com/x-deepin/p/blob/master/"

	scripts := ScanScript(".")
	zhW, err := os.Create("index.zh.md")
	if err != nil {
		fmt.Printf("E:", err)
		return
	}
	defer zhW.Close()

	enW, err := os.Create("index.md")
	if err != nil {
		fmt.Printf("E:", err)
		return
	}
	defer enW.Close()

	for _, s := range scripts {
		RenderScript(zhW, BaseURL, s, GetDoc("zh_CN", s))
	}
	for _, s := range scripts {
		RenderScript(enW, BaseURL, s, GetDoc("en_US", s))
	}
}
