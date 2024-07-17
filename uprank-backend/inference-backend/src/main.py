from server import Server
from services import InferenceService
import os
import json
from pinecone import Pinecone
import sys
import signal
from langchain_openai import OpenAIEmbeddings
import logging

logging.basicConfig(level=logging.INFO)
if __name__ == '__main__':
    env = os.getenv("ENV")
    server_logger = logging.getLogger("server")
    service_logger = logging.getLogger("services") 
    pinecone_api_key = ""
    openai_api_key = ""
    server_port = 0
    if env == "local" or not env:
        pinecone_api_key = os.getenv("PINECONE_API_KEY")
        openai_api_key = os.getenv("OPENAI_API_KEY")
        server_port = os.getenv("EXPOSED_INFERENCE_SERVER_PORT")
    elif env == "dev":
        logging.info("Running in dev environment")
        all_secrets = os.getenv("MAIN_BACKEND_SECRETS")
        if not all_secrets:
            logging.error("failed to load secrets")
            sys.exit(1)
        secret_map = json.loads(all_secrets)
        pinecone_api_key = secret_map["PINECONE_API_KEY"]
        openai_api_key = secret_map["OPENAI_API_KEY"]
        server_port = int(secret_map["EXPOSED_INFERENCE_SERVER_PORT"])
        
    pc = Pinecone(api_key=pinecone_api_key)
    model_name = "text-embedding-ada-002"
    embeddings = OpenAIEmbeddings(
        model=model_name,
        openai_api_key=openai_api_key,
    )
    infer = InferenceService(embed=embeddings, pc=pc, index_name="uprank-dev", logger=service_logger)
    s = Server(infer=infer, max_workers=10, port=server_port, logger=server_logger)

    def handle_signal(sig, frame):
        sys.exit(0)

    signal.signal(signal.SIGINT, handle_signal)
    signal.signal(signal.SIGTERM, handle_signal)
    s.serve()
