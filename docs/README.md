# What is this project about?

Zettelkasten is a note-taking and knowledge management system that was first developed by German sociologist Niklas Luhmann. The word "Zettelkasten" means "slip box" in German, and the system involves creating individual notes on small slips of paper, which are then organized and interconnected in a physical box.

In recent years, the Zettelkasten method has been adapted to digital note-taking using various software tools. One popular approach is to use plain text files, often written in Markdown syntax, to create and manage notes. This has the advantage of being simple, flexible, and easily searchable.

My aim is to develop a tool that makes it easy to create and manage a Zettelkasten using Markdown files. Specifically, given a list of folders, the tool will scan Markdown files and extract information from the headers to create a knowledge graph of interconnected notes. The header will contain fields like "id", "type", "project", and "tags", which will be used to organize and search the notes.

The "id" field will be a unique identifier for each note, while the "type" field will indicate the type of note (e.g. concept, definition, reflection). The "project" field will indicate the project or context to which the note belongs, while the "tags" field will contain additional metadata like keywords or hashtags.

The tool will be useful for a variety of use cases, including academic research, creative writing, and personal knowledge management. For example, a researcher could use the tool to create a knowledge base of related ideas and concepts, while a writer could use it to organize notes for a book or article. Users can also add additional features to the tool, such as a "todo" tag for notes that require follow-up action, or a "review" tag for notes that need to be revisited periodically.

## Markdown Headers 

The use of Markdown headers to create a standardized format for notes in a Zettelkasten system can be a powerful tool for organizing and managing knowledge. By creating a consistent header format for each note, users can easily sort, search, and link related notes together. This approach allows for greater flexibility in creating notes while still maintaining a standardized structure that can be used for categorization and searchability. Additionally, using a text-based format like Markdown allows for easy version control and portability across different devices and platforms. By adopting this approach, users can create a comprehensive and interconnected knowledge base that can be easily searched and accessed whenever needed.

Sure, here's an example of a standardized Markdown header:

```
id: e9eb51f7-0706-4eb8-a343-6c0c7f4f6e4d
type: concept
project: maths-book-writing
tags: algebra, equations, proof
```

In this example, the header contains four fields:

"id": a unique identifier for the note, generated using the "github.com/google/uuid" package.
"type": a description of the type of note, in this case "concept".
"project": the project or context to which the note belongs, in this case "maths-book-writing".
"tags": additional metadata that can be used to categorize or search the note, in this case "algebra", "equations", and "proof".
Using a standardized header like this can help to maintain consistency and make it easier to organize and search your notes. It's also flexible enough to accommodate different types of notes and projects, and can be customized as needed.

When creating your own notes, be sure to follow the same header format to ensure consistency across your zettelkasten. You can use any valid UUID for the "id" field, and create your own list of tags based on your needs. Additionally, you may want to use a text editor or note-taking app that supports Markdown syntax, to make it easy to create and edit notes.

## Supported Types

In addition to the fields included in the standardized Markdown header, I also suggest using a standard set of "type" categories to further organize and categorize notes within a Zettelkasten system. By using a set of pre-defined types, users can ensure that their notes are consistently categorized and easily searchable. Here is the list of types that Zettelo supports.

* Concept
* Definition
* Example
* Experiment
* Fact
* Idea
* Insight
* Note
* Observation
* Quote
* Reflection
* Summary
* Theory
* Miscellaneous

## Hashtags to extract info

In addition to the metadata included in the Markdown header, you can also use hashtags to add additional context and structure to your notes. Hashtags can be used to group related notes together, highlight important information, and even create to-do lists. By adding hashtags to individual lines within a note, you can quickly and easily filter and search through your notes to find relevant information.

For example, you might use the hashtag "#todo" to indicate items that need to be done, and add a due date using the hashtag "#due" followed by a date. This can help you to keep track of your tasks and ensure that you're meeting important deadlines. Other hashtags that you might find useful include "#important" to indicate particularly critical information, "#question" to flag notes that require further research or clarification, and "#reference" to highlight notes that provide background information for a project or concept.

Here are ten possible hashtags and their uses:

* "#todo": Use this hashtag to indicate tasks that need to be completed.
* "#due": Use this hashtag to indicate a due date for a task.
* "#important": Use this hashtag to highlight particularly critical information.
* "#question": Use this hashtag to flag notes that require further research or clarification.
* "#reference": Use this hashtag to highlight notes that provide background information for a project or concept.
* "#idea": Use this hashtag to indicate ideas that you want to explore further.
* "#contact": Use this hashtag to indicate notes related to a particular contact or person.
* "#location": Use this hashtag to indicate notes related to a particular location.
* "#event": Use this hashtag to indicate notes related to a particular event or meeting.
* "#note": Use this hashtag to indicate general notes that don't fit into a specific category.


## Inter-file links

Inter-file links are an important feature of a Zettelkasten system that allow you to link related notes together and create a network of information. By using inter-file links, you can establish connections between notes that might not be immediately obvious, allowing you to see patterns and relationships that might have otherwise gone unnoticed. In a Zettelkasten system, each note should have a unique identifier, and inter-file links can be created by simply including the identifier for the target note in double square brackets like this: [identifier].

For example, if you have a note with the identifier "1234" that relates to a concept you've discussed in another note with the file name "file-name.md" and the identifier "5678", you could create a link to that note by adding `[link text](/path-to/file-name.md?id=5678)` to the first note. This creates a bidirectional link between the two notes, allowing you to quickly navigate between them.

In addition to creating links between related notes, inter-file links can also be used to create a graph of your notes. By visualizing the connections between different notes, you can gain a better understanding of the relationships between different ideas and concepts. This can be especially useful when working on complex projects or research topics.

There are a number of tools available for visualizing the connections between notes in a Zettelkasten system, including graph visualization software like Gephi or yEd. By exporting your notes in a standardized format, such as GraphML or GML, you can create a visual representation of your knowledge network that can be used to explore your ideas in a more intuitive way.

Some possible uses of inter-file links and graph visualization include:

* Identifying patterns and relationships between different concepts and ideas.
* Identifying gaps in your knowledge and areas that require further research.
* Exploring the connections between different projects and tasks.
* Visualizing the progress of a project over time.
* Identifying areas for collaboration or cross-disciplinary research.

Ultimately, the use of inter-file links and graph visualization is intended to help you manage and organize your knowledge in a more intuitive and interconnected way. By using these tools, you can create a more comprehensive and useful knowledge base that can help you to tackle complex problems and develop new ideas.