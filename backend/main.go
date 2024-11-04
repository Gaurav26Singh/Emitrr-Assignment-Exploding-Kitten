package main

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
)

var rdb *redis.Client

func init() {
	// Attempt to load the .env file, but continue without it if it's missing
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found; using system environment variables")
	}

	// Initialize Redis client using environment variables
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0, // Adjust DB number if necessary
	})

	// Test the Redis connection
	if _, err := rdb.Ping(rdb.Context()).Result(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	} else {
		log.Println("Connected to Redis successfully.")
	}
}

type UserData struct {
	Username string `json:"username"`
	Points   int    `json:"points"`
}

func setUser(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	name := data["name"]

	exists, err := rdb.Exists(r.Context(), name).Result()
	if err != nil {
		log.Println("Error checking if user exists:", err)
		http.Error(w, "Failed to check if user exists", http.StatusInternalServerError)
		return
	}

	if exists == 0 {
		err := rdb.Set(r.Context(), name, 0, 0).Err()
		if err != nil {
			log.Println("Error setting points for new user:", err)
			http.Error(w, "Failed to set points for new user", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User points initialized successfully"))
}

func getUserPoints(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	pointsStr, err := rdb.Get(r.Context(), name).Result()
	if err == redis.Nil {

		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Println("Error getting user points:", err)
		http.Error(w, "Failed to get user points", http.StatusInternalServerError)
		return
	}

	points, err := strconv.Atoi(pointsStr)
	if err != nil {
		log.Println("Error converting points to integer:", err)
		http.Error(w, "Failed to convert points to integer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(points)
}

func updateUserPoints(w http.ResponseWriter, r *http.Request) {

	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	name := data["name"]

	currentPoints, err := rdb.Get(r.Context(), name).Int()
	if err != nil && err != redis.Nil {
		log.Println("Error getting user points:", err)
		http.Error(w, "Failed to get user points", http.StatusInternalServerError)
		return
	}

	newPoints := currentPoints + 1
	err = rdb.Set(r.Context(), name, newPoints, 0).Err()
	if err != nil {
		log.Println("Error updating user points:", err)
		http.Error(w, "Failed to update user points", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message": "User points updated successfully",
		"points":  newPoints,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getAllUserPoints(w http.ResponseWriter, r *http.Request) {

	keys := rdb.Keys(r.Context(), "*").Val()
	userPointsMap := make(map[string]int)
	for _, key := range keys {
		pointsStr, err := rdb.Get(r.Context(), key).Result()
		if err != nil && err != redis.Nil {
			log.Println("Error getting user points for key", key, ":", err)
			continue
		}

		points, err := strconv.Atoi(pointsStr)
		if err != nil {
			log.Println("Error parsing user points for key", key, ":", err)
			continue
		}

		userPointsMap[key] = points
	}

	responseJSON, err := json.Marshal(userPointsMap)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(responseJSON); err != nil {
		log.Println("Error writing response:", err)
	}
}

func getLeaderboard(w http.ResponseWriter, r *http.Request) {

	keys := rdb.Keys(r.Context(), "*").Val()
	leaderboard := make([]UserData, len(keys))
	for i, key := range keys {
		pointsStr, err := rdb.Get(r.Context(), key).Result()
		if err != nil && err != redis.Nil {
			log.Println("Error getting user points for key", key, ":", err)
			continue
		}

		points, err := strconv.Atoi(pointsStr)
		if err != nil {
			log.Println("Error parsing user points for key", key, ":", err)
			continue
		}

		leaderboard[i] = UserData{
			Username: key,
			Points:   points,
		}
	}

	sort.Slice(leaderboard, func(i, j int) bool {
		return leaderboard[i].Points > leaderboard[j].Points
	})

	responseJSON, err := json.Marshal(leaderboard)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(responseJSON); err != nil {
		log.Println("Error writing response:", err)
	}
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/api/user", setUser).Methods("POST")
	r.HandleFunc("/api/user/points", getUserPoints).Methods("GET")
	r.HandleFunc("/api/user/points", updateUserPoints).Methods("PUT")
	r.HandleFunc("/api/user/points/all", getAllUserPoints).Methods("GET")
	r.HandleFunc("/api/leaderboard", getLeaderboard).Methods("GET") // Add new endpoint for leaderboard

	c := cors.AllowAll()

	handler := c.Handler(r)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
