# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.1.2] - 2023-06-13
### Changed
- Fixed a bug that causes Scaffolder to fail if certain directories or files may not exist within a given template.

## [1.1.1] - 2023-06-10
### Changed
- No functional updates. Just fixed version info.

## [1.1.0] - 2023-06-07

### Changed
- Template tokens replacement have been updated to accommodate new templates:
  
  | Old Token         | New Token         |
  |-------------------|-------------------|
  | `{{module_name}}` | `MODULE_NAME`     |
  | `{{bin_name}}`    | `BIN_NAME`        |
  | `{{description}}` | `IMG_DESCRIPTION` |
  | `{{vendor_name}}` | `IMG_VENDOR`      |

  The old tokens are still supported for backwards compatibility with older templates.

- Updated default template to `https://github.com/twistingmercury/go-basic-tmpl.git`.

## [1.0.0] - 2024-05-22

### Added
- Initial release of Scaffolder, a command-line tool for creating Go projects from templates.
- `init` command to create a new project using a template from GitHub.
- Support for customizing the project using flags:
    - `--module-name` to set the Go module name.
    - `--bin-name` to set the name of the binary file.
    - `--template` to specify the GitHub repository to use as a template.
    - `--description` to provide a brief description of the project.
    - `--vendor-name` to specify the name of the vendor.
- Built-in template repository at `https://github.com/twistingmercury/gobasetmpl.git`.
- Template token replacement for `{{module_name}}`, `{{bin_name}}`, `{{description}}`, and `{{vendor_name}}`.
- README file with installation, usage, and contribution guidelines.

[1.1.2]: https://github.com/twistingmercury/scaffolder/compare/v1.1.1...v1.1.2
[1.1.1]: https://github.com/twistingmercury/scaffolder/compare/v1.1.0...v1.1.1
[1.1.0]: https://github.com/twistingmercury/scaffolder/compare/v1.0.0...v1.1.0
[1.0.0]: https://github.com/twistingmercury/scaffolder/releases/tag/v1.0.0
