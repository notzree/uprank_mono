import concurrent.futures

from queues import Queue
from queues.types import RankingRequest
from service import Service


class Server:
    def __init__(self, queue: Queue, service: Service, max_workers: int) -> None:
        self.queue = queue
        self.max_workers = max_workers
        self.service = service
        self.executor = concurrent.futures.ThreadPoolExecutor(
            max_workers=self.max_workers,
        )

    def start(self) -> None:
        while True:
            response = self.queue.poll_ranker_queue()
            if response:
                for item in response:
                    self.executor.submit(self.rank_and_delete, item)

    def rank_and_delete(self, item: RankingRequest) -> None:
        self.service.rank(item)
        self.queue.delete_message(item.receipt_handle)

