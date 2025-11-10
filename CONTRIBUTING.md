# Contributing to Golazy

Thank you for your interest in contributing! 🎉  
These guidelines ensure a smooth collaboration and maintain the quality of the library.

---

## 📌 Branch Strategy

- **Default branch:** `develop` → all contributions should target this branch via pull requests.
- **Stable branch:** `main` → production-ready releases, only the repository owner can merge or push here.
- **Feature branches:** Use `feature/<feature-name>` for your work.

**Do not:**
- Push directly to `main` or `develop`.
- Open pull requests to `main`.

---

## 🔧 How to Contribute

### 1. Fork the repository and clone your fork:

```bash
git clone https://github.com/<your-username>/golazy.git
cd golazy
```

### 2. Create a feature branch from `develop`:

```bash
git switch develop
git pull origin develop
git switch -c feature/my-feature
```

### 3. Make changes:
- Write clear, concise, and tested Go code.
- Follow Go conventions: go fmt ./... and go vet ./....
- Comment exported functions, methods, and types.


### 4. Run tests locally:

```bash
go test ./...
```
or
```bash
make test
```

### 5. Commit your changes

```bash
git add .
git commit -m "Add short descriptive message"
```

### 6. Push your branch to your fork

```bash
git push origin feature/my-feature
```

### 7. Open a Pull Request targeting develop:
- Include a description of your changes.
- Reference related issues using `Fixes #<issue-number>` if applicable.
- Follow the PR template if available.

---
## Code Guidelines

- **Formatting:** Always run `go fmt ./...` before committing.
- **Linting:** Optional but recommended: `golangci-lint run`.
- **Testing:** Add tests for all new functionality and bug fixes.
- **Commit messages:** Use concise, descriptive messages. Example:
    ```vbnet
    feat: add lazy loading for struct properties
    fix: handle nil pointer on Load function
    ```

---
## Testing & CI

- All pull requests must pass **CI checks** (unit tests + linting).
    
- You can run tests locally:
    

```bash
go test ./... golangci-lint run
```

---

## 🚫 What Not to Do

- Do not push directly to `develop` or `main`.
    
- Do not open PRs to `main`.
    
- Do not create release tags; only the repository owner can do this.
    
- Avoid large, unrelated changes in a single pull request.
    

---

## 🎯 Release Workflow (for maintainers)

- Releases are done via the `make release` command (patch, minor, major).
    
- Tags follow semantic versioning: `vX.Y.Z`.
    
- Contributors do not create or push tags.
    

---

## 📝 Reporting Issues

- Use **GitHub issues** to report bugs or suggest features.
    
- Include:
    
    - A clear description of the problem.
        
    - Steps to reproduce (if applicable).
        
    - Go version and OS details.
        

---
## Additional Resources

- [Go documentation](https://golang.org/doc/)
    
- [Effective Go](https://golang.org/doc/effective_go.html)
    
- [Go Modules](https://blog.golang.org/using-go-modules)
    

---

Thank you for helping improve Golazy! Your contributions make the library better for everyone. 🙌



