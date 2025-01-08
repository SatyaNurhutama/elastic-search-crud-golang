# Elasticsearch CRUD with Golang and Gin

This project demonstrates how to implement a simple CRUD RESTful API for managing student records using Elasticsearch as the database, written in Golang with the Gin framework.

## Features

- **Create Student**: Add a new student document to Elasticsearch.
- **Retrieve Student**: Fetch a student document by ID.
- **Update Student**: Partially update a student document.
- **Delete Student**: Remove a student document by ID.

## Prerequisites

- Golang 1.20+
- Elasticsearch (tested with Elasticsearch 8.x)
- Docker (optional, for running Elasticsearch locally)

## Environment Variables

```
ELASTICSEARCH_ADDRESS=your-elastic-search-address(e.g., "https://localhost:9200")
```
```
ELASTICSEARCH_USERNAME=your-username
```
```
ELASTICSEARCH_PASSWORD=your-password
```
