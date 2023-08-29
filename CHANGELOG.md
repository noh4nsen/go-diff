<a name="unreleased"></a>
## [Unreleased]


<a name="1.5.2"></a>
## [1.5.2] - 2023-08-29
### Features
- **ISSUE-0004:** fetch and pull on git config step on build workflow


<a name="1.4.0"></a>
## [1.4.0] - 2023-08-09

<a name="1.3.2"></a>
## [1.3.2] - 2023-08-09
### Fixes
- typo on entrypoint script
- remove hardcoded path on entrypoint


<a name="1.3.1"></a>
## [1.3.1] - 2023-08-09
### Fixes
- go-diff binary on /usr/bin


<a name="1.3.0"></a>
## [1.3.0] - 2023-08-09
### Fixes
- remove unused log in git service
- added workspace as safe directory
- change git fetch command


<a name="1.2.6"></a>
## [1.2.6] - 2023-08-09
### Fixes
- binary path on script


<a name="1.2.5"></a>
## [1.2.5] - 2023-08-09
### Fixes
- update dockerfile and entrypoint with workspace volume context


<a name="1.2.4"></a>
## [1.2.4] - 2023-08-09

<a name="1.2.3"></a>
## [1.2.3] - 2023-08-09
### Fixes
- fetch repo on validation


<a name="1.2.2"></a>
## [1.2.2] - 2023-08-09
### Fixes
- remove branch name from pull command


<a name="1.2.1"></a>
## [1.2.1] - 2023-08-09
### Features
- pull branches on git service


<a name="1.2.0"></a>
## [1.2.0] - 2023-08-09
### Features
- exec branch checkout on git service


<a name="1.1.0"></a>
## [1.1.0] - 2023-08-09
### Features
- change logging, updated example, entrypoint script and action outputs description
- included full go-diff full json output as output option on action

### Fixes
- broken variable name on entrypoint


<a name="1.0.0"></a>
## [1.0.0] - 2023-08-08
### Features
- jq added as dependency on Dockerfile
- split outputs on action
- change action from composite to docker

### Fixes
- binary path on action


<a name="0.3.0"></a>
## [0.3.0] - 2023-08-08
### Features
- include project output on description
- included ModifiedProjects on output

### Fixes
- change project extraction logic


<a name="0.2.0"></a>
## [0.2.0] - 2023-08-07
### Features
- separate jobs on build, change input description on action

### Fixes
- reestructure workflow jobs


<a name="0.1.0"></a>
## [0.1.0] - 2023-08-07
### Features
- verify output value on action execution

### Fixes
- update branch error message


<a name="1.1.0-beta"></a>
## [1.1.0-beta] - 2023-08-07

<a name="1.0.0-beta"></a>
## 1.0.0-beta - 2023-08-07
### Features
- separated invalid args error messaging method
- new method that fetch specified branches
- separate directories verification
- new method to determine if branch exists
- created Makefile and compiled binaries
- created core application


[Unreleased]: https://github.com/noh4nsen/go-diff/compare/1.5.2...HEAD
[1.5.2]: https://github.com/noh4nsen/go-diff/compare/1.4.0...1.5.2
[1.4.0]: https://github.com/noh4nsen/go-diff/compare/1.3.2...1.4.0
[1.3.2]: https://github.com/noh4nsen/go-diff/compare/1.3.1...1.3.2
[1.3.1]: https://github.com/noh4nsen/go-diff/compare/1.3.0...1.3.1
[1.3.0]: https://github.com/noh4nsen/go-diff/compare/1.2.6...1.3.0
[1.2.6]: https://github.com/noh4nsen/go-diff/compare/1.2.5...1.2.6
[1.2.5]: https://github.com/noh4nsen/go-diff/compare/1.2.4...1.2.5
[1.2.4]: https://github.com/noh4nsen/go-diff/compare/1.2.3...1.2.4
[1.2.3]: https://github.com/noh4nsen/go-diff/compare/1.2.2...1.2.3
[1.2.2]: https://github.com/noh4nsen/go-diff/compare/1.2.1...1.2.2
[1.2.1]: https://github.com/noh4nsen/go-diff/compare/1.2.0...1.2.1
[1.2.0]: https://github.com/noh4nsen/go-diff/compare/1.1.0...1.2.0
[1.1.0]: https://github.com/noh4nsen/go-diff/compare/1.0.0...1.1.0
[1.0.0]: https://github.com/noh4nsen/go-diff/compare/0.3.0...1.0.0
[0.3.0]: https://github.com/noh4nsen/go-diff/compare/0.2.0...0.3.0
[0.2.0]: https://github.com/noh4nsen/go-diff/compare/0.1.0...0.2.0
[0.1.0]: https://github.com/noh4nsen/go-diff/compare/1.1.0-beta...0.1.0
[1.1.0-beta]: https://github.com/noh4nsen/go-diff/compare/1.0.0-beta...1.1.0-beta
