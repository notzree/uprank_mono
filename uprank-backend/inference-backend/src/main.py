from server import Server
from services import InferenceService
import os
from pinecone import Pinecone
import sys
import signal
from langchain_openai import OpenAIEmbeddings
import logging

logging.basicConfig(level=logging.INFO)
if __name__ == '__main__':
    server_logger = logging.getLogger("server")
    service_logger = logging.getLogger("services") 
    pinecone_api_key = os.environ.get("PINECONE_API_KEY")
    openai_api_key = os.environ.get("OPENAI_API_KEY")
    server_port = os.environ.get("SERVER_PORT")
    pc = Pinecone(api_key=pinecone_api_key)
    model_name = "text-embedding-ada-002"
    embeddings = OpenAIEmbeddings(
        model=model_name,
    )
    infer = InferenceService(embed=embeddings, pc=pc, index_name="uprank-dev", logger=service_logger)
    s = Server(infer=infer, max_workers=10, port=server_port, logger=server_logger)

    def handle_signal(sig, frame):
        sys.exit(0)

    signal.signal(signal.SIGINT, handle_signal)
    signal.signal(signal.SIGTERM, handle_signal)
    s.serve()
