package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	_ "embed"

	"github.com/go-playground/validator"
	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

const (
	PubSubTopic = "broadcast"
	RedisKey    = "chat_messages"
	TimeFormat  = "15:04:05"
)

var (
	users       Users
	rdb         *redis.Client
	broadcaster *redis.PubSub
)

var (
	clients  = make(map[*websocket.Conn]bool)
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// ensure connection close when function returns
	defer ws.Close()
	clients[ws] = true

	// if it's zero, no messages were ever sent/saved
	if rdb.Exists(RedisKey).Val() != 0 {
		sendPreviousMessages(ws)
	}

	for {
		var msg ChatMessage
		// Read in a new message
		err := ws.ReadJSON(&msg)
		if err != nil {
			delete(clients, ws)
			break
		}

		if err := msg.Validate(); err != nil {
			log.Printf("error: %v", err)
			continue
		}

		msg.Timestamp = time.Now()

		// send new message to redis PUB/SUB channel
		if err := rdb.Publish(PubSubTopic, msg).Err(); err != nil {
			panic(err)
		}

		// store message in redis list
		storeInRedis(msg)
	}
}

func sendPreviousMessages(ws *websocket.Conn) {
	chatMessages, err := rdb.LRange(RedisKey, 0, -1).Result()
	if err != nil {
		panic(err)
	}

	// send previous messages
	for _, chatMessage := range chatMessages {
		var msg ChatMessage
		if err := msg.FromJson(chatMessage); err != nil {
			panic(err)
		}
		messageClient(ws, msg)
	}
}

// If a message is sent while a client is closing, ignore the error
func unsafeError(err error) bool {
	return !websocket.IsCloseError(err, websocket.CloseGoingAway) && err != io.EOF
}

func handleMessages() {
	for rmsg := range broadcaster.Channel() {
		// grab any next message from channel
		var msg ChatMessage
		if err := msg.FromJson(rmsg.Payload); err != nil {
			panic(err)
		}

		messageClients(msg)
	}
}

func storeInRedis(msg ChatMessage) {
	if err := rdb.RPush(RedisKey, msg).Err(); err != nil {
		panic(err)
	}
}

func messageClients(msg ChatMessage) {
	// send to every client currently connected
	for client := range clients {
		messageClient(client, msg)
	}
}

func messageClient(client *websocket.Conn, msg ChatMessage) {
	msg.Time = msg.Timestamp.Format(TimeFormat)

	err := client.WriteJSON(msg)
	if err != nil && unsafeError(err) {
		log.Printf("error: %v", err)
		client.Close()
		delete(clients, client)
	}
}

//go:embed public/index.html
var index string

//go:embed public/app.js
var appjs string

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	port := os.Getenv("PORT")

	redisURL := os.Getenv("REDIS_URL")
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		panic(err)
	}
	rdb = redis.NewClient(opt)

	b, err := os.ReadFile("secrets/users.json")
	if err != nil {
		panic(err)
	}
	if err := users.FromJson(string(b)); err != nil {
		panic(err)
	}

	broadcaster = rdb.Subscribe(PubSubTopic)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(index))
	})
	http.HandleFunc("/app.js", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(appjs))
	})
	http.HandleFunc("/websocket", handleConnections)
	go handleMessages()

	log.Print("Server starting at localhost:" + port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

type ChatMessage struct {
	Username  string    `json:"username" validate:"required,min=3,max=20"`
	Text      string    `json:"text" validate:"required"`
	Time      string    `json:"time"`
	Timestamp time.Time `json:"timestamp"`
}

func (c *ChatMessage) FromJson(in string) error {
	return json.Unmarshal([]byte(in), c)
}

func (c *ChatMessage) Validate() error {
	validate := validator.New()

	if !users.ValidUser(c.Username) {
		return errors.New("invalid username")
	}

	return validate.Struct(c)
}

func (c ChatMessage) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

type Users []string

func (u *Users) FromJson(in string) error {
	return json.Unmarshal([]byte(in), u)
}

func (u Users) ValidUser(in string) bool {
	for _, usr := range u {
		if usr == in {
			return true
		}
	}
	return false
}
