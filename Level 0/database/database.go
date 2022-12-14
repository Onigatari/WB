package database

import (
	"Level0/parser"
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"github.com/qustavo/dotsql"
	"log"
	"os"
)

const (
	host     = "localhost"
	port     = "5400"
	user     = "admin"
	password = "admin"
	database = "postgres"
)

func CreateDbTablesIfNotExist() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database)

	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, database)
	conn, err := pgx.Connect(context.Background(), databaseURL)

	if err != nil {
		log.Println(fmt.Sprintf("unable to connect to database: %v\n", err))
		os.Exit(1)
	}

	log.Println("Database is connect")
	defer conn.Close(context.Background())

	db, err := sql.Open("postgres", psqlInfo)
	dot, err := dotsql.LoadFromFile("database/create.sql")

	_, dbExistsCheck := conn.Query(context.Background(), "SELECT * FROM orders;")

	if dbExistsCheck != nil {
		_, err = dot.Exec(db, "create-table-deliveries")
		_, err = dot.Exec(db, "create-table-payments")
		_, err = dot.Exec(db, "create-table-orders")
		_, err = dot.Exec(db, "create-table-items")
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
}

func GetPostgresConnection() *pgxpool.Pool {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, database)
	conn, err := pgxpool.Connect(context.Background(), dsn)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

func GetSqlDB() (*sql.DB, *dotsql.DotSql) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, database)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	dot, err := dotsql.LoadFromFile("database/create.sql")

	conn := GetPostgresConnection()

	_, dbExistsCheck := conn.Query(context.Background(), "SELECT * FROM orders;")
	if dbExistsCheck != nil {
		_, err = dot.Exec(db, "create-table-deliveries")
		_, err = dot.Exec(db, "create-table-payments")
		_, err = dot.Exec(db, "create-table-orders")
		_, err = dot.Exec(db, "create-table-items")
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close()
	return db, dot
}

func InsertOrderToDB(order parser.Order) {
	conn := GetPostgresConnection()

	orderFromDB, _ := GetOrderFromDb(order.OrderUID)
	if orderFromDB.OrderUID == order.OrderUID {
		log.Println("such a record is already in DB")
		return
	}

	delivery := order.Delivery
	payment := order.Payment
	item := order.Items

	command, err := InsertDelivery(conn, &delivery)
	log.Println(command.String())
	if err != nil {
		log.Fatal(err)
	}

	command, err = InsertPayment(conn, &payment)
	log.Println(command.String())
	if err != nil {
		log.Fatal(err)
	}

	command, err = InsertOrders(conn, &order)
	log.Println(command.String())
	if err != nil {
		log.Fatal(err)
	}

	for _, elem := range item {
		command, err = InsertItem(conn, &elem, &order)
		log.Println(command.String())
		if err != nil {
			log.Fatal(err)
		}
	}

	defer conn.Close()
}

func InsertDelivery(conn *pgxpool.Pool, delivery *parser.Delivery) (pgconn.CommandTag, error) {
	command, err := conn.Exec(context.Background(),
		`INSERT INTO delivery
				(delivery_name, phone, zip, city, address, region, email) 
			VALUES($1, $2, $3, $4, $5, $6, $7)`,
		delivery.Name,
		delivery.Phone,
		delivery.Zip,
		delivery.City,
		delivery.Address,
		delivery.Region,
		delivery.Email)

	return command, err
}

func InsertOrders(conn *pgxpool.Pool, order *parser.Order) (pgconn.CommandTag, error) {
	command, err := conn.Exec(context.Background(),
		`INSERT INTO orders
				(order_uid, track_number, entry, delivery, payment, locale, 
				internal_signature, customer_id, delivery_service, shardkey, 
				sm_id, date_created, oof_shard)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`,
		order.OrderUID,
		order.TrackNumber,
		order.Entry,
		order.Delivery.Name,
		order.Payment.Transaction,
		order.Locale,
		order.InternalSignature,
		order.CustomerID,
		order.DeliveryService,
		order.Shardkey,
		order.SmID,
		order.DateCreated,
		order.OofShard)

	return command, err
}

