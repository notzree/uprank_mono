import { AvatarImage, AvatarFallback, Avatar } from "@/components/ui/avatar";
import { Button, buttonVariants } from "@/components/ui/button";
import Link from "next/link";
import { CardContent, Card } from "@/components/ui/card";
import { Freelancer } from "@/types/freelancer";
import Image from "next/image";
import { Badge } from "@/components/ui/badge";
import  craft_api_url  from "@/utils/api_utils/craft_api_url";
import { useUser } from "@clerk/nextjs";
import { useToast } from "@/components/ui/use-toast";
import { capitalizeFirstLetter } from "@/utils/client_utils/capitalize_first_letter";
import Router from "next/router";
interface FreelancerCardProps {
    freelancer: Freelancer;
}

export default function FreelancerCard({ freelancer }: FreelancerCardProps) {
    const { user } = useUser();
    const { toast } = useToast();
   const router = Router;

    return (
        <Card className="w-full max-w-xl flex flex-col md:flex-row items-start gap-6 p-6">
            <div className="flex-1 space-y-4">
                <div className="flex items-start gap-4">
                    <Avatar>
                        <AvatarImage
                            src={freelancer.profile_image_url as string}
                        />
                        <AvatarFallback>
                            {freelancer.first_name + " " + freelancer.last_name}
                        </AvatarFallback>
                    </Avatar>
                    <div className="flex-1 space-y-1">
                        <div className="flex items-center gap-2">
                            <h4 className="text-md font-semibold ">
                                {capitalizeFirstLetter(freelancer.first_name) +
                                    " " +
                                    capitalizeFirstLetter(freelancer.last_name)}
                            </h4>
                            <Badge
                                className="px-2 py-1 text-xs"
                                variant="outline"
                            >
                                Rank
                                {/*  todo: add ranking */}
                            </Badge>
                        </div>
                    </div>
                </div>
                <p className="text-sm leading-relaxed text-gray-500 dark:text-gray-400">
                    {freelancer.about_me}
                </p>
                <div className="flex items-center justify-between gap-4 text-sm text-gray-500 dark:text-gray-400">
                    <div className="flex item-center justify-evenly gap-6">
                        <div className="flex items-center gap-1">
                            <UsersIcon className="h-4 w-4" />
                            <span>1.2K</span>
                        </div>
                        <div className="flex items-center gap-1">
                            <UsersIcon className="h-4 w-4" />
                            <span>345</span>
                        </div>
                        <div className="flex items-center gap-1">
                            <DollarSignIcon className="h-4 w-4" />
                            <span>{freelancer.current_monthly_rate}</span>
                        </div>
                        <div className="flex items-center gap-1">
                            <CalendarDaysIcon className="h-4 w-4" />
                            <span>Monthly</span>
                        </div>
                    </div>
                </div>
                <div className="flex gap-2">
                    <Button className="flex-1" variant="outline">
                        Request Call
                    </Button>
                    <Button
                        className="flex-1"
                        variant="outline"
                        onClick={() => router.push(craft_api_url(`/client/dashboard/subscription/${freelancer.id}/create`))}
                    >
                        Subscribe
                    </Button>
                </div>
            </div>
            <div className="flex-1 space-y-4">
                <h4 className="text-lg font-semibold">Portfolio Preview</h4>
                <div className=" relative h-[100px]  flex justify-center items-center pb-1">
                    <Image
                        alt="Portfolio Image"
                        className="aspect-[5/3] object-cover rounded-md"
                        src="/placeholder.svg"
                        fill={true}
                    />
                </div>
                {freelancer.portfolio_url && (
                    <Link
                        href={freelancer.portfolio_url ?? "/"}
                        target="_blank"
                        className={buttonVariants({ variant: "outline" })}
                    >
                        View Full Portfolio
                    </Link>
                )}
            </div>
        </Card>
    );
}



function CalendarDaysIcon(props: any) {
    return (
        <svg
            {...props}
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
        >
            <rect width="18" height="18" x="3" y="4" rx="2" ry="2" />
            <line x1="16" x2="16" y1="2" y2="6" />
            <line x1="8" x2="8" y1="2" y2="6" />
            <line x1="3" x2="21" y1="10" y2="10" />
            <path d="M8 14h.01" />
            <path d="M12 14h.01" />
            <path d="M16 14h.01" />
            <path d="M8 18h.01" />
            <path d="M12 18h.01" />
            <path d="M16 18h.01" />
        </svg>
    );
}

function DollarSignIcon(props: any) {
    return (
        <svg
            {...props}
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
        >
            <line x1="12" x2="12" y1="2" y2="22" />
            <path d="M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6" />
        </svg>
    );
}

function UsersIcon(props: any) {
    return (
        <svg
            {...props}
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            strokeLinecap="round"
            strokeLinejoin="round"
        >
            <path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2" />
            <circle cx="9" cy="7" r="4" />
            <path d="M22 21v-2a4 4 0 0 0-3-3.87" />
            <path d="M16 3.13a4 4 0 0 1 0 7.75" />
        </svg>
    );
}
