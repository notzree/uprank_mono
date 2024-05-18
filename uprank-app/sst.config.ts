/// <reference path="./.sst/platform/config.d.ts" />

export default $config({
  app(input) {
    return {
      name: "uprank-dev",
      removal: input?.stage === "production" ? "retain" : "remove",
      home: "aws",
    };
  },
  async run() {
    const queue = new sst.aws.Queue("ScrapeRequestQueue",{
      fifo: true,
    });
    queue.subscribe("pages/api/private/job/")

    new sst.aws.Nextjs("MyWeb",
      {
        link: [queue]
      }
    );
  },
});
