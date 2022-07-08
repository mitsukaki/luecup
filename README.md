# luecup
Simple tag-oriented query engine to be used for Reverse Index searching. 

# Endpoints

## POST /api/fetch/{count}
Accepts a JSON object with the array child `tags` containing the tags to search for. Will return up to "count" results, sorted by how many tags an entry matched in descending order.

## GET /api/tags/{tag}
Gets all entries saved that are associated with the url-provided tag.

## PUT /api/tags/{tag}
Accepts a JSON array of objects that match the url-provided tag, to be enterd into the DB. Overwrites if the entry already exists.

## DELETE /api/tags/{tag}
Deletes the entry with the url-provided tag.