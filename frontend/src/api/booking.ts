"use server";

import { BookingCreatePayload } from '@/types/booking';
import axios from 'axios';
import { auth } from '../../auth';

export const getBookingsSummaries = async () => {
  const result = await axios.get(
    `${process.env.NEXT_PUBLIC_API_URL}/bookings/summaries`
  );
  return result.data;
};

export const postBooking = async (data: BookingCreatePayload) => {
  const result = await axios.post(
    `${process.env.NEXT_PUBLIC_API_URL}/bookings`,
    data
  );

  if (result.status !== 201) {
    throw new Error('Failed to create a booking');
  }

  return result.data;
};

export const getAllBookings = async () => {
  const session = await auth();

  // sessionからjwtを取得して、リクエストヘッダーにセットする
  const result = await axios.get(
    `${process.env.NEXT_PUBLIC_API_URL}/bookings`,
    {
      headers: {
        Authorization: `Bearer ${session?.backendToken}`,
      },
    }
  );
  return result.data;
};
