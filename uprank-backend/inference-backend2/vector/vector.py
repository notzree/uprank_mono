from __future__ import annotations

from abc import ABC, abstractmethod

from pinecone.core.grpc.protos.vector_service_pb2 import Vector

from custom_types import Job


class Vector(ABC):
    @abstractmethod
    def rank(self, job: Job, user_id: str)-> dict[str, float]:
        pass

    # @abstractmethod
    # def rank(self, job_id: str, freelancers: list[Any], job_vectors: dict[str, Vector]):
    #     pass







# What this class needs to be able to do:
# 1. Initialize a vector index
# 2. Embed and add a vector into the index
