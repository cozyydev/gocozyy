package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func writeEnvFile(cfg Config) error {
	envPath := filepath.Join(cfg.ProjectName, ".env")

	content := `PORT=42069
APP_ENV=local
GOCOZYY_DB_HOST=localhost
GOCOZYY_DB_PORT=5432
GOCOZYY_DB_DATABASE=gocozyy
GOCOZYY_DB_USERNAME=sammy
GOCOZYY_DB_PW=thisismypassword
GOCOZYY_DB_SCHEMA=public
`

	return os.WriteFile(envPath, []byte(content), 0o600)
}

func writeGitignore(cfg Config) error {
	path := filepath.Join(cfg.ProjectName, ".gitignore")

	content := `.env
# Go
bin/
*.exe
*.test
*.out

# Node/Bun/Vite
node_modules/
dist/
.vite/

# IDE/editor
.vscode/
.idea/
.DS_Store
`

	return os.WriteFile(path, []byte(content), 0o644)
}

func fixAppCssLogos(frontendDir string) error {
	appCssPath := filepath.Join(frontendDir, "src", "App.css")
	appTsxPath := filepath.Join(frontendDir, "src", "App.tsx")
	appJsxPath := filepath.Join(frontendDir, "src", "App.jsx")

	cssContent, err := os.ReadFile(appCssPath)
	if err != nil {
		return fmt.Errorf("reading App.css: %w", err)
	}

	newCssContent := string(cssContent)

	if !strings.Contains(newCssContent, ".logos") {
		newCssContent += `

.logos {
  display: flex;
  gap: 2rem;
  justify-content: center;
}
`
	}

	if err := os.WriteFile(appCssPath, []byte(newCssContent), 0o644); err != nil {
		return fmt.Errorf("writing App.css: %w", err)
	}

	if _, err := os.Stat(appTsxPath); err == nil {
		tsxContent, err := os.ReadFile(appTsxPath)
		if err != nil {
			return fmt.Errorf("reading App.tsx: %w", err)
		}
		newTsxContent := strings.Replace(
			strings.Replace(string(tsxContent), "<div>", `<div className="logos">`, 1),
			`<div>
        <a href="https://vite.dev"`,
			`<div className="logos">
        <a href="https://vite.dev"`,
			1,
		)
		if err := os.WriteFile(appTsxPath, []byte(newTsxContent), 0o644); err != nil {
			return fmt.Errorf("writing App.tsx: %w", err)
		}
	} else if _, err := os.Stat(appJsxPath); err == nil {
		jsxContent, err := os.ReadFile(appJsxPath)
		if err != nil {
			return fmt.Errorf("reading App.jsx: %w", err)
		}
		newJsxContent := strings.Replace(string(jsxContent), "<div>", `<div className="logos">`, 1)
		if err := os.WriteFile(appJsxPath, []byte(newJsxContent), 0o644); err != nil {
			return fmt.Errorf("writing App.jsx: %w", err)
		}
	}

	return nil
}
