import pika, json, time, os

RABBIT_HOST = os.getenv("RABBIT_HOST", "rabbitmq")

connection = pika.BlockingConnection(pika.ConnectionParameters(RABBIT_HOST))
channel = connection.channel()
channel.queue_declare(queue="events")

for i in range(5):
    msg = {"source": "producer-python", "message": f"event {i}"}
    channel.basic_publish(exchange="", routing_key="events", body=json.dumps(msg))
    print("Sent:", msg)
    time.sleep(1)

connection.close()
