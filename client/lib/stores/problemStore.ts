import { create } from 'zustand';
import { createProblemSlice, ProblemSlice } from '../slices/problemSlice';

type StoreState = ProblemSlice;

const useProblemStore = create<StoreState>()((...a) => ({
  ...createProblemSlice(...a),
}));

export default useProblemStore;
