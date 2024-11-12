'use client';

import { useFetchAllBookings } from '../hooks/useFetchAllBookings';

export default function Bookings() {
  const bookings = useFetchAllBookings();

  return (
    <div className="relative overflow-x-auto">
      <table className="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
        <thead className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
          <tr>
            <th scope="col" className="px-6 py-3">
              車両名
            </th>
            <th scope="col" className="px-6 py-3">
              利用開始日時
            </th>
            <th scope="col" className="px-6 py-3">
              利用終了日時
            </th>
            <th scope="col" className="px-6 py-3">
              料金
            </th>
            <th scope="col" className="px-6 py-3">
              ステータス
            </th>
          </tr>
        </thead>
        <tbody>
          {bookings.map((booking) => {
            return (
              <tr
                key={booking.id}
                className="bg-white border-b dark:bg-gray-800 dark:border-gray-700"
              >
                <td className="px-6 py-4">{booking.car?.name}</td>

                <th
                  scope="row"
                  className="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white"
                >
                  {booking.startTime}
                </th>
                <td className="px-6 py-4">{booking.endTime}</td>
                <td className="px-6 py-4">{booking.amount}</td>
                <td className="px-6 py-4">{booking.status}</td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </div>
  );
}
