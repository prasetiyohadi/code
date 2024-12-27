# Web Service Gin

## API Endpoints Design

You’ll build an API that provides access to a store selling vintage recordings on vinyl. So you’ll need to provide endpoints through which a client can get and add albums for users.

When developing an API, you typically begin by designing the endpoints. Your API’s users will have more success if the endpoints are easy to understand.

Here are the endpoints you’ll create in this tutorial.

`/albums`

 - GET – Get a list of all albums, returned as JSON.
 - POST – Add a new album from request data sent as JSON.

`/albums/:id`

 - GET – Get an album by its ID, returning the album data as JSON.

## Data Storage

To keep things simple for the tutorial, you’ll store data in memory. A more typical API would interact with a database.

Note that storing data in memory means that the set of albums will be lost each time you stop the server, then recreated when you start it.

## Handlers

### Return all items

When the client makes a request at GET /albums, you want to return all the albums as JSON.

To do this, you’ll write the following:

 - Logic to prepare a response
 - Code to map the request path to your logic

Note that this is the reverse of how they’ll be executed at runtime, but you’re adding dependencies first, then the code that depends on them.

### Add a new item

When the client makes a POST request at /albums, you want to add the album described in the request body to the existing albums’ data.

To do this, you’ll write the following:

 - Logic to add the new album to the existing list.
 - A bit of code to route the POST request to your logic.

### Return a specific item

When the client makes a request to GET /albums/[id], you want to return the album whose ID matches the id path parameter.

To do this, you will:

 - Add logic to retrieve the requested album.
 - Map the path to the logic.
