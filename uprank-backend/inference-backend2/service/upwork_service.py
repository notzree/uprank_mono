
import json

import requests

from custom_types import Job
from queues import RankingRequest
from vector import Vector

from .service import Service


class UpworkService(Service):
    """A class representing the Upwork service.

    Args:
    ----
        main_backend_url (str): The URL of the main backend.
        pc (any): The Pinecone client.

    Attributes:
    ----------
        main_backend_url (str): The URL of the main backend.
        pc (any): The Pinecone client.

    """

    def __init__(self, main_backend_url: str, inference_server_url: str, vector: Vector) -> None:
        self.main_backend_url = main_backend_url
        self.vector = vector
        self.inference_server_url = inference_server_url

    def rank(self, data: RankingRequest):
        url = f"{self.main_backend_url}/v1/private/jobs/{data.platform}/{data.platform_id}/all_data"
        headers = {
        # "Authorization": f"Bearer {data.short_lived_token}", # Switch to API key
            "Origin": self.inference_server_url,
            "User_id": data.user_id,
        }
        try:
            r = requests.get(
                url,
                headers=headers,
                timeout=10,
            )
        except requests.exceptions.RequestException as e:
            raise Exception(f"Failed to send request: {e}") from e
        if r.status_code != 200:
            raise Exception(f"Failed to get job data: {r.status_code}")

        job_data: Job = Job(**r.json())
        



        ranking_result = self.vector.rank(job_data, user_id=data.user_id)
        with open("ranking_sample_output.json", "w") as f:
            json.dump(ranking_result, f, indent=4)
        # self.vector.rank(job_id=job_data.id, user_id=data.user_id, job_vectors=job_vectors)


# TODO: Figure out how to create clerk jwt from the backend
# Make the get req to get job + freelancer data
# Do the ranking logic

# Or since SQS messsages are encrypted, we can include the users short lived token in the message attribute so we can use it to make the request.
