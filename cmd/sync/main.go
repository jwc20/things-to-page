package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/jwc20/things-to-page/internal/store"
)

const (
	dbPath   = "/Users/cjw/Library/Group Containers/JLMPQHK86H.com.culturedcode.ThingsMac/ThingsData-DC8DV/Things Database.thingsdatabase/main.sqlite"
	htmlPath = "./public/table.html"
)

const tableTemplate = `
<div id="data-table" class="fade-in">
    <table>
        <thead>
            <tr>{{ range .Headers }}<th>{{ . }}</th>{{ end }}</tr>
        </thead>
        <tbody>
            {{ range $row := .Rows }}
            <tr>
                {{ range $col := $.Headers }}
                <td>{{ index $row $col }}</td>
                {{ end }}
            </tr>
            {{ end }}
        </tbody>
    </table>
    <p>Last updated: {{ .Timestamp }}</p>
</div>
`

type PageData struct {
	Timestamp string
	Headers   []string
	Rows      []map[string]interface{}
}

func main() {
	dbConnection, err := store.NewDataProvider(dbPath)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer dbConnection.Close()

	data, err := dbConnection.FetchData()
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}

	var headers []string
	if len(data) > 0 {
		for k := range data[0] {
			headers = append(headers, k)
		}
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	pd := PageData{
		Timestamp: currentTime,
		Headers:   headers,
		Rows:      data,
	}

	f, err := os.Create(htmlPath)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer f.Close()

	tmpl := template.Must(template.New("table").Parse(tableTemplate))
	err = tmpl.Execute(f, pd)
	if err != nil {
		log.Fatalf("Failed to render template: %v", err)
	}

	if err := runGitCommands(); err != nil {
		log.Fatalf("Git automation failed: %v", err)
	}
}

func runGitCommands() error {
	commands := [][]string{
		{"git", "add", "public/table.html"},
		{"git", "commit", "-m", "chore: automated data update"},
		{"git", "push", "origin", "main"},
	}

	for _, args := range commands {
		cmd := exec.Command(args[0], args[1:]...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("command %v failed: %s (%v)", args, string(output), err)
		}
	}

	fmt.Println("Successfully pushed to GitHub")
	return nil
}
