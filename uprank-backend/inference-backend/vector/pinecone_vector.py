from langchain_openai import OpenAiEmbeddings
from pinecone import ServerlessSpec

from .vector import Vector


class PineconeVector(Vector):

    def __init__(self, pc: any, index_name: str, # noqa: D417
            embd: OpenAiEmbeddings) -> None:
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
        if index_name not in pc.list_indexs().names():
            # hardcode the index details for now
            pc.create_index(
                name=index_name,
                dimension=1536,
                metric="cosine",
                spec=ServerlessSpec(
                    cloud="aws",
                    region="us-east-2",
                ),
            )
        self.index = self.pc.Index(index_name)

