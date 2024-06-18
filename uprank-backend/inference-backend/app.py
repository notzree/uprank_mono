
import os
import boto3

from queues import SQSQueue
from server import Server
from service import UpworkService

# Initialize SQS client


if __name__ == '__main__':
    sqs_client = boto3.client('sqs', region_name='us-east-2')
    ranking_queue_url = os.environ.get("RANKING_QUEUE_URL")
    q = SQSQueue(sqs_client, ranking_queue_url)
    svc = UpworkService("http://localhost:8080")
    server = Server(q, svc, 4)
    try:
        print("starting listener")
        server.start()
    except KeyboardInterrupt:
        print("Exiting")
        exit(0)
