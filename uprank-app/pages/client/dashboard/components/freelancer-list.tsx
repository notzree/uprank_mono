

import { cn } from "@/lib/utils"
import { ScrollArea } from "@/components/ui/scroll-area"
import { Freelancer } from "@/types/freelancer"
import FreelancerCard from "./freelancer-card"

interface FreelancerListProps {
  freelancers: Freelancer[],
}

export function FreelancerList({ freelancers: freelancers }: FreelancerListProps) {
  return (
    <ScrollArea className="h-screen">
      <div className="flex flex-col gap-2 p-4 pt-0">
        {freelancers.map((freelancer) => (
          //TODO: MOVE THE BUTTON INTO THE FREELANCER CARD TO AVOID HYDRATION HERROR
            <FreelancerCard key = {freelancer.id} freelancer={freelancer}  />
          
        ))}
      </div>
    </ScrollArea>
  )
}

