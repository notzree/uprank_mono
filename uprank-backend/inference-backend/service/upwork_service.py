from .service import Service
from queues import RankingRequest
import requests

class UpworkService(Service):
    def __init__(self, main_backend_url: str):
        self.main_backend_url = main_backend_url

    def rank(self, data: RankingRequest):
        #Do ranking
        print(data)
        r = requests.get(self.main_backend_url +"/v1/private/jobs/{data.job_id}/upwork/)

        return

## TODO: Figure out how to create clerk jwt from the backend
## Make the get req to get job + freelancer data
## Do the ranking logic