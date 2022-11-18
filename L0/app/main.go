package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
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

func startHttpServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})
}

func connectDB() *sql.DB {
	connStr := "postgres://lumina:2408mM305@localhost/wb?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return db
}

func insertDB(db *sql.DB, order OrderModel) {
	// Upload to postgres
}

func fetchDB(db *sql.DB, orderList *[]OrderModel) {
	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rows.Close()

	for rows.Next() {
		var order OrderModel
		// err := rows.Scan(&order)
		// if err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }

		*orderList = append(*orderList, order)
	}
}

func main() {

	// Plan:
	// 1. Connect to postgres
	// 2. Restore cache from postgres
	// 3. Host http webserver or web interface
	// 4. Tell webserver to serve id requests
	// 5. Connect to NATS
	// 6. Subscribe to NATS
	// 7. Handle NATS messages along with http requests

	var ordersCache []OrderModel

	db := connectDB()
	defer db.Close()

	fetchDB(db, &ordersCache)

	fmt.Println(ordersCache)

	//startHttpServer()

	/////////////////////////////////////////////////
	// NATS connection example //
	/////////////////////////////////////////////////

	// Nats streaming connection
	// nc, err := nats.Connect(nats.DefaultURL)
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// defer nc.Close()

	// Subscribe to subject
	// _, err = nc.Subscribe("channel", func(msg *nats.Msg) {
	// 	fmt.Printf("Received a message: %s", string(msg.Data))
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// Publish message
	// err = nc.Publish("channel", []byte("Hello World"))
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// nc.Flush()

	/////////////////////////////////////////////////
	// JSON parsing example //
	/////////////////////////////////////////////////

	// filename := "./model.json"
	// plan, error := os.ReadFile(filename)
	// if error != nil {
	// 	fmt.Println(error)
	// }

	// var data OrderModel
	// err = json.Unmarshal(plan, &data)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(data)

	/////////////////////////////////////////////////
	// HTTP server example //
	/////////////////////////////////////////////////

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello World")
	// })

	// http.ListenAndServe(":8080", nil)

}
