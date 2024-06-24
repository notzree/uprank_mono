from generated import inference_pb2_grpc, inference_pb2
from langchain_openai import OpenAIEmbeddings
from pinecone.grpc import PineconeGRPC
from pinecone.core.grpc.protos.vector_service_pb2 import Vector
import grpc
from logging import Logger


class InferenceService(inference_pb2_grpc.InferenceServicer):
    def __init__(self, embed: OpenAIEmbeddings, pc: PineconeGRPC, logger: Logger):
        self.embed = embed
        self.pc = pc
        self.logger = logger

    def EmbedText(self, request, context):
        vector = self.embed.embed_query(request.text)
        return inference_pb2.EmbedTextResponse(
            vector=vector
        )

    def UpsertVector(self, request, context):
        transformed_vectors = []
        for vector in request.vectors:
            transformed_vectors.append(Vector(
                id=vector.id,
                vector=vector.vector,
                metadata=vector.metadata
            ))
        try:
            upsert_response = self.pc.upsert(
                namespace=request.namespace,
                vectors=transformed_vectors
            )
            print(upsert_response)
        except Exception as e:
            context.set_code(grpc.StatusCode.UNKNOWN)
            return inference_pb2.UpsertVectorResponse(
                id="",
                error=e
            )
        context.set_code(grpc.StatusCode.OK)
        return inference_pb2.UpsertVectorResponse(
            id="",
            error=""
        )

    def QueryVector(self, request, context):
        try:
            query_response = self.pc.query(
                namespace=request.namespace,
                vector=request.vector,
                top_k=request.top_k,
                filter=request.filter
            )
        except Exception as e:
            context.set_code(grpc.StatusCode.UNKNOWN)
            return inference_pb2.QueryVectorResponse(
                matches=[],
                error=e
            )
        context.set_code(grpc.StatusCode.OK)
        return inference_pb2.QueryVectorResponse(
            matches=query_response["matches"],
            error=""
        )

    def DeleteVector(self, request, context):
        try:
            delete_response = self.pc.delete(
                namespace=request.namespace,
                ids=request.ids
            )
        except Exception as e:
            context.set_code(grpc.StatusCode.UNKNOWN)
            return inference_pb2.DeleteVectorResponse(
                error=e
            )
        context.set_code(grpc.StatusCode.OK)
        return inference_pb2.DeleteVectorResponse(
            error=""
        )
