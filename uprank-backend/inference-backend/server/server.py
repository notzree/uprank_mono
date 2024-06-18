from queues import Queue
from service import Service
import concurrent.futures


class Server:
    def __init__(self, queue: Queue, service: Service, max_workers: int) -> None:
        self.queue = queue
        self.max_workers = max_workers
        self.service = service
        self.executor = concurrent.futures.ThreadPoolExecutor(max_workers=self.max_workers)

    def start(self):
        while True:
            response = self.queue.poll_ranker_queue()
            if response:
                for item in response:
                    self.executor.submit(self.service.rank, item)
