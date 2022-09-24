import { serve } from "https://deno.land/std@0.91.0/http/server.ts";

const server = serve({ hostname: "0.0.0.0", port: 8080 });
console.log(`HTTP Webserver is running. Access it at http://localhost:8000/`);

for await (const request of server) {
  let bodyContent = "Your user-agent is:\n";
  bodyContent += request.headers.get("user-agent") || "Unknown";
  bodyContent += "\nHello World!\n";

  request.respond({ status: 200, body: bodyContent });
}
