import {Configuration} from "@shared/api-client";

export const config = new Configuration({
  basePath: "https://api.example.com", // your API base URL
  // optional: add default headers
  // headers: { Authorization: `Bearer ${token}` }
});

console.log(config);
