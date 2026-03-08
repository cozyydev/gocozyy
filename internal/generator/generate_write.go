package generator

import (
	"os"
	"path/filepath"
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
