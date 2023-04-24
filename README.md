# twitter-account-scraper

## Description
This is a Golang-based tool that can be used to scrape tweets, metadata, images, and videos from a set of Twitter accounts. The tool requires configuration through a settings.json file. 

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Installation
To install and use the tool, follow these steps:
1. Clone the repository: `git clone https://github.com/TaylorGitRep/twitter-account-scraper.git`
2. Install the required dependencies: 
    - Go to the root of the repository: `cd twitter-account-scraper`
    - Install the dependencies: `go mod download`
3. Configure the settings.json file according to your needs. 
4. Run the tool with the following command: `go run main.go`

## Usage
The tool can be used to scrape tweets, metadata, images, and videos from a set of Twitter accounts. To use the tool:
1. Follow the installation steps mentioned above.
2. Configure the settings.json file according to your needs. You can specify the Twitter accounts to scrape, the type of data to scrape, and the directory to save the scraped data.
3. Run the tool with the following command: `go run main.go`
4. The tool will scrape the specified data from the specified Twitter accounts and save it in the specified directory.

## Contributing
If you are interested in contributing to the tool, please follow these steps:
1. Fork the repository.
2. Create a new branch for your changes: `git checkout -b my-feature-branch`
3. Make your changes and commit them: `git commit -am 'Add some feature'`
4. Push the changes to your fork: `git push origin my-feature-branch`
5. Create a pull request.

## License
This tool is licensed under the [MIT License](https://github.com/TaylorGitRep/twitter-account-scraper/LICENSE.md).
