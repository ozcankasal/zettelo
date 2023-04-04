# Zettelo: The Hassle-free Hashtag Extractor

Zettelo is a simple command-line tool that helps you extract hashtags from your Markdown files and export them as JSON data. The tool allows you to quickly organize your notes and research, while avoiding the hassle of manually sifting through your files for specific hashtags.

With Zettelo, all you have to do is specify the folder containing your Markdown files, and the tool will automatically extract all hashtags and output them in a neatly organized JSON format.

Zettelo is highly configurable, allowing you to customize the tagging format and output, as well as add your own custom tag mappings.

Zettelo is built with Golang and is designed to be fast, efficient, and easy to use. Try it out today and take the first step towards effortless note-taking and research organization!

## Requirements

* Go 1.16 or later
* Git

## Steps

1. Clone the repository: `git clone https://github.com/ozcankasal/zettelo.git`
2. Navigate to the project directory: `cd your-repo`
3. Create a YAML configuration file with your tag mappings. Here is an example `config.yaml` file (which is in `samples` folder)
   
  ```
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
