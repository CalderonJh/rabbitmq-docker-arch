FROM python:3.11-slim
WORKDIR /app
RUN apt-get update && apt-get install -y --no-install-recommends \
    netcat-openbsd \
    ca-certificates \
    apt-transport-https \
 && rm -rf /var/lib/apt/lists/*
COPY consumer.py .
RUN pip install pika
CMD ["python", "consumer.py"]
