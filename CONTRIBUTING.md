# Contributing Guide

Thank you for contributing! To maintain a high-quality codebase and an automated audit trail for our releases, we follow a specific workflow.

---

## 🏗 1. Branching Strategy
We use the `develop` branch for integration. All work should happen in feature branches using the following naming convention:

**Format:** `type/description`

| Branch Prefix | Purpose | Auto-Label |
| :--- | :--- | :--- |
| `feat/` | New features or functionality | `enhancement` |
| `fix/` | Bug fixes | `bug` |
| `docs/` | Documentation updates | `documentation` |
| `refactor/` | Code cleanup (no logic change) | `refactor` |
| `chore/` | Maintenance & Dependencies | `internal` |

---

## 📝 2. Pull Request Template
When you open a Pull Request, please use the following structure in your description. This helps us track changes for our automated release notes.

### Description
**Linked Issue:** #

### Type of Change
- [ ] `feat/` (New feature)
- [ ] `fix/` (Bug fix)
- [ ] `refactor/` (Code cleanup)
- [ ] `chore/` (Maintenance/Dependencies)
- [ ] `docs/` (Documentation only)
- [ ] `perf/` (Performance improvement)

---

### How Has This Been Tested?
- [ ] Unit Tests passed
- [ ] Integration Tests passed (if applicable)
- [ ] Manually verified on [Environment]

---

### Checklist
- [ ] My branch name follows the convention: `type/description`
- [ ] My code follows the style guidelines of this project
- [ ] I have performed a self-review of my own code
- [ ] I have commented my code in hard-to-understand areas
- [ ] I have made corresponding changes to the documentation
- [ ] My changes generate no new warnings

---

## 🚀 3. Pull Request Titles
We use **Conventional Commits** for PR titles to help the Release Drafter.
**Example:** `feat(api): add parsing logic for incoming messages`

---

## 🏆 4. Contributor Credits
By contributing to this project, your GitHub profile will automatically be featured in the **"New Contributors"** section of our official Release Notes once your PR is merged into `main`!