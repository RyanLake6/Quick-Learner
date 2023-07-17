# Quick Learner

This service allows for the ingestion of a large text input and creates hyperlinks to wikipedia pages of key words found.

# Versions:

## V1 Naive Iterative Approach:

Version 1 is a quick solution that iteratively calls the wiki api causing for a slow response to return all wiki links mapped to found keywords

## V2 Multi-threaded Concurrent Approach:

Version 2 is a multi-threaded solution that creates worker pools and splits the large data ingestion into seperate jobs for the workers to handle. It attempts to optimize the number of workers and jobs based on the size of the data ingestion. This solution can be more than 10x faster than v1.

TODO: Implement a queue service to hold all sent api requests as individual jobs that need to be completed within the pipeline

# Display:

Currently the display is just returning a map of all wiki links to their respective key words.

TODO: planning on creating html to hyperlink each word as necessary for the user to be able to read their inputted text with the option to follow the wiki page to learn more

# Endpoints:

```
/v<versionNumber>/quickLearn
```

Pass in the text as a json body like the example below:

```
{
    input: "Here is your input text that will be hyperlinked"
}
```
