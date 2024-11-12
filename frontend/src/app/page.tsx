import Link from 'next/link';

export default function Page() {
  return (
    <div>
      <h1>トップページ</h1>
      <div>
        <Link href="/bookings">予約画面へ</Link>
      </div>
      <div>
        <Link href="/admins/login">管理画面へ</Link>
      </div>
    </div>
  );
}
