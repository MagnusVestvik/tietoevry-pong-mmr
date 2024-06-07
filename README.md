# Ping Pong MMR Website

![](https://media1.giphy.com/media/v1.Y2lkPTc5MGI3NjExYnVkeWw4OG51cjlqZzdkY3hrZmI3cHRhNmY2OTFxOG90NG03eGU0MSZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/pWncxUrrNHdny/giphy.webp)
This project is a web platform for tracking and ranking players' Matchmaking Rating (MMR) in ping pong using the chess MMR equation.

## Features
- Coming Soon

## Technologies Used
- **Frontend:** HTML, CSS, JavaScript, Svelte<3, tailwind
- **Backend:** Go

## Getting Started
1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourusername/ping-pong-mmr.git
   cd ping-pong-mmr
   ```
2. **Create database and set secret**
   ```bash
   docker run --name mmr-postgres -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres
   export DATABASE_SECRET="host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Oslo"
   ```
3. **Start backend**
   ```bash
   cd api && go run main.go
   ```
