# Console-Based Library Management System in Go

## Overview

A simple console-based library management system built in Go.
Demonstrates:

* Structs (`Book` and `Member`)
* Interfaces (`LibraryManager`)
* Methods, slices, maps, and basic error handling

Users can add/remove books, borrow/return books, and list available or borrowed books.

---

## Features

* Add and remove books
* Borrow and return books
* List available books
* List borrowed books by member

---

## Folder Structure

```
BACKEND_GO/
├── main.go
├── controllers/
│   └── library_controller.go
├── models/
│   ├── book.go
│   └── member.go
├── services/
│   └── library_service.go
├── docs/
│   └── documentation.md
└── go.mod
```

---

## Usage

1. Clone the repository
2. Run:

```bash
go run main.go
```

3. Use the console menu to manage the library

---