func InsertPayment(conn *pgxpool.Pool, payment *parser.Payment) (pgconn.CommandTag, error) {
	command, err := conn.Exec(context.Background(),
		`INSERT INTO payment
				(payment_transaction, request_id, currency, payment_provider, 
				amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		payment.Transaction,
		payment.RequestID,
		payment.Currency,
		payment.Provider,
		payment.Amount,
		payment.PaymentDt,
		payment.Bank,
		payment.DeliveryCost,
		payment.GoodsTotal,
		payment.CustomFee)

	return command, err
}

func InsertItem(conn *pgxpool.Pool, item *parser.Item, order *parser.Order) (pgconn.CommandTag, error) {
	command, err := conn.Exec(context.Background(),
		`INSERT INTO item
				(chrt_id, order_uid, track_number, price, rid, item_name, 
				sale, size, total_price, nm_id, brand, status)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
		item.ChrtID,
		order.OrderUID,
		item.TrackNumber,
		item.Price,
		item.Rid,
		item.Name,
		item.Sale,
		item.Size,
		item.TotalPrice,
		item.NmID,
		item.Brand,
		item.Status)

	return command, err
}

func GetOrderFromDb(orderUid string) (parser.Order, error) {
	conn := GetPostgresConnection()

	var order parser.Order
	rows := conn.QueryRow(context.Background(),
		`SELECT * FROM orders WHERE order_uid = $1`, orderUid)
	err := rows.Scan(
		&order.OrderUID,
		&order.TrackNumber,
		&order.Entry,
		&order.Delivery.Name,
		&order.Payment.Transaction,
		&order.Locale,
		&order.InternalSignature,
		&order.CustomerID,
		&order.DeliveryService,
		&order.Shardkey,
		&order.SmID,
		&order.DateCreated,
		&order.OofShard)

	if err != nil {
		return order, err
	}

	rows = conn.QueryRow(context.Background(),
		`SELECT * FROM delivery WHERE delivery_name = $1`, order.Delivery.Name)
	err = rows.Scan(
		&order.Delivery.Name,
		&order.Delivery.Phone,
		&order.Delivery.Zip,
		&order.Delivery.City,
		&order.Delivery.Address,
		&order.Delivery.Region,
		&order.Delivery.Email)

	if err != nil {
		return order, err
	}

	rows = conn.QueryRow(context.Background(),
		`SELECT * FROM payment WHERE payment_transaction = $1`, order.Payment.Transaction)
	err = rows.Scan(
		&order.Payment.Transaction,
		&order.Payment.RequestID,
		&order.Payment.Currency,
		&order.Payment.Provider,
		&order.Payment.Amount,
		&order.Payment.PaymentDt,
		&order.Payment.Bank,
		&order.Payment.DeliveryCost,
		&order.Payment.GoodsTotal,
		&order.Payment.CustomFee)

	if err != nil {
		return order, err
	}

	itemsRow, err := conn.Query(context.Background(),
		`SELECT * FROM item WHERE order_uid = $1`, orderUid)
	if err != nil {
		return order, err
	}

	var item = parser.Item{}
	for itemsRow.Next() {
		_ = itemsRow.Scan(
			&item.ChrtID,
			&order.OrderUID,
			&item.TrackNumber,
			&item.Price,
			&item.Rid,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.TotalPrice,
			&item.NmID,
			&item.Brand,
			&item.Status)

		order.Items = append(order.Items, item)
	}

	defer conn.Close()
	return order, err
}

func GetAllOrdersFromDB() []parser.Order {
	conn := GetPostgresConnection()
	var orders []parser.Order
	var orderUid string

	rows, err := conn.Query(context.Background(),
		`SELECT order_uid FROM orders;`)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&orderUid)
		if err != nil {
			return nil
		}

		order, _ := GetOrderFromDb(orderUid)
		orders = append(orders, order)
	}
	defer conn.Close()
	return orders
}
