export const ROUTES = {
  home: '/',
  about: '/about',
  blog: '/blog',
  works: '/works',
  story: '/story',
  gallery: '/gallery',
  contact: '/contact',

  division: {
    list: '/division',
    detail: (slug: string) => `/division/${slug}`,
  },

  admin: {
    dashboard: '/admin',
    login: '/admin/login',
    coreTeam: '/admin/core-team',
    blog: '/admin/blog',
  },
} as const
