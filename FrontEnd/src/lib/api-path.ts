export const API_PATH = {
  auth: {
    login: "/auth/login",
    register: "/auth/register",
    refresh: "/auth/refresh",
    logout: "/auth/logout",
  },

  user: {
    me: "/user/me",
    profile: "/user/profile",
    changePassword: "/user/change-password",
    list: "/user",
    detail: (id: number | string) => `/user/${id}`,
  },

  admin: {
    user: {
      list: "/admin/user",
      superAdmin: "/admin/user/super-admin",
      changePassword: (id: number | string) =>
        `/admin/user/${id}/change-password`,
    },
    blogs: {
      list: "/admin/blogs",
      detail: (id: number | string) => `/admin/blogs/${id}`,
    },
    gallery: {
      list: "/admin/gallery",
      detail: (id: number | string) => `/admin/gallery/${id}`,
    },
    works: {
      list: "/admin/works",
      detail: (id: number | string) => `/admin/works/${id}`,
      status: (id: number | string) => `/admin/works/${id}/status`,
    },
    pengurus: {
      list: "/admin/pengurus",
      detail: (id: number | string) => `/admin/pengurus/${id}`,
      byUser: (userId: number | string) => `/admin/pengurus/by-user/${userId}`,
      delete: (id: number | string) => `/admin/pengurus/delete/${id}`,
    },
  },

  blogs: {
    list: "/blogs",
    detail: (id: number | string) => `/blogs/${id}`,
  },

  gallery: {
    list: "/gallery",
  },

  works: {
    list: "/works",
    detail: (id: number | string) => `/works/${id}`,
    types: "/works/types",
  },

  pengurus: {
    byDivision: (division: string) => `/pengurus/division/${division}`,
    createProfile: "/pengurus",
    profile: "/pengurus/profile",
    me: "/pengurus/me",
  },

  upload: {
    file: "/upload/file",
    files: "/upload/files",
  },
} as const;
