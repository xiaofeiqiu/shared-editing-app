# Collaborative Text Editor with WebSocket

This project is a collaborative text editor built using React, Yjs, and Quill. It allows multiple users to edit text in real-time, with synchronization handled via a WebSocket server implemented in Go.

## Features

- Real-time collaborative text editing
- WebSocket-based synchronization
- Peer-to-peer and local synchronization disabled for server-only syncing

## Prerequisites

- [Go](https://golang.org/doc/install) (1.16 or later)
- [Node.js and npm](https://nodejs.org/en/download/) (Node 14.x or later)

## Setup

### Backend (Go)

1. Clone the repository:

   ```sh
   git clone https://github.com/your-username/collaborative-text-editor.git
   cd collaborative-text-editor

2. Navigate to the "server" directory:
  cd server

3. Run the Go server:
  go run main.go

The server will start on "http://localhost:8080".

Frontend (React):
1. Navigate to the "client" directory:
  cd client

2. Install dependencies:
  npm install

3. Build the React application:
  npm run build

4. Serve the built files using a static server (e.g., "serve"):
  npm install -g serve
  serve -s build
