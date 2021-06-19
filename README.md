# Setup 

Update password in `db.go` and `go run .`


# Create

## POST
```
/users/_doc
/admins/_doc
/orgs/_doc
/opps/_doc
```

# Read

## GET

```
/users/_search
/admins/_search
/orgs/_search
/opps/_search

// basic text search
{
    "query": {
        "match": {
            "key": "value"
        }
    }
}

// search phrase
{
    "query": {
        "match_phrase": {
            "key": "a phrase"
        }
    }
}

// complex queries
{
    "query": {
        "bool": {
            "must": [
                "match": {
                    "key": "value"
                },
                 "match_phrase": {
                    "key": "a phrase"
                }       
            ]
        }
    }
}

// also have must_not for negations, weighted searches based on certain queries, gte (range instead of bool), sort, etc.

// you can even have it highlight the keyword in the text automatically (sends it back as response)
```

# Update

## PUT

```
/users/_doc/:id
/admins/_doc/:id
/orgs/_doc/:id
/opps/_doc/:id

// bulk updates
POST /users/_bulk
{
    { "index": { "_id": 0 } }
    {
        "updates"....
    }
    { "index": { "_id": 1 } }
    {
        "updates"....
    }
    ....
}
```

# Delete

## DELETE

```
/users
/admins
/orgs
/opps
```
