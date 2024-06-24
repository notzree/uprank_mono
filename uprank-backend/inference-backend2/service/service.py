from abc import ABC, abstractmethod

from queues import RankingRequest


class Service(ABC):
    @abstractmethod
    def rank(self, data: RankingRequest):
        pass
