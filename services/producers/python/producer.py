import pika, json, time, os

RABBIT_HOST = os.getenv("RABBIT_HOST", "rabbitmq")
credentials = pika.PlainCredentials("admin", "admin")

while True:
    try:
        connection = pika.BlockingConnection(pika.ConnectionParameters(RABBIT_HOST, credentials=credentials))
        channel = connection.channel()
        channel.exchange_declare(exchange="events-topic-exchange", exchange_type="topic", durable=True)

        i = 0
        while True:
            msg = {"source": "producer-python", "message": f"event {i}"}
            channel.basic_publish(
                exchange="events-topic-exchange",
                routing_key="events.python",
                body=json.dumps(msg),
                properties=pika.BasicProperties(delivery_mode=2),
            )
            print("Sent:", msg)
            i += 1
            time.sleep(5)
    except Exception as e:
        print("Connection failed, retrying...", e)
        time.sleep(3)