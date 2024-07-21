FROM python:3.11-slim

WORKDIR /app

COPY requirements.txt requirements.txt
RUN pip install --no-cache-dir -r requirements.txt

# Download grpc-health-probe
ADD https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v0.3.6/grpc_health_probe-linux-amd64 /bin/grpc_health_probe
RUN chmod +x /bin/grpc_health_probe

COPY . .

EXPOSE 50051

CMD ["python", "src/main.py"]
