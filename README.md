# Exploding Kittens Card Game 😼


![ReactJS](https://img.shields.io/badge/-ReactJs-61DAFB?logo=react&logoColor=white)
![Redux](https://img.shields.io/badge/-Redux-764ABC?logo=redux&logoColor=white)
![GoLang](https://img.shields.io/badge/-GoLang-00ADD8?logo=go&logoColor=white)
![Redis](https://img.shields.io/badge/-Redis-DC382D?logo=redis&logoColor=white)
![Gorilla Mux](https://img.shields.io/badge/-Gorilla%20Mux-00ADD8?logo=go&logoColor=white)



Welcome to the Exploding Kittens card game! This is an online single-player card game where you try to draw cards from a deck without exploding kittens blowing up in your face. Let's dive into how the game works and how you can access it locally.

## Working

  <video controls height="130">
  <source src="Screenshots/gameFilm.mp4" type="video/quicktime">
    Sorry! Your browser does not support the video tag.
  </video>


## How the Game Works

### Rules:
- **Types of Cards:**
  - 😼 **Cat Card**: Harmless card, just a cute cat.
  - 🙅‍♂️ **Defuse Card**: Used to defuse an exploding kitten.
  - 🔀 **Shuffle Card**: Shuffles the deck and resets the game.
  - 💣 **Exploding Kitten Card**: Instant game over if drawn.

- **Gameplay:**
  - Clicking on the deck reveals a card.
  - If it's a cat card, it's removed from the deck.
  - If it's an exploding kitten card and you don't have a defuse card, you lose.
  - If it's a defuse card, it's removed from the deck and can be used to defuse an exploding kitten later.
  - If it's a shuffle card, the game restarts with a freshly shuffled deck.
  - Win the game by drawing all non-exploding kitten cards.

## Accessing the Game Locally

### File Directory of the Game
```
   |--frontend
   |   |-- node_modules/
   |   |-- public/
   |   |   |-- index.html
   |   |   |-- cat.ico
   |   |   |-- manifest.json
   |   |-- src/
   |   |   |-- components/
   |   |   |   |-- CardBlock.js
   |   |   |   |-- Info.js
   |   |   |   |-- LeaderBoard.js
   |   |   |   |--NameForm.js
   |   |   |   |-- Navbar.js
   |   |   |   |-- Toast.js
   |   |   |-- assets/
   |   |   |   |-- card_back.png
   |   |   |   |-- card_border.png
   |   |   |   |-- pixx.png
   |   |   | 
   |   |   |-- App.js
   |   |   |-- App.css
   |   |   |-- index.js
   |   |   |-- index.css
   |   |-- package.json
   |   |-- package-lock.json
   | 
   |--backend
   |   |--go.mod
   |   |--go.sum
   |   |--main.go
```

### For Frontend (React):
1. Install dependencies: `npm install`
2. Start the frontend server: `npm start`
3. Access the game at: `http://localhost:3000`

### For Backend (Go Lang):
1. Install dependencies: `go mod tidy`
2. Build the Go server: `go build -o main`
3. Start the server: `go run main.go`

Sure, here are the steps to set up Redis for your project along with some example code snippets:

### For Redis DB (Backend):

#### Step 1: Install Redis

First, you need to install Redis on your system. You can download and install Redis from the official website or use package managers like `apt` or `brew` depending on your operating system.

#### Step 2: Start Redis Server

Start the Redis server on your local machine. You can do this by running the following command in your terminal:

```bash
redis-server
redis-server-start
```

This command will start the Redis server with default configurations.


#### Step 3: Close Redis Connection

Don't forget to close the Redis connection when you're done:

```bash
redis-server-stop
```
### Environment Variables:
That's it! You've successfully set up Redis for your project and connected to it using Go Lang. Now you can use Redis to store and retrieve data as needed.
This project requires certain environment variables to be configured. Since the .env file is included in .gitignore, it is not present in the repository. After cloning the project, create a .env file in the root directory of the project with the following keys:

```
REDIS_ADDR=<your_redis_address>
REDIS_PASSWORD=<your_redis_password>
REDIS_DB=<your_redis_db_number>
```

#### Steps to Set Up Environment Variables

#### Step 1:
Create a .env file: In the project root, create a new file named .env.

#### Step 2:
Add the variables: Replace <your_redis_address>, <your_redis_password>, and <your_redis_db_number> with your actual Redis configuration values.

#### Step 3:
Save the file: Ensure the file is saved and accessible from the project root.

Alternatively, you can set these environment variables directly on your server or deployment platform instead of creating a .env file, especially for production deployments.

## Tech Stack Used
- ReactJs
- Redux
- GoLang
- Redis
