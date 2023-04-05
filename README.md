# Zettelo: The Hassle-free Zettelkasten

Zettelo is a simple command-line tool that helps you convert your Markdown files to a Zettelkasten knowledge base. The tool allows you to quickly organize your notes and research, while avoiding the hassle of manually sifting through your files for specific information.

With Zettelo, all you have to do is specify the folder containing your Markdown files, and the tool will automatically extract all relevant information and output them in a neatly organized table format.

Zettelo is highly configurable, allowing you to customize the tagging format and output, as well as add your own custom tag mappings.

Zettelo is built with Golang and is designed to be fast, efficient, and easy to use. Try it out today and take the first step towards effortless knowledge organization!

## An Opinionated Zettelkasten Approach

In this system, we use a set of predefined tags to create a well-structured and interconnected web of notes that reflects our thoughts and ideas.

At the top level, we have the `#note` tag, which is used for any kind of note you create in your zettelkasten system. Underneath that, we have several subtags that allow you to categorize your notes in different ways:

* `#reference`: For notes that contain reference material you want to keep in your zettelkasten system, such as articles, research papers, or books.
* `#task`: For notes that contain tasks or action items you need to complete.
* `#idea`: For notes that contain ideas or insights you want to explore further.
* `#question`: For notes that contain questions or topics you want to research or explore further.
* `#moc`: For notes that are part of a "map of content" you are creating in your zettelkasten system.
* `#project`: For notes that are related to a specific project you are working on.
* `#goal`: For notes that relate to your goals, objectives, or aspirations.
* `#log`: For notes that contain daily logs or journal entries.
* `#people`: For notes that relate to people you interact with, such as colleagues, friends, or family members.

By using these tags consistently, you can create a hierarchy of notes that allows you to quickly find and navigate related content. Additionally, you can use the `[[double bracket syntax]]` to create links between notes, allowing you to build a rich and interconnected knowledge base.


## Requirements

* Go 1.16 or later
* Git

## Usage

1. Clone the repository: `git clone https://github.com/ozcankasal/zettelo.git`
2. Navigate to the project directory: `cd your-repo`
3. Create a YAML configuration file with your settings. Here is an example `config.yaml` file (which is in `samples` folder). Add your folders (multiple folders support added now!)
   
  ```yaml
# Web server configuration
web:
  port: 8080
  host: localhost

# Application-specific settings
app:
  tag_mappings:
    #todo: #todo
    #to-do: #todo
    #todo: #todo
  folders:
    - /path/to/folder1
    - /path/to/folder2
  ```

4. Move this file to `~/.zettelo/config.yaml`

5. Retrieve the dependencies with `go get ./...`
3. Build the binary: `go build -o zettelo cmd/zettelo/zettelo.go`
4. Run the binary with the path to your markdown files directory as an argument: `./zettelo`
5. Open your web browser and go to `localhost:8080` to view tags and their corresponding file locations.

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

