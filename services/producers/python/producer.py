import pika, json, time, os

RABBIT_HOST = os.getenv("RABBIT_HOST", "rabbitmq")

credentials = pika.PlainCredentials("admin", "admin")
connection = pika.BlockingConnection(pika.ConnectionParameters(RABBIT_HOST, credentials=credentials))
channel = connection.channel()

# Ensure same durability everywhere
channel.queue_declare(queue="events", durable=True)

for i in range(5):
    msg = {"source": "producer-python", "message": f"event {i}"}
    channel.basic_publish(
        exchange="",
        routing_key="events",
        body=json.dumps(msg),
        properties=pika.BasicProperties(delivery_mode=2)  # make message persistent
    )
    print("Sent:", msg)
    time.sleep(1)

connection.close()