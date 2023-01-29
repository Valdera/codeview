import { Counter } from '@lib/types/counter';
import { StateCreator } from 'zustand';

export interface CounterSlice {
  counter: Counter;
  increment: () => void;
  decrement: () => void;
}

export const createCounterSlice: StateCreator<CounterSlice> = (set, get) => ({
  counter: {
    current: 0,
  },
  increment: () => {
    const counter = get().counter;
    counter.current += 1;
    set({ counter });
  },
  decrement: () => {
    const counter = get().counter;
    counter.current -= 1;
    set({ counter });
  },
});
