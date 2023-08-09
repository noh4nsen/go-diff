# go-diff

[![Build](https://github.com/noh4nsen/go-diff/actions/workflows/build.yml/badge.svg?branch=release/1.3.2)](https://github.com/noh4nsen/go-diff/actions/workflows/build.yml)

Tiny action written in Go that returns the files/directories/projects that have been modified between two branches.

### Example:

[Usage example](.github/workflows/example.yml):

```yaml
name: Usage example

on: [workflow_dispatch]

jobs:
    diff:
        runs-on: ubuntu-latest
        steps:
          - name: Checkout code
            uses: actions/checkout@v3
          - name: Test Go-diff
            id: test-go-diff
            uses: noh4nsen/go-diff@1.3.2
            with:
              head_branch: dev
              base_branch: main
          - name: Print output
            run: |
              echo ${{ steps.test-go-diff.outputs.full_output }}
              echo ${{ steps.test-go-diff.outputs.modified_files }}
              echo ${{ steps.test-go-diff.outputs.modified_dirs }}
              echo ${{ steps.test-go-diff.outputs.modified_projects }}
```

Outputs examples:

- full_output
```json
{"modifiedFiles":[".github/workflows/example.yml","build/go-diff-1.3.2"],"modifiedDirs":[".github/workflows","build"],"modifiedProjects":[".github","build"]}
```
- modified_files:
```json
{"files":[".github/workflows/example.yml","build/go-diff-1.3.2"]}
```
- modified_dirs:
```json
{"directories":[".github/workflows","build"]}
```
- modified_projects:
```json
{"projects":[".github","build"]}
```
### Inputs:

- head_branch
    - "Head branch for comparison. Only alphanumeric, dash and underscores are accepted."
    - required
    - default: dev
- base_branch
    - "Base branch for comparison. Only alphanumeric, dash and underscores are accepted."
    - required
    - default: main

### Outputs:

- full_output
  - Full JSON output generated by go-diff
- modified_files
  - JSON containing only modified files
- modified_dirs
  - JSON containing only modified directories
- modified_projects
  - JSON containing only directories at root level