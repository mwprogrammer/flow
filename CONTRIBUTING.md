# Contributing to Flow

Thank you for helping us build Flow! To ensure a high-quality audit trail and to respect everyone's time, we follow a specific workflow.

---

## 🚦 Our Workflow: "Issue First"
To keep the project organized, we require an approved issue before any code is written.

1.  **Open an Issue:** Use our templates to report a bug or suggest a feature.
2.  **Wait for Acceptance:** A maintainer will review the issue and "Accept" it by adding a label or a comment.
3.  **Fork the Project:** Once the issue is accepted, [fork the repository](https://github.com/mwprogrammer/flow/fork) to your own GitHub account.
4.  **Create a Branch:** Create your branch on your fork using the naming convention below.
5.  **Submit PR:** Open a Pull Request from your fork back to our `develop` branch.

---

## 🏗 Branching Strategy
Target the `develop` branch. Use the following naming convention for your branches:

**Format:** `type/short-description`

| Type | Description |
| :--- | :--- |
| `feat/` | New features or functional requirements |
| `fix/` | Bug fixes |
| `docs/` | Documentation changes |
| `refactor/` | Code cleanup (no logic change) |
| `perf/` | Performance improvements |
| `chore/` | Maintenance (dependencies, CI/CD) |

*Example: `feat/message-parser`*

---

## 🚀 The Pull Request Process
1. **Implement your changes** following our style guidelines.
2. **Open a PR** targeting the `develop` branch.
3. **Fill out the Template:** The PR description will pre-fill with our checklist. Please ensure you link the **Linked Issue** number in the description so we can track the progress.

## 📝 Pull Request Titles
We use **Conventional Commits** for automated release notes:
`type(scope): description` (e.g., `feat(parser): add validation logic`)

---

## 🛠 Go Style Guide

To maintain consistency across the codebase, we adhere to the official Go standards.

### 1. Formatting
All code must be formatted using `gofmt` (or `goimports`). 
- **Tip:** Most IDEs can be configured to run `gofmt` automatically on save.

### 2. Standard Guidelines
We follow the official community recommendations:
- **[Effective Go](https://golang.org/doc/effective_go.html):** The primary reference for writing "idiomatic" Go.
- **[Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments):** A collection of common style mistakes and how to avoid them.

### 3. Naming Conventions
- **Packages:** Short, lowercase, single-word names (no underscores or mixedCaps).
- **Interfaces:** Usually end in "er" (e.g., `Reader`, `Writer`, `Parser`).
- **Variables/Functions:** Use `MixedCaps` or `mixedCaps` rather than underscores.

### 4. Error Handling
- Errors should be handled explicitly.
- Use `fmt.Errorf` with the `%w` verb for error wrapping where appropriate.
- **Example:** `if err != nil { return fmt.Errorf("parsing failed: %w", err) }`


## 🏆 Release Notes
Once merged, your contribution and GitHub profile will be automatically added to the next Release once published!