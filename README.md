# Go + React Calorie Tracker

A full-stack calorie tracker application built using Go for the backend and React for the frontend, with MongoDB as the database.

## Features

- CRUD operations for tracking meals and calories
- Interactive user interface for adding, editing, and deleting meal entries
- Responsive design for mobile and desktop usage
- Backend REST API with robust data validation

## Tech Stack

### Backend:
- **Go**: Handles the RESTful API and business logic
- **Gin Gonic**: Lightweight and fast web framework for Go
- **MongoDB**: NoSQL database for storing user and meal data

### Frontend:
- **React**: Builds the user interface
- **React Router**: Enables navigation between pages
- **Axios**: Handles HTTP requests to the backend

## Getting Started

Follow these instructions to set up and run the project locally.

### Prerequisites

- Go (v1.18+)
- Node.js (v16+)
- MongoDB (running instance or a cloud database like MongoDB Atla)

### Installation

#### Clone the Repository:
```bash
git clone https://github.com/razyneko/go-react-calorie-tracker.git
cd go-react-calorie-tracker
```

#### Backend Setup:
1. Navigate to the `backend` directory:
   ```bash
   cd backend
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Create a `.env` file for environment variables:
   ```
   MONGO_URI=your_mongodb_connection_string
   ```
4. Run the server:
   ```bash
   go run main.go
   ```

#### Frontend Setup:
1. Navigate to the `frontend` directory:
   ```bash
   cd ../frontend
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Start the development server:
   ```bash
   npm start
   ```

### Running the Application

- Access the application in your browser at `http://localhost:3000`.
- The backend runs by default at `http://localhost:8080`.

## API Endpoints

### Meals
- `GET /api/meals`: Retrieve all meals
- `POST /api/meals`: Add a new meal
- `PUT /api/meals/:id`: Update an existing meal
- `DELETE /api/meals/:id`: Delete a meal

## Screenshots

Add screenshots of your application to showcase its features.

## Contributing

Contributions are welcome! Feel free to submit a pull request or open an issue.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

---

Happy tracking! 🎉