FROM python:3.11-slim
WORKDIR /app
RUN apt-get update && apt-get install -y --no-install-recommends \
    netcat-traditional \
    ca-certificates \
    apt-transport-https \
 && rm -rf /var/lib/apt/lists/*
COPY producer.py .
RUN pip install pika
CMD ["python", "producer.py"]