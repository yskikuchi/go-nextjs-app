import { auth } from '../auth';

// NextAuthConfigのauthorized callbackが呼び出される
// TODO: 型エラーを修正する
export default auth((req) => {
  const { nextUrl } = req;
  const isLoggedIn = !!req.auth;
  const publicRoutes: string[] = ['/', '/bookings'];
  const authRoutes: string[] = ['/admins/login'];
  const DEFAULT_LOGIN_REDIRECT: string = '/admins/dashboard';

  console.log('pathname', nextUrl.pathname);
  console.log('isLoggedIn', isLoggedIn);

  const isPublicRoute = publicRoutes.includes(nextUrl.pathname);
  const isAuthRoute = authRoutes.includes(nextUrl.pathname);

  if (isAuthRoute) {
    if (isLoggedIn) {
      return Response.redirect(new URL(DEFAULT_LOGIN_REDIRECT, nextUrl));
    }
    return null;
  }
  if (!isLoggedIn && !isPublicRoute) {
    return Response.redirect(new URL('/admins/login', nextUrl));
  }

  return null;
});

export const config = {
  matcher: ['/((?!api|_next/static|_next/image|.*\\.png$).*)'],
};
