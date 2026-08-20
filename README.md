# Blog API

A clean and simple RESTful API for managing a personal blog platform. Built with a focus on core CRUD operations, article tagging, filtering, and structured data handling.

---

## Features

- **Get All Articles**: Retrieve a list of all blog articles with support for query filtering (by publication date or tags).
- **Get Single Article**: Get detailed information for a specific article by its unique ID.
- **Create Article**: Publish new blog posts with title, content, and tags.
- **Update Article**: Modify existing articles by ID.
- **Delete Article**: Remove articles by ID.

---

## Tech Stack

- **Language**: Go
- **Router**: go-chi/chi
- **Database**: PostgreSQL

---

## Running the app
```bash
docker compose up -d
```
