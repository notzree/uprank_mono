

class RankingRequest:
    def __init__(self, body: str, job_id: str, user_id: str, receipt_handle: str):
        self.body = body
        self.job_id = job_id
        self.user_id = user_id
        self.receipt_handle = receipt_handle

