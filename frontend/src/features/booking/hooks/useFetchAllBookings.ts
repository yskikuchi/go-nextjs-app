import { getAllBookings } from '@/api/booking';
import { Booking } from '@/types/booking';
import { useEffect, useState } from 'react';

export const useFetchAllBookings = () => {
  const [bookings, setBookings] = useState<Booking[]>([]);

  useEffect(() => {
    const fetchBookings = async () => {
      setBookings(await getAllBookings());
    };

    fetchBookings();
  }, []);

  return bookings;
};
