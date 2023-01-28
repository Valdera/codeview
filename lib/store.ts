import { create } from 'zustand';
import { CounterSlice, createCounterSlice } from './slices/counterSlice';

type StoreState = CounterSlice;

export const useAppStore = create<StoreState>()((...a) => ({
  ...createCounterSlice(...a),
}));
