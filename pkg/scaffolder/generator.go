package scaffolder

import (
	"bytes"
	"embed"
	"fmt"
	"strings"
	"text/template"
	"time"
)

//go:embed templates/*
var templateFS embed.FS

// Scaffolder generates community health files.
type Scaffolder struct {
	Stack       TechStack
	Owner       string
	LicenseType string
	BotName     string
	BotIconURL  string
}

func NewScaffolder(stack TechStack, owner, licenseType, botName, botIconURL string) *Scaffolder {
	if licenseType == "" {
		licenseType = "mit"
	}
	if botName == "" {
		botName = "cbog"
	}
	return &Scaffolder{
		Stack:       stack,
		Owner:       owner,
		LicenseType: strings.ToLower(licenseType),
		BotName:     botName,
		BotIconURL:  botIconURL,
	}
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

	// 2. CODE_OF_CONDUCT.md
	coc, err := templateFS.ReadFile("templates/code_of_conduct.md")
	if err == nil {
		files["CODE_OF_CONDUCT.md"] = string(coc)
	}

	// 3. SECURITY.md
	sec, err := templateFS.ReadFile("templates/security.md")
	if err == nil {
		files["SECURITY.md"] = string(sec)
	}

	// 4. SUPPORT.md
	supportTmpl, err := templateFS.ReadFile("templates/support.md")
	if err == nil {
		sTmpl, err := template.New("support").Parse(string(supportTmpl))
		if err == nil {
			var sBuf bytes.Buffer
			sData := struct {
				Owner string
				Repo  string
			}{
				Owner: s.Owner,
				Repo:  "project",
			}
			if err := sTmpl.Execute(&sBuf, sData); err == nil {
				files["SUPPORT.md"] = sBuf.String()
			}
		}
	}

	// 5. LICENSE (Supports: mit, apache-2.0, bsd-3-clause, gpl-3.0, mpl-2.0)
	licFile := fmt.Sprintf("templates/licenses/%s.txt", s.LicenseType)
	licTmpl, err := templateFS.ReadFile(licFile)
	if err != nil {
		licTmpl, _ = templateFS.ReadFile("templates/licenses/mit.txt")
	}

	if len(licTmpl) > 0 {
		lTmpl, err := template.New("license").Parse(string(licTmpl))
		if err == nil {
			var lBuf bytes.Buffer
			lData := struct {
				Year  int
				Owner string
			}{
				Year:  time.Now().Year(),
				Owner: s.Owner,
			}
			if err := lTmpl.Execute(&lBuf, lData); err == nil {
				files["LICENSE"] = lBuf.String()
			}
		}
	}

	// 6. PR Template
	prTmpl, err := templateFS.ReadFile("templates/pr_template.md")
	if err == nil {
		files[".github/PULL_REQUEST_TEMPLATE.md"] = string(prTmpl)
	}

	// 7. Issue Forms
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

	// 8. cbog.yml configuration with custom bot branding
	cbogCfgTmpl, err := templateFS.ReadFile("templates/cbog.yml")
	if err == nil {
		cfgStr := string(cbogCfgTmpl)
		cfgStr = strings.ReplaceAll(cfgStr, "{{.BotName}}", s.BotName)
		cfgStr = strings.ReplaceAll(cfgStr, "{{.BotIconURL}}", s.BotIconURL)
		files[".github/cbog.yml"] = cfgStr
	}

	return files, nil
}
