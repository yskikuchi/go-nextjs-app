import LoginForm from '@/features/admin/components/LoginForm';
import { auth } from '../../../../auth';

export default async function Page() {
  const session = await auth();

  console.table(session);
  return (
    <div>
      <LoginForm />
    </div>
  );
}
