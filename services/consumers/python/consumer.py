import pika, json, os

RABBIT_HOST = os.getenv("RABBIT_HOST", "rabbitmq")

credentials = pika.PlainCredentials("admin", "admin")
connection = pika.BlockingConnection(pika.ConnectionParameters(RABBIT_HOST, credentials=credentials))
channel = connection.channel()

channel.queue_declare(queue="events", durable=True)

def callback(ch, method, properties, body):
    try:
        if not body:
            print("Received empty message")
            ch.basic_ack(delivery_tag=method.delivery_tag)
            return

        data = json.loads(body.decode('utf-8'))
        print(f"Received from {data.get('source', 'unknown')}: {data.get('message', body)}")
    except json.JSONDecodeError:
        print(f"Non-JSON message received: {body}")
    finally:
        ch.basic_ack(delivery_tag=method.delivery_tag)


channel.basic_consume(queue="events", on_message_callback=callback)
print(" [*] Waiting for messages. To exit press CTRL+C")
channel.start_consuming()