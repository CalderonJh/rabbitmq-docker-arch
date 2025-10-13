FROM openjdk:17-jdk-slim
WORKDIR /app
RUN apt-get update && apt-get install -y netcat && rm -rf /var/lib/apt/lists/*
COPY . .
RUN javac -cp "lib/*" src/main/java/com/example/producer/Producer.java -d out
CMD ["java", "-cp", "out:lib/*", "com.example.producer.Producer"]