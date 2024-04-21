package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type MyFile struct {
	*os.File
}

func main() {
	// https://www.alexedwards.net/blog/serving-static-sites-with-go
	path := "hello.md"
	html := Read(path)
	CheckFile("static/index.html", html)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	fmt.Println("Server is running on port 6969...")
	err := http.ListenAndServe(":6969", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Read(filename string) string { // ta uma bosta isso aqui. Lendo e escrevendo. Arrumar dps
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var html strings.Builder

	scanner := bufio.NewScanner(file)
	// lendo o arquivo linha por linha
	for scanner.Scan() {
		// salvando a linha atual
		content := scanner.Text()

		// checando se linha atual tem caractere '#'
		// if strings.Contains(content, "#") {
		// println(content)
		// return ConvertMarkdownToHTML(content, scanner)
		// }
		html.WriteString(ConvertMarkdownToHTML(content))
	}

	err = scanner.Err()
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo: %s", err)
	}

	return html.String()
}

// pega o conteudo e o scanner vai permitir ler o arquivo
func ConvertMarkdownToHTML(content string) string {
	// para escrever uma nova string
	var html strings.Builder

	// removendo espaços em branco -> ' ', '\t', '\n', '\v', '\f', '\r'
	// line := bytes.TrimSpace(scanner.Bytes())
	switch {
	case strings.HasPrefix(content, "#"):
		count := strings.Count(content, "#")
		switch count {
		case 1:
			str := strings.TrimPrefix(content, "# ")
			html.WriteString(fmt.Sprintf("<h1>%s</h1>", str))
		case 2:
			str := strings.TrimPrefix(content, "## ")
			html.WriteString(fmt.Sprintf("<h2>%s</h2>", str))
		case 3:
			str := strings.TrimPrefix(content, "### ")
			html.WriteString(fmt.Sprintf("<h3>%s</h3>", str))
		case 4:
			str := strings.TrimPrefix(content, "#### ")
			html.WriteString(fmt.Sprintf("<h4>%s</h4>", str))
		case 5:
			str := strings.TrimPrefix(content, "##### ")
			html.WriteString(fmt.Sprintf("<h5>%s</h5>", str))
		case 6:
			str := strings.TrimPrefix(content, "###### ")
			html.WriteString(fmt.Sprintf("<h6>%s</h6>", str))
		}
	case strings.HasPrefix(content, "```"):
		html.WriteString("<pre><code>")
	case strings.HasSuffix(content, "```"):
		html.WriteString("</code></pre>")
	default:
		html.WriteString(fmt.Sprintf("<p>%s</p>", content))
	}

	return html.String()
}

func CheckFile(filename, content string) {
	// if _, err := os.Stat(filename); err == nil {
	// 	// https://stackoverflow.com/questions/33851692/golang-bad-file-descriptor
	// 	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 06444)
	// 	if err != nil {
	// 		log.Fatalf("Erro ao abrir o arquivo: %s", err)
	// 	}
	// 	defer func() {
	// 		if err := f.Close(); err != nil {
	// 			log.Fatalf("Erro ao fechar o arquivo: %s", err)
	// 		}
	// 	}()

	// 	// escrevendo no arquivo
	// 	outfile := &MyFile{f}
	// 	outfile.WriteFile(content)

	// 	return

	// } else if os.IsNotExist(err) {
	// 	// criando o arquivo
	// 	f, err := os.Create(filename)
	// 	if err != nil {
	// 		log.Fatalf("Erro ao criar o arquivo: %s", err)
	// 	}
	// 	defer func() {
	// 		if err := f.Close(); err != nil {
	// 			log.Fatalf("Erro ao fechar o arquivo: %s", err)
	// 		}
	// 	}()

	// 	// escrevendo no arquivo
	// 	outfile := &MyFile{f}
	// 	outfile.WriteFile(content)

	// 	return
	// }

	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Erro ao criar arquivo: %s", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("Erro ao fechar arquivo: %s", err)
		}
	}()
	// writing to the file
	outfile := &MyFile{f}
	outfile.WriteFile(content)
}

func (f *MyFile) WriteFile(content string) {
	// escrevendo no arquivo
	if _, err := f.WriteString(content); err != nil {
		log.Fatalf("Erro ao escrever no arquivo: %s", err)
	}
}
