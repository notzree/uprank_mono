import Link from "next/link";
import { IoMdArrowBack } from "react-icons/io";
export default function Component() {
    return (
        <div className="container mx-auto max-w-3xl py-12 px-4 sm:px-6 lg:px-8">
             <Link className="h-14 justify-start"  href="/">
            <div className="flex items-center space-x-2">
                <IoMdArrowBack />
                <span className="font-semibold">Back to home</span>
            </div>
        </Link>
            <div className="space-y-8">
                <div>
                    <h1 className="text-3xl font-bold tracking-tight text-foreground">
                        Privacy Policy
                    </h1>
                    <p>Effectie Date: July 26, 2024</p>
                    <p className="mt-4 text-muted-foreground">
                        Uprank (&quot;Uprank,&quot; &quot;we,&quot;
                        &quot;us,&quot; or &quot;our&quot;) provides technology
                        to allow our customers to hire freelancers. This privacy
                        policy is designed to help you understand how we
                        collect, use, and share your personal infromation and to
                        help you understand and exercise your privacy rights. By
                        accessing or using the website and its services, you
                        agree to this privacy policy. If you do not agree with
                        the practices described in this policy, please do not
                        use this website.
                    </p>
                </div>
                <h2>1. Information We Collect</h2>
                <h3>1.1 Personal Data</h3>
                <p>We collect the following personal information from you:</p>
                <ul>
                    <li>
                        <strong>Name:</strong> We collect your name to
                        personalize your experience and communicate with you
                        effectively.
                    </li>
                    <li>
                        <strong>Email:</strong> We collect your email address to
                        send you important information regarding your orders,
                        updates, and communication.
                    </li>
                    <li>
                        <strong>Payment Information:</strong> We collect payment
                        details to process your orders securely. However, we do
                        not store your payment information on our servers.
                        Payments are processed by trusted third-party payment
                        processors (Stripe).
                    </li>
                </ul>
                <p>
                    We also use a trusted third-party authentication provider, <a href="https://clerk.com/legal/dpa" target="_blank">Clerk</a>.
                    Clerk may store your Name, Email, and other social login information. Clerk also offers you the ability to register and log in 
                    using your third-party social media account details (like Google, Apple, Microsoft). Where you choose to do this, We and Clerk will receive certain profile
                    information about you from your social media provideer. The profile information we receive may vary depending on the social media provider concerned, but will often include your name, email address, 
                    profile picture as well as other information you choose to make public. We will use the information we receive only for the purposes that are described in this privacy policy or that are otherwise made clear to you on the website.
                    Please note that we do not control, and are not responsible for, other uses of your personal information by your third-party social media provider. 
                    We recommend that you review their privacy policy to understand how they collect, use and share your personal information, and how you can set your privacy preferences on their sites and apps.

                    
                </p>
                <h3>1.2 Non-Personal Data</h3>
                <p>
                    We may use web cookies and similar technologies to collect
                    non-personal information such as your IP address, browser
                    type, device information, and browsing patterns. This
                    information helps us to enhance your browsing experience,
                    analyze trends, and improve our services.
                </p>
                <h2>2. Purpose of Data Collection</h2>
                <p>
                    We collect and use your personal data for the sole purpose
                    of order processing. This includes processing your orders,
                    sending order confirmations, providing customer support, and
                    keeping you updated about the status of your orders.
                </p>
                <h2>3. Data Sharing</h2>
                <p>
                    We do not share your personal data with any third parties
                    except as required for order processing (e.g., sharing your
                    information with payment processors). We do not sell, trade,
                    or rent your personal information to others.
                </p>
                <h2>4. Children&apos;s Privacy</h2>
                <p>
                    Uprank is not intended for children under the age of 13.
                    We do not knowingly collect personal information from
                    children. If you are a parent or guardian and believe that
                    your child has provided us with personal information, please
                    contact us at the email address provided below.
                </p>
                <h2>5. Updates to the Privacy Policy</h2>
                <p>
                    We may update this Privacy Policy from time to time to
                    reflect changes in our practices or for other operational,
                    legal, or regulatory reasons. Any updates will be posted on
                    this page, and we may notify you via email about significant
                    changes.
                </p>
                <h2>6. Contact Information</h2>
                <p>
                    If you have any questions, concerns, or requests related to
                    this Privacy Policy, you can contact us at:
                </p>
                <p>
                    <strong>Email:</strong> richard@uprank.app
                </p>
                <p>
                    By using Uprank, you consent to the terms of this Privacy
                    Policy.
                </p>
            </div>
        </div>
    );
}
