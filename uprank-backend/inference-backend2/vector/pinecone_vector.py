from __future__ import annotations

from typing import Any

from langchain_openai import OpenAIEmbeddings
from pinecone import ServerlessSpec
from pinecone.core.grpc.protos.vector_service_pb2 import Vector
from pinecone.grpc import PineconeGRPC

from custom_types import Job

# from pinecone import Pinecone
from .vector import Vector


class PineconeVector(Vector):

    def __init__(self, pc: PineconeGRPC, index_name: str, # noqa: D417
            embd: OpenAIEmbeddings) -> None:
        """Initialize the vector class, create the index if it does not exist, and assign the index.

        Args:
        ----
            pc (any): The Pinecone client object.
            index_name (str): The name of the index.

        Returns:
        -------
            None

        """
        self.pc = pc
        self.embd = embd
        self.description_index = "description"
        self.skills_index = "skills"
        if index_name not in pc.list_indexes().names():
            # hardcode the index details for now
            index = pc.create_index(
                name=index_name,
                dimension=1536,
                metric="cosine",
                spec=ServerlessSpec(
                    cloud="aws",
                    region="us-east-2",
                ),
            )
        else:
            index = pc.Index(index_name)
        self.index = index


    def rank(self, job:Job, user_id: str)-> dict[str, float]:
        result = {}
        # try:
        #     description_vector = self.create_vector(f"{job.id}_description", job.description, {"user_id": user_id, "platform": "upwork", "job_id": job.id})
        #     upsert_description_response = self.index.upsert(
        #         namespace=self.description_index,
        #         vectors=[
        #             description_vector,
        #         ],
        #     )
        # except Exception as e:
        #     error_message = "Failed to upsert description vector"
        #     raise Exception(error_message) from e # noqa: TRY002
        # try:
        #     skills_vector = self.create_vector(f"{job.id}_skills", " ".join(job.skills), {"user_id": user_id, "platform": "upwork", "job_id": job.id})
        #     upsert_skills_response = self.index.upsert(
        #         namespace=self.skills_index,
        #         vectors=[
        #             skills_vector,
        #         ],
        #     )
        # except Exception as e:
        #     error_message = "Failed to upsert skills vector"
        #     raise Exception(error_message) from e  # noqa: TRY002

        # try:
        #     upwork_freelancers: list[Any] = job.edges["upworkfreelancer"]
        #     for freelancer in upwork_freelancers:
        #         freelancer_id = freelancer["id"]
        #         freelancer_description_vector = self.create_vector(freelancer["id"], freelancer["description"], {"user_id": user_id, "job_id": job.id, "type": "freelancer_description"})
        #         upsert_freelancer_response = self.index.upsert(
        #             namespace=self.description_index,
        #             vectors=[
        #                 freelancer_description_vector,
        #             ],
        #         )
        #         freelancer_skill_vector = self.create_vector(freelancer["id"], " ".join(freelancer["skills"]), {"user_id": user_id, "job_id": job.id, "type": "freelancer_skills"})
        #         upsert_freelancer_response = self.index.upsert(
        #             namespace=self.skills_index,
        #             vectors=[
        #                 freelancer_skill_vector,
        #             ],
        #         )

        #         work_histories = freelancer["edges"].get("work_histories", [])
        #         if len(work_histories) == 0:
        #             continue
        #         for work_history in work_histories:
        #             description = work_history.get("description", "")
        #             if description == "":
        #                 continue
        #             try:
        #                 work_history_vector = self.create_vector(f"{freelancer['id']}_{work_history['id']}", description, {"freelancer_id": freelancer_id, "job_id": job.id, "type": "work_history"})
        #                 upsert_work_history_response = self.index.upsert(
        #                     namespace=self.description_index,
        #                     vectors=[
        #                         work_history_vector,
        #                     ],
        #                 )
        #             except Exception as e:
        #                 print("exception occured")
        #                 print(e)
        # except Exception as e:
        #     print(e)
        #     error_message = "Failed to upsert freelancer vector"
        #     raise Exception(error_message) from e
        # print("computing similarity")
        description_vector = self.embd.embed_query(job.description)
        upwork_freelancers: list[Any] = job.edges["upworkfreelancer"]
        for freelancer in upwork_freelancers:
            try:
                response = self.index.query(
                    namespace=self.description_index,
                    vector=description_vector,
                    top_k=100,
                    filter={
                        "job_id": job.id,
                        "type": "work_history",
                        "freelancer_id": freelancer["id"],
                    },
                )
                similarities = [match["score"] for match in response["matches"]]
                result[freelancer["id"]] = similarities
            except Exception as e:
                print(e)
                error_message = "Failed to query vectors"
                raise Exception(error_message) from e
        print("Finsihed ranking")
        return result


    def create_vector(self, vector_id: str, text: str, metadata: dict[str, str]) -> Vector:
        vector = self.embd.embed_query(text)
        return (
            vector_id, vector, metadata,
        )
    

