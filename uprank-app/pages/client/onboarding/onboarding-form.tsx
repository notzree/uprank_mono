import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { useEffect, useState } from "react";
import { Button } from "@/components/ui/button";
import  craft_api_url  from "@/utils/api_utils/craft_api_url";
import {
    Form,
    FormControl,
    FormDescription,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { useUser } from "@clerk/nextjs";
import { SyncUserBody } from "@/pages/api/private/syncUser";
import { useToast } from "@/components/ui/use-toast";
import { useRouter } from "next/router";
import { Checkbox } from "@/components/ui/checkbox";
import { User } from "@prisma/client";

const formSchema = z.object({
    first_name: z.string().trim().min(1, {
        message: "First Name is required",
    }),
    email: z.string().email({
        message: "Please enter a valid email address",
    }),
    company: z
        .string()
        .trim()
        .min(1, {
            message: "Company name is required",
        }),
});

export default function Home( ){
    const { isLoaded, user } = useUser();
    const { toast } = useToast();
    const router = useRouter();
    const form = useForm<z.infer<typeof formSchema>>({
        resolver: zodResolver(formSchema),
        defaultValues: {
            first_name: user?.firstName ? user.firstName : "",
            email: user?.emailAddresses[0].emailAddress,
            company: "",
        },
    });
    useEffect(() => {
        if (isLoaded && user) {
            form.reset({
                first_name: user.firstName || "",
                email: user?.emailAddresses?.[0]?.emailAddress || '',
            });
        }
    // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [isLoaded, user]); // Depend on isLoaded and user
    async function onSubmit(values: z.infer<typeof formSchema>) {
        console.log("Submitting");
        if (!user){
            return;
        }
        const request_body: SyncUserBody = {
            "user_data":{
                "id": user.id,
                "email": values.email,
                "first_name": values.first_name,
                "company_name": values.company,
                "created_at": new Date(),
                "updated_at": new Date(),
                "last_login": new Date(),
            } as User,
            "metadata": {
                "completed_onboarding": user.unsafeMetadata.completed_onboarding as boolean
            },

        }
        const sync_user = await fetch(craft_api_url("/api/private/syncUser"),{
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({request_body})
        })
        if (!sync_user.ok){
            alert("Failed to sync user data, please try again.");
            return;
        }
        user.update({
            unsafeMetadata: {
                completed_onboarding: true,
            }
          })
        toast({description: "success"});
        setTimeout(() => {
            router.push('client/dashboard');
        }, 500);
    }

    return (
        <div className=" py-16  mx-auto max-w-4xl px-4">
            <div className="space-y-6">
                <div className="space-y-2">
                    <h1 className="text-3xl font-bold">
                        Complete your registration to hire Uprankers
                    </h1>
                    <p className="text-gray-500 dark:text-gray-400">
                        We require a few more details to get started :)
                    </p>
                </div>
                <Form {...form}>
                    <form
                        onSubmit={form.handleSubmit(onSubmit)}
                        className="space-y-8"
                    >
                        <div className="grid grid-cols-2 gap-4">
                            <FormField
                                control={form.control}
                                name="first_name"
                                render={({ field }) => (
                                    <FormItem className="flex-1">
                                        <FormLabel>First Name</FormLabel>
                                        <FormControl>
                                            <Input
                                                {...field}
                                            />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                            <FormField
                                control={form.control}
                                name="email"
                                render={({ field }) => (
                                    <FormItem>
                                        <FormLabel>Email</FormLabel>
                                        <FormControl>
                                            <Input
                                                {...field}
                                            />
                                        </FormControl>
                                        <FormDescription>
                                            Your Email is used for notifying you
                                            of new requests and clients.
                                        </FormDescription>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                            <FormField
                                control={form.control}
                                name="company"
                                render={({ field }) => (
                                    <div className=" flex flex-col space-y-2">
                                        <FormItem>
                                            <FormLabel>Company Name</FormLabel>
                                            <FormControl>
                                                <Input
                                                    placeholder="Uprank Design"
                                                    {...field}
                                                />
                                            </FormControl>
                                            <FormDescription>
                                                Name of your organization or
                                                personal brand.
                                            </FormDescription>
                                            <FormMessage />
                                        </FormItem>
                                    </div>
                                )}
                            />
                        </div>
                        <Button type="submit">Complete sign-up</Button>
                    </form>
                </Form>
            </div>
        </div>
    );
}


