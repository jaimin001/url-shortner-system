# URL Shortener

Hey there! Thanks for checking out this project. 

This is a simple URL shortener service I put together. It takes long, messy URLs and turns them into short, easy-to-share links. I built this to learn how these things work under the hood.

## How it works
The system is split into two main parts:
- **Backend:** Written in Go, it handles the requests, talks to the database, and keeps everything running. It uses MongoDB for storage and a caching layer for speed.
- **Frontend:** A clean React interface built with Vite and Tailwind CSS. It lets you quickly shorten your URLs and see your recent links.

## Getting Started

### Prerequisites
You'll need Go and Node.js installed on your machine.

### Running the Backend
1. Head over to the `src/backend` folder.
2. Make sure you have your database running.
3. Run the server:
   ```bash
   go run main.go
   ```

### Running the Frontend
1. Go into the `frontend` folder.
2. Install the packages:
   ```bash
   npm install
   ```
3. Start the dev server:
   ```bash
   npm run dev
   ```

## Built With
- **Language:** Go
- **Frontend:** React, TypeScript, Tailwind CSS
- **Database:** MongoDB
- **Caching:** Valkey (Redis-compatible)

## Any Questions?
I'm still tinkering with this, so feel free to look around! If you see something that could be better or just want to say hi, feel free to open an issue or reach out. Thanks again for stopping by!
