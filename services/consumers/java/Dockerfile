FROM openjdk:17-jdk-slim
WORKDIR /app
COPY . .
RUN javac -cp "lib/*" src/main/java/com/example/consumer/Consumer.java -d out
CMD ["java", "-cp", "out:lib/*", "com.example.consumer.Consumer"]