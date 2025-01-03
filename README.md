# GCMS (GitHub Content Management System) 🗃️
GCMS is a command-line interface (CLI) tool designed to manage content on GitHub repositories efficiently. It provides a set of commands to initialize, configure, and manage content through github, making it easier to handle content directly from the terminal.

[![Tech Stack](https://skillicons.dev/icons?i=go)]()

## Features
- **Initialize Repositories**: 
  - `gcms init --empty`: Initializes an empty repository and sets it as the GCMS repository.
  - `gcms init --from <github_link>`: Clones an existing repository and configures it as the GCMS repository.
- **Repository Information**:
  - `gcms info repo`: Displays information about the local GCMS repository.
  - `gcms info remote`: Shows details about the remote GCMS repository.
- **Detach Repositories**:
  - `gcms detach`: Soft detaches the local repository from the remote.
  - `gcms detach --hard`: Hard detaches and deletes the remote repository, given sufficient permissions.
- **File Management**:
  - `gcms push <filename>`: Pushes a file to the repository.
  - `gcms push (-g / --group) <group_name>`: Pushes files grouped under a specific name.
  - `gcms push (-m / --metadata) <metadata_key>:<metadata_value>`: Adds metadata to pushed files.
  - `gcms list`: Lists all files in the local repository.
  - `gcms remove <filename>`: Removes files from both the local and remote repositories.
  - `gcms update`: Pulls from remote and updates
 
- **Health**
  -  `gcms doctor`: Checks health of the installation

## Commands
| Command | Description |
|---------|-------------|
| `gcms` | Show menu |
| `gcms config set <key> <value>` | Sets the configuration value (For example: github.pat_token required for github auth) |
| `gcms config get <key>` | Gets the configuration value |
| `gcms init --from <github_link>` | Clones the specified repository and sets it as the GCMS repository |
| `gcms init --empty` | Initializes an empty repository and sets it as the GCMS repository |
| `gcms info repo` | Gets information about the local GCMS repository |
| `gcms info remote` | Gets information about the remote GCMS repository |
| `gcms detach` | Soft detaches the remote repository, removing only the remotes |
| `gcms detach --hard` | Hard detaches and deletes the remote repository, given sufficient permissions |
| `gcms delete-local` | [WARN] Deletes the local repository |
| `gcms push <filename>` | Pushes the specified file to the repository |
| `gcms push (-g / --group) <group_name>` | Pushes files grouped under a specific name |
| `gcms push (-m / --metadata) <metadata_key>:<metadata_value>` | Adds metadata to pushed files |
| `gcms list` | Lists all files in the local repository |
| `gcms remove <filename>` | Removes the specified file from both the local and remote repositories |
| `gcms update` | Pulls from remote and updates |
| `gcms doctor` | Checks health of the installation

## Installation
There are two primary methods to install GCMS:

### Method 1: Building from Source
Clone the repository and build the CLI:
```bash
git clone https://github.com/saphalpdyl/gcms.git
cd gcms
go build -o gcms
```

### Method 2: Go Install
Use Go's installation method to directly install the CLI:
```bash
go install github.com/saphalpdyl/gcms
```

After installation, you can run GCMS directly from your command line.

## Usage
After installing, you can run commands using:
```bash
gcms <command> [flags] [arguments]
```

### Example Usage
Initialize an empty repository:
```bash
gcms init --empty
```

Push a file with grouping and metadata:
```bash
gcms push -g blog -m tags=golang;cli=mypost.html
```

## Contributing
Contributions are welcome! Please fork the repository and submit a pull request with your changes. Ensure that your code follows the project's coding standards and includes relevant tests.

## License
This project is licensed under the MIT License. See the `LICENSE` file for more details.

## Support
For any issues or questions, please open an issue on the GitHub repository or contact the maintainers directly.
