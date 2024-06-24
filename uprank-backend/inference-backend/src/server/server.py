# src/server.py
import grpc
from concurrent import futures
from generated import inference_pb2_grpc
from services import InferenceService
from logging import Logger


class Server:
    def __init__(self,  infer: InferenceService, logger: Logger, max_workers=10, port=50051):
        self.port = port
        self.logger = logger
        self.infer = infer
        self.server = grpc.server(futures.ThreadPoolExecutor(max_workers=max_workers))
        self.server.add_insecure_port(f'[::]:{port}')

    def serve(self):
        inference_pb2_grpc.add_InferenceServicer_to_server(self.infer, self.server)
        self.server.start()
        self.logger.info(f"Server started on port ${self.port}")
        self.server.wait_for_termination()

