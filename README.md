# simple-blog

The purpose of this project is to demonstrate how to use cleango to create a 
simple blog post application. 

## Routes
* GET /blogs - returns all of the blogs.
* GET /blogs/artifactID - returns a specific blog.
* GET /blogs/artifactID/edit - returns the edit page for a given blog.
* POST /blogs - creates a new blog post.
    if there is an issue with validation, need to redirect to the original 
  /blogs URL, this could be the /blogs/artifactID/edit or the original /blogs.