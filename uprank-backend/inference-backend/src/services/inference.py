from generated import inference_pb2_grpc, inference_pb2
from langchain_openai import OpenAIEmbeddings
from pinecone.grpc import PineconeGRPC
from pinecone import ServerlessSpec
from pinecone.core.grpc.protos.vector_service_pb2 import Vector as Vector
from typing import List
import grpc
import pprint
from logging import Logger


class InferenceService(inference_pb2_grpc.InferenceServicer):
    def __init__(self, embed: OpenAIEmbeddings, pc: PineconeGRPC, index_name: str, logger: Logger):
        self.embed = embed
        self.pc = pc
        self.logger = logger
        if index_name not in pc.list_indexes().names():
            # hardcode the index details for now
            index = pc.create_index(
                name=index_name,
                dimension=1536,
                metric="cosine",
                spec=ServerlessSpec(
                    cloud="aws",
                    region="us-east-1",
                ),
            )
        else:
            index = pc.Index(index_name)
        self.index = index

    def EmbedText(self, request, context):
        try:
            embeddings = self.embed.embed_query(request.text)
            vector = inference_pb2.Vector(
                id=request.id,
                vector=embeddings,
                metadata=request.metadata
            )
            return inference_pb2.EmbedTextResponse(
                vector=vector
            )
        except Exception as e:
            context.set_code(grpc.StatusCode.UNKNOWN)
            return inference_pb2.EmbedTextResponse(
                vector=inference_pb2.Vector(
                    id="",
                    vector=[],
                    metadata={}
                ),
                error=e
            )

    def UpsertVector(self, request, context):
        transformed_vectors = []
        for v in request.vectors:
            transformed_vectors.append((v.id, v.vector, dict(v.metadata)))
        try:
            upsert_response = self.index.upsert(
                namespace=request.namespace,
                vectors=transformed_vectors
            )
        except Exception as e:
            context.set_code(grpc.StatusCode.UNKNOWN)
            return inference_pb2.UpsertVectorResponse(
                error=e
            )
        context.set_code(grpc.StatusCode.OK)
        return inference_pb2.UpsertVectorResponse(
            error=""
        )

    def QueryVector(self, request, context):
        matches: List[inference_pb2.Match] = []
        demo ={
            'job_id':'1788207506953621504',
            'type': 'work_history_description',
        }
        print(type(demo), type(dict(request.filter)))
        print(demo == dict(request.filter))
        print(dict(request.filter))
        try:
            query_response = self.index.query(
                namespace=request.namespace,
                vector=list(request.vector),
                top_k=request.top_k,
                include_metadata=True,
                filter=dict(request.filter)
            )
        except Exception as e:
            context.set_code(grpc.StatusCode.UNKNOWN)
            context.set_details(str(e))
            return inference_pb2.QueryVectorResponse(
                matches=[],
                namespace=request.namespace,
                usage=inference_pb2.Usage(
                    read_units=0
                )
            )
        for result in query_response['matches']:
            matches.append(
                inference_pb2.Match(
                    id=result['id'],
                    score=result['score'],
                    metadata=result['metadata']
                )
            )
        with open('query_response.txt', 'w') as f:
            pprint.pprint(query_response, f)
        context.set_code(grpc.StatusCode.OK)
        return inference_pb2.QueryVectorResponse(
            matches=matches,
            namespace=request.namespace,
            usage=inference_pb2.Usage(
                read_units= query_response['usage']['read_units']
            )
        )

    def DeleteVector(self, request, context):
        try:
            delete_response = self.index.delete(
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
