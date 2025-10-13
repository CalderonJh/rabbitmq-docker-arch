import pika, json, os

RABBIT_HOST = os.getenv("RABBIT_HOST", "rabbitmq")

connection = pika.BlockingConnection(pika.ConnectionParameters(RABBIT_HOST))
channel = connection.channel()
channel.queue_declare(queue="events")

def callback(ch, method, properties, body):
    data = json.loads(body)
    print(f"Received from {data['source']}: {data['message']}")
    ch.basic_ack(delivery_tag=method.delivery_tag)

channel.basic_consume(queue="events", on_message_callback=callback)
print(" [*] Waiting for messages. To exit press CTRL+C")
channel.start_consuming()