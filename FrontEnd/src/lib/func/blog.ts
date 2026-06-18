import type { Blog } from "../types";

export function unwrapBlogDetail(data: { blog?: Blog } & Blog): Blog {
  return data.blog ?? data;
}
