# Setup

Update `ES_PASSWORD` and `DB_PASSWORD` in `.env` and `go run .`

# Server setup

1. clone this repo and monstache repo
2. build and move monstache binary to `/usr/bin`
3. update `ES_PASSWORD` and `DB_PASSWORD` in `.env`, `chmod +x setup.sh`, `./setup.sh`
4. build and move volunteery binary: `go build -o volunteery .`, `mv ./volunteery /usr/bin`
5. create systemd services for backend (`/usr/bin/volunteery`) and monstache (`/usr/bin/monstache`) and start them!

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
