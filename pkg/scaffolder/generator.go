package scaffolder

import (
	"bytes"
	"embed"
	"fmt"
	"text/template"
)

//go:embed templates/*
var templateFS embed.FS

// Scaffolder generates community health files.
type Scaffolder struct {
	Stack TechStack
}

func NewScaffolder(stack TechStack) *Scaffolder {
	return &Scaffolder{Stack: stack}
}

// GenerateFiles generates all community guidelines and issue templates.
func (s *Scaffolder) GenerateFiles() (map[string]string, error) {
	files := make(map[string]string)

	// 1. CONTRIBUTING.md
	contribTmplContent, err := templateFS.ReadFile("templates/contributing.md")
	if err != nil {
		return nil, fmt.Errorf("failed to read contributing template: %w", err)
	}

	tmpl, err := template.New("contributing").Parse(string(contribTmplContent))
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	data := struct {
		TestInstructions string
	}{
		TestInstructions: GetTestInstructions(s.Stack),
	}
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}
	files["CONTRIBUTING.md"] = buf.String()

	// 2. PR Template
	prTmpl, err := templateFS.ReadFile("templates/pr_template.md")
	if err == nil {
		files[".github/PULL_REQUEST_TEMPLATE.md"] = string(prTmpl)
	}

	// 3. Issue Forms
	bugReport, err := templateFS.ReadFile("templates/issue_forms/bug_report.yml")
	if err == nil {
		files[".github/ISSUE_TEMPLATE/bug_report.yml"] = string(bugReport)
	}

	featureReq, err := templateFS.ReadFile("templates/issue_forms/feature_request.yml")
	if err == nil {
		files[".github/ISSUE_TEMPLATE/feature_request.yml"] = string(featureReq)
	}

	configYaml, err := templateFS.ReadFile("templates/issue_forms/config.yml")
	if err == nil {
		files[".github/ISSUE_TEMPLATE/config.yml"] = string(configYaml)
	}

	// 4. cbog.yml configuration
	cbogCfg, err := templateFS.ReadFile("templates/cbog.yml")
	if err == nil {
		files[".github/cbog.yml"] = string(cbogCfg)
	}

	return files, nil
}
