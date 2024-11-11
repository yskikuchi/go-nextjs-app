'use server';

import { redirect } from 'next/navigation';
import { signIn } from '../../auth';
import { AuthError } from 'next-auth';
import { isRedirectError } from 'next/dist/client/components/redirect';

export type LoginState = {
  errors?: {
    name?: string[];
    email?: string[];
    password?: string[];
  };
  value?: {
    name?: string;
    email?: string;
    password?: string;
  };
  message?: string;
};

export const signin = async (
  _state: LoginState,
  formData: FormData
): Promise<LoginState> => {
  // login処理
  try {
    // NEXT_REDIRECTが投げられ，catchでリダイレクトされる
    await signIn('credentials', formData);
    return { message: 'success' };
  } catch (error) {
    if (error instanceof AuthError) {
      switch (error.type) {
        case 'CredentialsSignin':
          console.error('Signin error:', error);
          return {
            message: 'メールアドレスまたはパスワードが間違っています',
          };
      }
    }
    // リダイレクトエラーの場合はリダイレクト
    if (isRedirectError(error)) {
      redirect('/admins/dashboard');
    }
    return {
      message: 'An unexpected error occurred during signin',
    };
  }
};
