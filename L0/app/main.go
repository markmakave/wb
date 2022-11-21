package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/lib/pq"
	"github.com/nats-io/nats.go"
)

type DeliveryModel struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type PaymentModel struct {
	Transaction  string `json:"transaction"`
	RequestId    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type ItemModel struct {
	ChrtId      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmId        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

type OrderModel struct {
	OrderUid           string        `json:"order_uid"`
	TrackNumber        string        `json:"track_number"`
	Entry              string        `json:"entry"`
	Delivery           DeliveryModel `json:"delivery"`
	Payment            PaymentModel  `json:"payment"`
	Items              []ItemModel   `json:"items"`
	Locale             string        `json:"locale"`
	Internal_signature string        `json:"internal_signature"`
	CustomerId         string        `json:"customer_id"`
	DeliveryService    string        `json:"delivery_service"`
	Shardkey           string        `json:"shardkey"`
	SmId               int           `json:"sm_id"`
	DateCreated        string        `json:"date_created"`
	OofShard           string        `json:"oof_shard"`
}

// ///////////////////////////////////////////////

func startHttpServer(ordersCache *map[string]OrderModel) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./html/index.html")
		fmt.Println("[ HTTP ] Served index.html")
	})

	// handle get request in format /orders/{id}
	http.HandleFunc("/orders/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/orders/"):]
		order, ok := (*ordersCache)[id]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Println("[ HTTP ] Order not found with id", id)
			return
		}

		json.NewEncoder(w).Encode(order)
		fmt.Println("[ HTTP ] Served order with id", id)
	})

	// starts http server main loop
	http.ListenAndServe(":8080", nil)
}

func connectDB() *sql.DB {
	connStr := "postgres://lumina:password@localhost/wb?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		// error connecting to database
		// usually means that the database is not running of credentials are wrong
		fmt.Println("[ POSTGRES ] Error occurred:", err)
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		// error pinging database
		// usually means that the database is not responding
		fmt.Println("[ POSTGRES ] Error occurred:", err)
		os.Exit(1)
	}

	return db
}

func insertDB(db *sql.DB, id string, rawJson json.RawMessage) {
	_, err := db.Exec("INSERT INTO orders (id, data) VALUES ($1, $2)", id, rawJson)
	if err != nil {
		// error inserting to database
		// usually is caused by dudplicate id
		if pgerr, ok := err.(*pq.Error); ok {
			// duplicate id handling
			if pgerr.Code == "23505" {
				// handle duplicate insert
				// I update the duplicate row but duplicates can also be ignored or handled in some other way
				_, err := db.Exec("UPDATE orders SET data = $1 WHERE id = $2", rawJson, id)
				if err != nil {
					// error while updating duplicate row
					// is database dead?
					fmt.Println(err)
					os.Exit(1)
				}
				fmt.Println("[ POSTGRES ] Updated at index", id)
				return
			}
		}

		// other error
		// no way to handle it correctly from application
		fmt.Println("[ POSTGRES ] Error occurred:", err)
		os.Exit(1)
	}

	fmt.Println("[ POSTGRES ] Inserted at index", id)
}

func fetchDB(db *sql.DB, orderList *map[string]OrderModel) {
	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		// database error
		// no way to handle it correctly from application
		fmt.Println("[ POSTGRES ] Error occurred:", err)
		os.Exit(1)
	}
	defer rows.Close()

	for rows.Next() {
		var order OrderModel

		var id string
		var rawJson json.RawMessage
		err := rows.Scan(&id, &rawJson)
		if err != nil {
			fmt.Println("[ POSTGRES ] Error occurred:", err)
			os.Exit(1)
		}

		// Parse json to struct
		err = json.Unmarshal(rawJson, &order)
		if err != nil {
			// error parsing json from database object
			// literaly impossible to happen
			// but I handle it anyway
			fmt.Println("[ POSTGRES ] Invalid data from database:", err)
			os.Exit(1)
		}

		(*orderList)[id] = order
	}

	fmt.Println("[ POSTGRES ] Fetched from DB")
}

func main() {

	// Plan:
	// 1. Connect to postgres
	// 2. Restore cache from postgres
	// 3. Connect to NATS
	// 4. Subscribe to NATS channel
	// 5. Handle NATS messages
	// 6. Host http webserver for web interface
	// 7. Tell webserver to serve id requests
	// 8. Sit in http server main loop until cthulhu engulfs the world or Ctrl+C is pressed

	// Order model cache array
	var ordersCache = make(map[string]OrderModel)

	// Connect to db
	db := connectDB()
	defer db.Close()

	// Restore cache from db
	fetchDB(db, &ordersCache)

	// Establish NATS connection
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		// NATS connection error
		// usually forgot to start NATS server
		fmt.Println("[ NATS ] Error occurred:", err)
		os.Exit(1)
	}
	defer nc.Close()

	// Subscribe to NATS channel
	channelName := "orders"
	nc.Subscribe(channelName, func(m *nats.Msg) {
		var order OrderModel

		// Parse json to struct
		err = json.Unmarshal(m.Data, &order)
		if err != nil {
			// error parsing json from NATS message
			// usually means that NATS message is not in correct format
			// or that the message is not json at all

			// I ignore the message and continue
			fmt.Println("[ NATS ] Error parsing json (garbage data in channel?):", err)
			return
		}
		fmt.Println("[ NATS ] Received order with id", order.OrderUid)

		// Insert into or update cache
		ordersCache[order.OrderUid] = order

		// Insert into or update on db
		insertDB(db, order.OrderUid, m.Data)
	})

	// Start http server
	// main loop is provided by http server
	startHttpServer(&ordersCache)

}
