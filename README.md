# React Calorie Tracker

A **React-based** web application that helps users track their daily calorie intake by logging food dishes, ingredients, calories, and fat content. The application provides an interactive UI for users to input and manage their calorie entries. Data is stored in **MongoDB** for persistence and the app provides options to add, update, and delete entries.

## Features

- **Track Calories**: Users can input calorie information for each dish they consume.
- **Add/Edit/Delete Entries**: Easily manage calorie data entries.
- **Ingredient Management**: Users can update ingredients for their entries.
- **Real-time Updates**: The app fetches and updates calorie entries dynamically.
- **Responsive UI**: Built using **React** and styled with **Bootstrap** for an intuitive user experience.

## Technologies Used

- **Frontend**: React, Axios, React-Bootstrap
- **Backend**: Go (with Gin Gonic) - API Server (optional, if using a full-stack version)
- **Database**: MongoDB (optional, if using a full-stack version)

## Installation

### Prerequisites

- **Node.js**: Ensure Node.js is installed. If not, install it from [the official website](https://nodejs.org/).
- **MongoDB** (if using the full-stack version): Install MongoDB and run it locally, or use a cloud instance such as MongoDB Atlas.

### Steps to Set Up

1. Clone the repository:
   ```bash
   git clone <repository_url>
   ```

2. Navigate to the project directory:
   ```bash
   cd react-calorie-tracker
   ```

3. Install dependencies:
   ```bash
   npm install
   ```

4. Set up MongoDB connection in your backend server (if applicable):
   - The backend is optional for a fully functional app. If you want to store the entries in MongoDB, make sure to configure the connection in the server.
   - You can use a local instance of MongoDB or MongoDB Atlas (cloud version).

5. Run the React app:
   ```bash
   npm start
   ```
   This will start the app on `http://localhost:3000`.

## Folder Structure

- **/src**
  - **/components**: Contains all the React components for managing the UI, such as `Entries.js`, `SingleEntry.js`.
  - **/services**: Contains functions for making HTTP requests to the API server (if applicable).
  - **/App.js**: The main entry point for the React app.

## How to Use

### 1. **Add a New Entry**

- Click the **Track today's calories** button.
- Fill out the form with the dish name, ingredients, calories, and fat content.
- Click **Add** to save the entry to the database.

### 2. **View Entries**

- All the tracked entries are displayed in a list, showing the dish name, ingredients, calories, and fat.
- You can click on an entry to update or delete it.

### 3. **Edit an Entry**

- Click **Change Entry** on any entry to edit its details.
- Update the dish, ingredients, calories, or fat, then click **Change** to save the updates.

### 4. **Delete an Entry**

- Click **Delete Entry** on any entry to remove it from the list.

### 5. **Change Ingredients**

- Click **Change Ingredients** on any entry to update the ingredients for that dish.
- Enter the new ingredients and click **Add** to save.

## API Endpoints (If Using a Full-Stack Backend)

The following API endpoints are available for interaction with the database (if using the Go backend):

### `GET /entries`
- Fetches all calorie entries.

### `POST /entry/create`
- Creates a new calorie entry.
- Request body:
  ```json
  {
    "dish": "Dish Name",
    "ingredients": "Ingredients List",
    "calories": 200,
    "fat": 5
  }
  ```

### `PUT /entry/update/:id`
- Updates an existing calorie entry by ID.
- Request body:
  ```json
  {
    "dish": "Updated Dish Name",
    "ingredients": "Updated Ingredients List",
    "calories": 250,
    "fat": 10
  }
  ```

### `DELETE /entry/delete/:id`
- Deletes an entry by ID.

### `PUT /ingredient/update/:id`
- Updates the ingredients for a specific entry by ID.
- Request body:
  ```json
  {
    "ingredients": "Updated Ingredients"
  }
  ```

## Styling

The app is styled using **React-Bootstrap**, which provides responsive and accessible UI components. You can customize the styles further as per your requirements.

## Testing

You can use tools like **Postman** or **cURL** to interact with the backend API for testing purposes.

Example cURL commands:

- **Create a New Entry**:
  ```bash
  curl -X POST http://localhost:8000/entry/create -d '{"dish": "Pizza", "ingredients": "Cheese, Sauce, Dough", "calories": 300, "fat": 15}' -H "Content-Type: application/json"
  ```

- **Update an Entry**:
  ```bash
  curl -X PUT http://localhost:8000/entry/update/12345 -d '{"dish": "Burger", "ingredients": "Beef, Lettuce, Tomato", "calories": 400, "fat": 20}' -H "Content-Type: application/json"
  ```

- **Delete an Entry**:
  ```bash
  curl -X DELETE http://localhost:8000/entry/delete/12345
  ```
