import Link from "next/link";
import { IoMdArrowBack } from "react-icons/io";
export default function Component() {
    return (
        <div className="container mx-auto max-w-3xl py-12 px-4 sm:px-6 lg:px-8">
            <Link className="h-14 justify-start" href="/">
                <div className="flex items-center space-x-2">
                    <IoMdArrowBack />
                    <span className="font-semibold">Back to home</span>
                </div>
            </Link>
            <div className="space-y-8">
                <h1>Terms and Conditions for Uprank</h1>
                <p>Last Updated: July 27, 2024</p>

                <p>Welcome to Uprank!</p>

                <p>
                    These Terms of Service (&quot;Terms&quot;) govern your use
                    of the Uprank website at{" "}
                    <a href="https://uprank.io">https://uprank.io</a>{" "}
                    (&quot;Website&quot;) and the services provided by Uprank
                    including the uprank extension. By accessing our Website or using our
                    services, you agree to these Terms.
                </p>

                <h2>1. Description of Uprank</h2>
                <p>
                    Uprank is a platform that helps business owners hire skilled
                    freelancers by providing accurate and unbiased data on
                    freelancers.
                </p>

                <h2>2. Ownership and Usage Rights</h2>
                <p>
                    When you subscribe to Uprank or purchase a one-time job
                    package, you gain the right to access and use our systems to
                    analyze the freelancers that have applied to your job. You
                    own the right to access and use that data to make hiring
                    decisions, but you do not have the right to resell any data
                    provided by Uprank.
                </p>

                <h2>3. User Data and Privacy</h2>
                <p>
                    We collect and store user data, including name, email, and
                    payment information, as necessary to provide our services.
                    For details on how we handle your data, please refer to our
                    Privacy Policy at{" "}
                    <a href="https://uprank.app/privacy-policy">
                        https://uprank.app/privacy-policy
                    </a>
                    .
                </p>

                <h2>4. Non-Personal Data Collection</h2>
                <p>
                    We use web cookies to collect non-personal data for the
                    purpose of improving our services and user experience.
                </p>

                <h2>5. Governing Law</h2>
                <p>These Terms are governed by the laws of Canada.</p>

                <h2>6. Updates to the Terms</h2>
                <p>
                    We may update these Terms from time to time. Users will be
                    notified of any changes via the email address associated
                    with their account.
                </p>

                <p>
                    For any questions or concerns regarding these Terms of
                    Service, please contact us at{" "}
                    <a href="mailto:support@uprank.io">richard@uprank.app</a>.
                </p>
                <h2>7. Disclaimer of Liability</h2>
                <p>
                    Uprank provides AI-powered
                    predictions and tools to assist businesses in hiring
                    freelancers. While we strive to provide accurate and
                    reliable information, Uprank makes no representations or
                    warranties of any kind, express or implied, about the
                    completeness, accuracy, reliability, suitability, or
                    availability with respect to the software or the
                    information, services, or related graphics contained on the
                    platform for any purpose. Any reliance you place on such
                    information is therefore strictly at your own risk. Uprank
                    shall not be responsible or liable for any loss, damage,
                    delay, or other issues that may arise from using our
                    software to decide on hiring a freelancer. The user is solely
                    responsible for their hiring decisions and actions.
                </p>
                <h2>8. Payment Terms</h2>
                <p>
                    All payments for subscriptions or one-time job packages are due in advance. We accept various payment methods and billing cycles. In case of non-payment, your account may be suspended or terminated.
                </p>

                <h2>9. Termination</h2>
                <p>
                    We may terminate or suspend your account if you violate these Terms. Upon termination, your access to the service will be revoked, and your data may be deleted.
                </p>

                <h2>10. Refund Policy</h2>
                <p>
                    We do not offer refunds on our monthly subscriptions, however we do offer refunds on one-time job orders, provided that no job has been processed.
                </p>

                <p>Thank you for using Uprank!</p>

                <p>Terms of service | Uprank</p>
            </div>
        </div>
    );
}
