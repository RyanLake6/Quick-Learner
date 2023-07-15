# Quick Learner

This service allows for the ingestion of a large text input and creates hyperlinks to wikipedia pages of key words found.

# Versions:

## V1 (Iterative):

This quick solution iteratively calls the wiki api causing for a slow response to return all wiki links mapped to found keywords

## V2 (Concurrent):

TODO: Using worker pools and channels concurrently hit the api within many threads to lower response time

TODO: Implement a queue service to hold all sent api requests as individual jobs that need to be completed within the pipeline

# Display:

Currently the display is just returning a map of all wiki links to their respective key words. TODO: planning on creating html to hyperlink each word as necessary for the user to be able to read their inputted text with the option to follow the wiki page to learn more

# Endpoints:

```
/v1/quickLearn
```

Pass in the text as a json body like the example below:

```
{
    input: "Here is your input text that will be hyperlinked"
}
```
