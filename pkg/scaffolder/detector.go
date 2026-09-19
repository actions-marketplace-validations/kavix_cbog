package scaffolder

import (
	"os"
)

// TechStack represents the detected stack of the repository.
type TechStack string

const (
	StackGo         TechStack = "go"
	StackNode       TechStack = "node"
	StackPython     TechStack = "python"
	StackRust       TechStack = "rust"
	StackUniversal  TechStack = "universal"
)

// DetectStack inspects root repository files to determine the programming language.
func DetectStack(repoDir string) TechStack {
	if fileExists(repoDir + "/go.mod") {
		return StackGo
	}
	if fileExists(repoDir + "/package.json") {
		return StackNode
	}
	if fileExists(repoDir+"/pyproject.toml") || fileExists(repoDir+"/requirements.txt") || fileExists(repoDir+"/setup.py") {
		return StackPython
	}
	if fileExists(repoDir + "/Cargo.toml") {
		return StackRust
	}
	return StackUniversal
}

// GetTestInstructions returns test/build instructions based on the stack.
func GetTestInstructions(stack TechStack) string {
	switch stack {
	case StackGo:
		return "```bash\n   go test -v ./...\n   go vet ./...\n   ```"
	case StackNode:
		return "```bash\n   npm install\n   npm test\n   npm run lint\n   ```"
	case StackPython:
		return "```bash\n   pytest\n   ruff check .\n   ```"
	case StackRust:
		return "```bash\n   cargo test\n   cargo clippy\n   ```"
	default:
		return "```bash\n   # Run your project's test suite and linters locally\n   ```"
	}
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
