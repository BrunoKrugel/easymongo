<p align="center">
<h1 align="center">Easy Mongo</h1>
<p align="center">Simple wraper for MongoDB</p>
</p>

## Features

## Usage

The following samples will assist you to become as comfortable as possible with easymongo library.

```go
// Import easymongo into your code and refer it as `easymongo`.
import "github.com/BrunoKrugel/easymongo"
```

#### Create Client

```go
easyMongo.NewMongoInstance("uri", "db", "collection")
```

#### Simple FindOne

```go
filter := bson.D{
    {Key: "id", Value: "123"},
}

easyMongo.NewMongoInstance("uri", "db", "collection").FindOne(filter)
```
