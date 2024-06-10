# Scaffolder

Scaffolder is a command-line tool that allows you to quickly create new Go projects based on templates hosted on GitHub. It simplifies the process of setting up a new project by cloning a template repository and replacing template tokens with user-provided values.

## Installation

To install Scaffolder, make sure you have Go installed on your system. Use the following command to install the tool:

```bash
go install github.com/twistingmercury/scaffolder@v1.1.1
```

## Usage

| Available Commands | Description                                                         |
|--------------------|---------------------------------------------------------------------|
| completion         | Generate the autocompletion script for the specified shell          |
| help               | Help about any command                                              |
| init               | Initializes a new project using a Go template project from GitHub.  |
| version            | Returns the current scaffolder version                              |


## init Flags

To create a new project using Scaffolder, use the `init` command followed by the required flags:
```
scaffolder init -m <module-name> -b <bin-name> [-t <template-url>] [-d <description>] [-v <vendor-name>]
```

#

| Flag                | Required | DescriptionKey                                                                                      |
|---------------------|----------|-----------------------------------------------------------------------------------------------------|
| `-m, --module-name` | Yes      | The name of the Go module (e.g., github.com/username/project-name) as set in the `go.mod` file.     |
| `-b, --bin-name`    | Yes      | The name of the binary file to be compiled. It sets the service name in the `conf.go` file.         |
| `-t, --template`    | No       | The URL of the GitHub repository to clone and use as a template. Defaults to [go-basic-tmpl](https://github.com/twistingmercury/go-basic-tmpl) |
| `-d, --description` | No       | A brief description of the project, set in the Dockerfile as a label.                               |
| `-v, --vendor-name` | No       | The name of the vendor, set in the Dockerfile as a label.                                           |
| `-h, --help`        | No       | Display help information for the `init` command.                                                    |

### Example

To create a new project using the default template, you can run the following command:

```bash
scaffolder init -m github.com/username/myproject -b "myproject" -d "My awesome project" -v "John Doe"
```

This will clone the default template repository, replace the template tokens with the provided values, and set up the project for development.

## Template Repositories

Scaffolder uses GitHub repositories as templates for creating new projects. The template repository should contain the necessary files and directory structure for a Go project. The files can include template tokens that will be replaced with user-provided values during project creation.

The default template repository is `https://github.com/twistingmercury/gobasetmpl.git`, but you can specify a different template using the `-t` flag.

## Template Tokens

Scaffolder supports the following template tokens that can be used in the template files:

| Token              | DescriptionKey                                         |
|--------------------|--------------------------------------------------------|
| `MODULE_NAME`      | Represents the name of the Go module.                  |
| `BIN_NAME`         | Represents the name of the binary file to be compiled. |
| `IMG_DESCCRIPTION` | Represents the description of the project.             |
| `IMG_VENDOR_NAME`  | Represents the name of the vendor.                     |

These tokens will be replaced with the corresponding user-provided values during project creation.

## Contributing

Contributions to Scaffolder are welcome! If you find any issues or have suggestions for improvement, please open an issue or submit a pull request on the [GitHub repository](https://github.com/twistingmercury/scaffolder).

## License

Scaffolder is open-source software released under the [MIT License](./LICENSE).

