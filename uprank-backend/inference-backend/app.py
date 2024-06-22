import os

import boto3
from langchain_openai import OpenAIEmbeddings
from pinecone.grpc import PineconeGRPC as Pinecone

from queues import SQSQueue
from server import Server
from service import UpworkService
from vector import Vector

# Initialize SQS client


if __name__ == "__main__":
    ranking_queue_url = os.environ.get("RANKING_QUEUE_URL")
    main_backend_url = os.environ.get("MAIN_BACKEND_URL")
    pinecone_api_key = os.environ.get("PINECONE_API_KEY")
    inference_server_url = os.environ.get("INFERENCE_SERVER_URL")
    openai_api_key = os.environ.get("OPENAI_API_KEY")
    sqs_client = boto3.client("sqs", region_name="us-east-2")

    pc = Pinecone(api_key="YOUR_API_KEY")
    model_name = "text-embedding-ada-002"
    embeddings = OpenAIEmbeddings(
        model=model_name,
        openai_api_key=openai_api_key,
    )
    embeddings.
    pc_vector = Vector(pc=pc, index_name="uprank-dev", embd=embeddings)


    q = SQSQueue(sqs_client, ranking_queue_url)
    svc = UpworkService(main_backend_url=main_backend_url, inference_server_url=inference_server_url, vector=pc_vector)
    server = Server(q, svc, 4)
    try:
        print("starting listener")
        server.start()
    except KeyboardInterrupt:
        print("Exiting")
        exit(0)
