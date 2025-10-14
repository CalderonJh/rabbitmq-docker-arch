import pika, json, time, os

RABBIT_HOST = os.getenv("RABBIT_HOST", "rabbitmq")
credentials = pika.PlainCredentials("admin", "admin")

connection = pika.BlockingConnection(pika.ConnectionParameters(RABBIT_HOST, credentials=credentials))
channel = connection.channel()

# Exchange fanout
channel.exchange_declare(exchange="events-exchange", exchange_type="fanout", durable=True)

for i in range(5):
    msg = {"source": "producer-python", "message": f"event {i}"}
    channel.basic_publish(
        exchange="events-exchange",
        routing_key="",
        body=json.dumps(msg),
        properties=pika.BasicProperties(delivery_mode=2)  # persistent
    )
    print("Sent:", msg)
    time.sleep(1)

connection.close()
