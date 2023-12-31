# User Management System

## Overview

This project comprises a user management system designed to facilitate CRUD (Create, Read, Update, Delete) operations on user data. It consists of a frontend web application developed with TypeScript using React and Next.js, complemented by a backend service built with Go. The system's primary purpose is to provide users with a seamless interface for managing user data via a master view and a detailed view.

### Features

- **Master View:**
  - Displays users in a data grid.
  - Allows CRUD operations (New, Edit, Delete) via buttons.
  - Requires row selection for Edit and Delete actions.

- **Detailed View:**
  - Exhibits user details in a form format.
  - Contains "Action" and "Back" buttons.

- **RESTful API Endpoints:**
  - Follows RESTful conventions for API design.
  - Supports operations such as:
    - Retrieve all users.
    - Retrieve a specific user by ID.
    - Save a new user.
    - Update user data by ID.
    - Delete a user by ID.

- **Backend Implementation:**
  - Developed in Go language.
  - Utilizes SQLite for database management.
  - Database file included within the project folder for data persistence.

- **Frontend Implementation:**
  - Implemented using TypeScript, React, and Next.js.
  - Provides an intuitive user interface for managing user data.
  - Communicates with the backend through RESTful API endpoints.

### API Specifications

- **GET /api/getUsers:** Retrieves all users.
- **GET /api/getUser:** Retrieves a user by ID.
- **POST /api/saveUser:** Saves a new user.
- **PUT /api/updateUser:** Updates user data by ID.
- **DELETE /api/deleteUser:** Deletes a user by ID.

### Technologies Used

- **Frontend:** TypeScript, React, Next.js
- **Backend:** Go, SQLite

### Build Instructions

#### Frontend

1. **Install Dependencies:**
   ```bash
   cd UserManagement/frontend
   npm i
   ```
2. **Start the Frontend**
   ```bash
    npm run dev
   ```
### Backend
1. **Start the Backend**
    ```bash
    go run cmd/main.go
    ```
