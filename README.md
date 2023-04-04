# Zettelo: The Hassle-free Zettelkasten

Zettelo is a simple command-line tool that helps you convert your Markdown files to a Zettelkasten knowledge base. The tool allows you to quickly organize your notes and research, while avoiding the hassle of manually sifting through your files for specific information.

With Zettelo, all you have to do is specify the folder containing your Markdown files, and the tool will automatically extract all relevant information and output them in a neatly organized table format.

Zettelo is highly configurable, allowing you to customize the tagging format and output, as well as add your own custom tag mappings.

Zettelo is built with Golang and is designed to be fast, efficient, and easy to use. Try it out today and take the first step towards effortless knowledge organization!

## Requirements

* Go 1.16 or later
* Git

## Usage

1. Clone the repository: `git clone https://github.com/ozcankasal/zettelo.git`
2. Navigate to the project directory: `cd your-repo`
3. Create a YAML configuration file with your tag mappings. Here is an example `config.yaml` file (which is in `samples` folder)
   
  ```yaml
  tag_mappings:
    "#todo": todo
    "#to-do": todo
    "#todo:": todo
  ```

4. Export the path to the YAML configuration file as an environment variable. 
This tells the program where to find the configuration file.
```
export ZETTELO_CONFIG=/path/to/config.yaml
```

4. Retrieve the dependencies with `go get ./...`
5. Build the binary: `go build -o zettelo cmd/zettelo/zettelo.go`
6. Run the binary with the path to your markdown files directory as an argument: `./zettelo /path/to/your/files`
7. Open your web browser and go to `localhost:8080` to view tags and their corresponding file locations.

The output will be a table with the following format:

|Tag|File|Content
|---|---|---|
|tag1|/path/to/file1.md|Related text with #tag1|
|tag2|/path/to/file2.md|Related text with #tag2|

## Realtime Updates

Zettelo supports realtime updates using websockets. When the app is running, it will serve the output on localhost:8080. Anytime a file in the specified directory is updated, added or deleted, the output table will automatically update in your browser.

## Configuration

Zettelo is configurable via a YAML configuration file. To use a custom configuration, set the ZETTELO_CONFIG environment variable to the path of the YAML file.

Here's an example YAML configuration file:

```yaml
tag_mappings:
  "#todo": "todo"
  "#to-do": "todo"
  "#todo:": "todo"
```


## Features

This project is in early days and most of the intended features are missing. Currently, the following features are available:

* Extract hashtags from Markdown files
* Tag them with custom tag mappings
* Output them in a table format
* Serve the output as a webpage on localhost:8080

## Notice

This repository is still in development and may contain bugs or missing features. Please use with caution and report any issues or feature requests to the repository's issue tracker.

