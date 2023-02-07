import { create } from 'zustand';
import { createProblemSlice, ProblemSlice } from '../slices/problemSlice';

type StoreState = ProblemSlice;

const useAppStore = create<StoreState>()((...a) => ({
  ...createProblemSlice(...a),
}));

export default useAppStore;
