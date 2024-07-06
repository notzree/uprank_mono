import Image from "next/image";
import { Inter } from "next/font/google";
const inter = Inter({ subsets: ["latin"] });
import Link from "next/link";
import { Button } from "@/components/ui/button";
import { CardHeader, CardContent, Card } from "@/components/ui/card";
import { SignOutButton } from "@clerk/nextjs";
import { useUser } from "@clerk/nextjs";
type CardData = {
    title: string;
    content: string;
};

const FreelancerCardData: CardData[] = [
    {
        title: "AI powered insights",
        content:
            "Make smarter hiring decisions with our AI-powered insights. Our advanced algorithms provide a job compatibility score, analyzing how well a freelancer's past experiences align with your project needs. Additionally, get AI-driven estimated completion times to better plan and manage your projects.",
    },
    {
        title: "Comprehensive Freelancer Metrics",
        content:
            "Make informed hiring decisions with our extensive freelancer metrics. Access vital statistics like budget adherence rates, average client hire count, and more, all designed to help you assess the quality and reliability of freelancers effectively.",
    },
    {
        title: " Custom Filtering and Data Analysis",
        content:
            "Our powerful spreadsheet interface allows you to apply custom filters on specific freelancer data, making it easy to narrow down your search and find the ideal candidate for your project. Leverage data analysis tools to gain deeper insights and make well-informed decisions.",
    },
];
export default function Component() {
    const { user } = useUser();
    return (
        <div key="1" className="flex flex-col min-h-screen">
            <header className="flex items-center justify-between px-8 py-6 border-b">
                <div className="flex items-center">
                    <ReplyIcon className="w-6 h-6 mr-2" />
                    <Link className="text-lg font-bold" href="#">
                        Uprank
                    </Link>
                </div>
                <nav className="flex gap-4">
                    {user?.id ? (
                        <Link
                            className="text-sm font-medium hover:underline"
                            href="/client/dashboard"
                        >
                            Dashboard
                        </Link>
                    ) : (
                        <>
                            <Link
                                className="text-sm font-medium hover:underline"
                                href="/sign-up"
                            >
                                Sign up
                            </Link>
                            <p>or</p>
                            <Link
                                className="text-sm font-medium hover:underline"
                                href="/sign-in"
                            >
                                Sign in
                            </Link>
                        </>
                    )}
                </nav>
            </header>
            <main className="flex-1">
                <section className="py-16 px-10 text-center">
                    <h1 className="text-4xl font-bold">
                        We help business owners hire the best freelancers
                    </h1>
                    <div className="flex justify-center items-center">
                        <p className="mt-4 text-lg text-foreground w-[40rem] text-left">
                            Discover a smarter way to hire freelancers with our
                            cutting-edge SaaS solution. Our Chrome extension and
                            AI recommender algorithm streamline the hiring
                            process, offering powerful insights and
                            comprehensive data to help you find the perfect
                            match for your project. Elevate your hiring strategy
                            with custom filters, AI-powered insights, and
                            essential metrics, all within an intuitive
                            spreadsheet ui.
                        </p>
                    </div>
                </section>
                <section className="py-2 px-10 bg-background" id="features">
                    <h2 className="text-3xl font-bold text-center">Features</h2>
                    <div className="mt-10 grid gap-10 grid-cols-1 sm:grid-cols-2 lg:grid-cols-3">
                        {FreelancerCardData.map(
                            (card: CardData, index: any) => (
                                <Card key={index}>
                                    <CardHeader>
                                        <h3 className="text-lg font-bold">
                                            {card.title}
                                        </h3>
                                    </CardHeader>
                                    <CardContent>
                                        <p className="text-sm text-gray-600">
                                            {card.content}
                                        </p>
                                    </CardContent>
                                </Card>
                            )
                        )}
                    </div>
                </section>

                <section className="py-24 px-10 text-center">
                    <h2 className="text-3xl font-bold">Sign up for the closed beta today</h2>
                    <p className="mt-6 text-lg text-gray-600">
                        Join thousands of business owners who have already enhanced their hiring strategy with Uprank.
                    </p>
                    <Button className="mt-10">Register for Beta</Button>
                </section>
            </main>
            <footer className="py-6 px-10 text-center text-sm text-gray-600">
                © Uprank Design. All Rights Reserved.
            </footer>
        </div>
    );
}

function ReplyIcon(props: any) {
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
            <polyline points="9 17 4 12 9 7" />
            <path d="M20 18v-2a4 4 0 0 0-4-4H4" />
        </svg>
    );
}
