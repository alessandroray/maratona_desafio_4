package main_test

import (
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	err := godotenv.Load()
	
	if err != nil {
		t.Errorf("Falha Carregar .env")
	}
}
func TestConnect (t *testing.T) { 
	dsn := "amqp://" + os.Getenv("RABBITMQ_DEFAULT_USER") + ":" + os.Getenv("RABBITMQ_DEFAULT_PASS") + "@" + os.Getenv("RABBITMQ_DEFAULT_HOST") + ":" + os.Getenv("RABBITMQ_DEFAULT_PORT") + os.Getenv("RABBITMQ_DEFAULT_VHOST")
	_, err := amqp.Dial(dsn)

	if err  != nil {
		t.Error("Falha Conection")
		return
	}
}
