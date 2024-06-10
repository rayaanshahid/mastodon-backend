# Mastodon Live Stream Application

## Overview

This project streams public Mastodon posts in real-time and displays them on a web application.

## Technologies Used

- Backend: Go
- Frontend: React, TypeScript
- Containerization: Docker
- Communication: WebSocket

## Setup Instructions

1. Clone the repository:

git clone https://github.com/yourusername/mastodon-live-stream.git cd mastodon-live-stream

2. Set up environment variables:

- Create a `.env` file in the backend directory and add your Mastodon access token:
  ```
  MASTODON_ACCESS_TOKEN=your_access_token_here
  ```

3. Build and run the application using Docker Compose:

docker-compose up --build

4. Open your browser and navigate to `http://localhost:3000` to see the live Mastodon posts.

## Running Tests

To run the backend tests, use the following command:

docker-compose run backend go test ./...

## Architecture

The application consists of two main services:

- **Backend**: Connects to the Mastodon API and streams public posts using WebSocket.
- **Frontend**: Displays the live posts received from the backend WebSocket.

Both services are containerized using Docker and orchestrated with Docker Compose.

## CI/CD

Continuous Integration is set up using GitHub Actions to automatically build and test the application on every push to the main branch.

## Contributing

Feel free to open issues or submit pull requests if you have suggestions for improvements or bug fixes.
