import { create } from 'zustand';
import { CounterSlice, createCounterSlice } from './slices/counterSlice';
import { createProblemSlice, ProblemSlice } from './slices/problemSlice';

type StoreState = CounterSlice & ProblemSlice;

export const useAppStore = create<StoreState>()((...a) => ({
  ...createCounterSlice(...a),
  ...createProblemSlice(...a),
}));
