package main

import (
	"net"
	"fmt"
	"time"
	"os"
	"os/signal"
	"math/rand"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"strings"
)

var port int
var trappedCount int

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("SERVICETRAP")
	
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
        panic(err.Error())
    }
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
func RandString() string {
	n := rand.Intn(60) + 10
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

func handleConnection(connection net.Conn) {
	defer connection.Close()
	trappedCount++

	fmt.Println("Currently handling", trappedCount, "trapped connections")

	for {
		_, err := connection.Write([]byte(RandString()))
		if err != nil {
			fmt.Println("Error writing: %+v, closing connection", err)
			trappedCount--
			fmt.Println("Currently handling %+v trapped connections", trappedCount)
			return
		}

		time.Sleep(10 * time.Second)
	}
}

func addDBrow(ipAddr string) {

		/* construct sql connection string */
	cfg := &mysql.Config{
		AllowNativePasswords: true,
		User:   viper.GetString("DB_USER"),
		Passwd: viper.GetString("DB_PASSWORD"),
		Net:    "tcp",
		Addr:   viper.GetString("DB_HOST") + ":" + viper.GetString("DB_PORT"),
		DBName: viper.GetString("DB_NAME"),
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	
	// if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }
	defer db.Close()
	
	reverseHostname, err := net.LookupAddr(ipAddr)
	if err != nil {
        panic(err.Error())
    }
	reverseHostnameString := strings.Join(reverseHostname,", ")
	
	hostname, err := os.Hostname()
	if err != nil {
		panic(err.Error())
	}
		

    sql := "INSERT INTO `" + viper.GetString("DB_TABLE") + "` (`id`, `ip`, `port`, `protocol`, `category`, `hostname`, `country`, `rdns`, `timestamp`, `remove`) VALUES (NULL, '" + ipAddr + "', '" + viper.GetString("LISTEN_PORT") + "', 'tcp', '" + viper.GetString("DATA_CATEGORY") + "', '" + hostname + "', '', '" + reverseHostnameString + "', current_timestamp(), '0')"
    res, err := db.Exec(sql)

    if err != nil {
        panic(err.Error())
    }

    lastId, err := res.LastInsertId()

    if err != nil {
        panic(err.Error())
    }

    fmt.Printf("Successfully registered " + ipAddr + " in DB row id %d\n", lastId)
}


func main() {
	trappedCount = 0
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt)
	go func() {
		_ = <-sigc
		
		fmt.Println("Received interrupt, closing")
		os.Exit(1)
	}()

	fmt.Println("Starting up service-trap on", viper.GetString("LISTEN_ADDRESS"), "port" ,viper.GetString("LISTEN_PROTOCOL") + "/" + viper.GetString("LISTEN_PORT"))

	sock, err := net.Listen(viper.GetString("LISTEN_PROTOCOL"), fmt.Sprintf(viper.GetString("LISTEN_ADDRESS") + ":" + viper.GetString("LISTEN_PORT")))
	if err != nil {
		panic(err.Error())
	}
	
	fmt.Println("Waiting for connections...")

	for {
		conn, err := sock.Accept()
		if err != nil {
			fmt.Println("Error on accept: %+v", err)
		}
		IPAddress := conn.RemoteAddr().(*net.TCPAddr).IP.String()
		fmt.Println("Connection accepted from ", conn.RemoteAddr())
		
		addDBrow(IPAddress)
		go handleConnection(conn)
	}	
}